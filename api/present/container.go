package present

import (
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
)

func Container(container db.Container) atc.Container {
	var stepType string
	if container.Type != db.ContainerTypeCheck {
		stepType = container.Type.String()
	}
	return atc.Container{
		ID:                   container.Handle,
		WorkerName:           container.ContainerMetadata.WorkerName,
		PipelineName:         container.PipelineName,
		JobName:              container.JobName,
		BuildName:            container.BuildName,
		BuildID:              container.ContainerMetadata.BuildID,
		StepType:             stepType,
		StepName:             container.StepName,
		ResourceName:         container.ResourceName,
		WorkingDirectory:     container.WorkingDirectory,
		EnvironmentVariables: container.EnvironmentVariables,
	}
}
