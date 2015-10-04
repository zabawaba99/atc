package resource

import (
	"github.com/concourse/atc"
	"github.com/concourse/atc/worker"
	"github.com/concourse/baggageclaim"
	"github.com/pivotal-golang/lager"
)

type ResourceType string
type ContainerImage string

type Session struct {
	ID        worker.Identifier
	Ephemeral bool
}

//go:generate counterfeiter . Tracker

type Tracker interface {
	Init(lager.Logger, Metadata, Session, ResourceType, atc.Tags) (Resource, error)
	InitWithCache(lager.Logger, Metadata, Session, ResourceType, atc.Tags, CacheIdentifier) (Resource, Cache, error)
}

//go:generate counterfeiter . Cache

type Cache interface {
	IsInitialized() (bool, error)
	Initialize() error
}

type Metadata interface {
	Env() []string
}

type EmptyMetadata struct{}

func (m EmptyMetadata) Env() []string { return nil }

type tracker struct {
	workerClient worker.Client
}

type TrackerFactory struct{}

func (TrackerFactory) TrackerFor(client worker.Client) Tracker {
	return NewTracker(client)
}

func NewTracker(workerClient worker.Client) Tracker {
	return &tracker{
		workerClient: workerClient,
	}
}

type VolumeMount struct {
	Volume    baggageclaim.Volume
	MountPath string
}

func (tracker *tracker) Init(logger lager.Logger, metadata Metadata, session Session, typ ResourceType, tags atc.Tags) (Resource, error) {
	logger = logger.Session("init")

	logger.Debug("start")
	defer logger.Debug("done")

	container, found, err := tracker.workerClient.FindContainerForIdentifier(logger, session.ID)
	if err != nil {
		logger.Error("failed-to-look-for-existing-container", err)
		return nil, err
	}

	if found {
		logger.Info("found-existing", lager.Data{"container": container.Handle()})
		return NewResource(container), nil
	}

	logger.Info("creating-container", lager.Data{"identifier": session.ID})

	container, err = tracker.workerClient.CreateContainer(logger, session.ID, worker.ResourceTypeContainerSpec{
		Type:      string(typ),
		Ephemeral: session.Ephemeral,
		Tags:      tags,
		Env:       metadata.Env(),
	})
	if err != nil {
		return nil, err
	}

	logger.Info("created", lager.Data{"container": container.Handle()})

	return NewResource(container), nil
}

func (tracker *tracker) InitWithCache(logger lager.Logger, metadata Metadata, session Session, typ ResourceType, tags atc.Tags, cacheIdentifier CacheIdentifier) (Resource, Cache, error) {
	logger = logger.Session("init-with-cache")

	logger.Debug("start")
	defer logger.Debug("done")

	container, found, err := tracker.workerClient.FindContainerForIdentifier(logger, session.ID)
	if err != nil {
		logger.Error("failed-to-look-for-existing-container", err)
		return nil, nil, err
	}

	if found {
		logger.Info("found-existing", lager.Data{"container": container.Handle()})

		var cache Cache

		volumes := container.Volumes()
		switch len(volumes) {
		case 0:
			logger.Info("no-cache")
			cache = noopCache{}
		default:
			logger.Info("found-cache")
			cache = volumeCache{volumes[0]}
		}

		return NewResource(container), cache, nil
	}

	logger.Info("no-existing-container")

	resourceSpec := worker.WorkerSpec{
		ResourceType: string(typ),
		Tags:         tags,
	}

	chosenWorker, err := tracker.workerClient.Satisfying(resourceSpec)
	if err != nil {
		logger.Info("no-workers-satisfying-spec", lager.Data{
			"error": err.Error(),
		})
		return nil, nil, err
	}

	vm, hasVM := chosenWorker.VolumeManager()
	if !hasVM {
		logger.Info("creating-container-without-cache", lager.Data{"identifier": session.ID})

		container, err := chosenWorker.CreateContainer(logger, session.ID, worker.ResourceTypeContainerSpec{
			Type:      string(typ),
			Ephemeral: session.Ephemeral,
			Tags:      tags,
			Env:       metadata.Env(),
		})
		if err != nil {
			logger.Error("faild-to-create-container", err)
			return nil, nil, err
		}

		logger.Info("created", lager.Data{"container": container.Handle()})

		return NewResource(container), noopCache{}, nil
	}

	cachedVolume, cacheFound, err := cacheIdentifier.FindOn(logger, vm)
	if err != nil {
		logger.Error("failed-to-look-for-cache", err)
		return nil, nil, err
	}

	if cacheFound {
		logger.Info("found-cache", lager.Data{"volume": cachedVolume.Handle()})
	} else {
		logger.Info("no-cache-found")

		cachedVolume, err = cacheIdentifier.CreateOn(logger, vm)
		if err != nil {
			return nil, nil, err
		}

		logger.Info("new-cache", lager.Data{"volume": cachedVolume.Handle()})
	}

	defer cachedVolume.Release()

	logger.Info("creating-container-with-cache", lager.Data{"identifier": session.ID})

	container, err = chosenWorker.CreateContainer(logger, session.ID, worker.ResourceTypeContainerSpec{
		Type:      string(typ),
		Ephemeral: session.Ephemeral,
		Tags:      tags,
		Env:       metadata.Env(),
		Cache: worker.VolumeMount{
			Volume:    cachedVolume,
			MountPath: ResourcesDir("get"),
		},
	})
	if err != nil {
		logger.Error("failed-to-create-container", err)
		return nil, nil, err
	}

	logger.Info("created", lager.Data{"container": container.Handle()})

	return NewResource(container), volumeCache{cachedVolume}, nil
}
