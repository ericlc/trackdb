package main

// sqlexecutor is an interface to execute
// sql by the different vendors

type SQLExecutor interface {
	Execute(es ExecutableSQL) bool
	//Commit() bool
}
