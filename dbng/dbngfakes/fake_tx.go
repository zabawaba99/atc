// This file was generated by counterfeiter
package dbngfakes

import (
	"database/sql"
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeTx struct {
	CommitStub        func() error
	commitMutex       sync.RWMutex
	commitArgsForCall []struct{}
	commitReturns     struct {
		result1 error
	}
	ExecStub        func(query string, args ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		query string
		args  []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	PrepareStub        func(query string) (*sql.Stmt, error)
	prepareMutex       sync.RWMutex
	prepareArgsForCall []struct {
		query string
	}
	prepareReturns struct {
		result1 *sql.Stmt
		result2 error
	}
	QueryStub        func(query string, args ...interface{}) (*sql.Rows, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		query string
		args  []interface{}
	}
	queryReturns struct {
		result1 *sql.Rows
		result2 error
	}
	QueryRowStub        func(query string, args ...interface{}) *sql.Row
	queryRowMutex       sync.RWMutex
	queryRowArgsForCall []struct {
		query string
		args  []interface{}
	}
	queryRowReturns struct {
		result1 *sql.Row
	}
	RollbackStub        func() error
	rollbackMutex       sync.RWMutex
	rollbackArgsForCall []struct{}
	rollbackReturns     struct {
		result1 error
	}
	StmtStub        func(stmt *sql.Stmt) *sql.Stmt
	stmtMutex       sync.RWMutex
	stmtArgsForCall []struct {
		stmt *sql.Stmt
	}
	stmtReturns struct {
		result1 *sql.Stmt
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTx) Commit() error {
	fake.commitMutex.Lock()
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct{}{})
	fake.recordInvocation("Commit", []interface{}{})
	fake.commitMutex.Unlock()
	if fake.CommitStub != nil {
		return fake.CommitStub()
	} else {
		return fake.commitReturns.result1
	}
}

func (fake *FakeTx) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeTx) CommitReturns(result1 error) {
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Exec", []interface{}{query, args})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(query, args...)
	} else {
		return fake.execReturns.result1, fake.execReturns.result2
	}
}

func (fake *FakeTx) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeTx) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].query, fake.execArgsForCall[i].args
}

func (fake *FakeTx) ExecReturns(result1 sql.Result, result2 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) Prepare(query string) (*sql.Stmt, error) {
	fake.prepareMutex.Lock()
	fake.prepareArgsForCall = append(fake.prepareArgsForCall, struct {
		query string
	}{query})
	fake.recordInvocation("Prepare", []interface{}{query})
	fake.prepareMutex.Unlock()
	if fake.PrepareStub != nil {
		return fake.PrepareStub(query)
	} else {
		return fake.prepareReturns.result1, fake.prepareReturns.result2
	}
}

func (fake *FakeTx) PrepareCallCount() int {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return len(fake.prepareArgsForCall)
}

func (fake *FakeTx) PrepareArgsForCall(i int) string {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return fake.prepareArgsForCall[i].query
}

func (fake *FakeTx) PrepareReturns(result1 *sql.Stmt, result2 error) {
	fake.PrepareStub = nil
	fake.prepareReturns = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	fake.queryMutex.Lock()
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("Query", []interface{}{query, args})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(query, args...)
	} else {
		return fake.queryReturns.result1, fake.queryReturns.result2
	}
}

func (fake *FakeTx) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *FakeTx) QueryArgsForCall(i int) (string, []interface{}) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].query, fake.queryArgsForCall[i].args
}

func (fake *FakeTx) QueryReturns(result1 *sql.Rows, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeTx) QueryRow(query string, args ...interface{}) *sql.Row {
	fake.queryRowMutex.Lock()
	fake.queryRowArgsForCall = append(fake.queryRowArgsForCall, struct {
		query string
		args  []interface{}
	}{query, args})
	fake.recordInvocation("QueryRow", []interface{}{query, args})
	fake.queryRowMutex.Unlock()
	if fake.QueryRowStub != nil {
		return fake.QueryRowStub(query, args...)
	} else {
		return fake.queryRowReturns.result1
	}
}

func (fake *FakeTx) QueryRowCallCount() int {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return len(fake.queryRowArgsForCall)
}

func (fake *FakeTx) QueryRowArgsForCall(i int) (string, []interface{}) {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return fake.queryRowArgsForCall[i].query, fake.queryRowArgsForCall[i].args
}

func (fake *FakeTx) QueryRowReturns(result1 *sql.Row) {
	fake.QueryRowStub = nil
	fake.queryRowReturns = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *FakeTx) Rollback() error {
	fake.rollbackMutex.Lock()
	fake.rollbackArgsForCall = append(fake.rollbackArgsForCall, struct{}{})
	fake.recordInvocation("Rollback", []interface{}{})
	fake.rollbackMutex.Unlock()
	if fake.RollbackStub != nil {
		return fake.RollbackStub()
	} else {
		return fake.rollbackReturns.result1
	}
}

func (fake *FakeTx) RollbackCallCount() int {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	return len(fake.rollbackArgsForCall)
}

func (fake *FakeTx) RollbackReturns(result1 error) {
	fake.RollbackStub = nil
	fake.rollbackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) Stmt(stmt *sql.Stmt) *sql.Stmt {
	fake.stmtMutex.Lock()
	fake.stmtArgsForCall = append(fake.stmtArgsForCall, struct {
		stmt *sql.Stmt
	}{stmt})
	fake.recordInvocation("Stmt", []interface{}{stmt})
	fake.stmtMutex.Unlock()
	if fake.StmtStub != nil {
		return fake.StmtStub(stmt)
	} else {
		return fake.stmtReturns.result1
	}
}

func (fake *FakeTx) StmtCallCount() int {
	fake.stmtMutex.RLock()
	defer fake.stmtMutex.RUnlock()
	return len(fake.stmtArgsForCall)
}

func (fake *FakeTx) StmtArgsForCall(i int) *sql.Stmt {
	fake.stmtMutex.RLock()
	defer fake.stmtMutex.RUnlock()
	return fake.stmtArgsForCall[i].stmt
}

func (fake *FakeTx) StmtReturns(result1 *sql.Stmt) {
	fake.StmtStub = nil
	fake.stmtReturns = struct {
		result1 *sql.Stmt
	}{result1}
}

func (fake *FakeTx) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	fake.stmtMutex.RLock()
	defer fake.stmtMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTx) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ dbng.Tx = new(FakeTx)
