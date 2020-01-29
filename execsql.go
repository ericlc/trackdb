package main

// there are two kinds of sql code execution
// 1) sql originated by --track
// 2) the full sql file
// they are represented by differents types
// and have the interface ExecutableSQL

// all sql executed have these properties
type SQLProperty struct {
	id        string
	filename  string
	sql       string
	sqlhash   string
	operation string
}

// the full sql file is represented by SQLFile type
type SQLFile struct {
	SQLProperty
}

// the --track has these properties
type TrackProperty struct {
	failOnError bool
	multiThread bool
}

// the --track is represented by Track type
// with SQLProperty and TrackProperty
type Track struct {
	SQLProperty
	TrackProperty
}

// interface ExecutableSQL
type ExecutableSQL interface {
	Id() string
	Filename() string
	SQL() string
	SQLHash() string
	Operation() string
	Property(p string) bool
	IsEqual(es1 ExecutableSQL) bool
}

// SQLFile interace methods
func (f SQLFile) Id() string {
	return f.id
}

func (f SQLFile) Filename() string {
	return f.filename
}

func (f SQLFile) SQL() string {
	return f.sql
}

func (f SQLFile) SQLHash() string {
	return f.sqlhash
}

func (f SQLFile) Operation() string {
	return f.operation
}

func (f SQLFile) Property(p string) bool {
	return false
}

func (f SQLFile) IsEqual(es1 ExecutableSQL) bool {

	sqlFile, ok := es1.(SQLFile)

	if ok {
		if (f.id == sqlFile.id && f.filename == sqlFile.filename && 
			f.sql == sqlFile.sql && f.sqlhash == sqlFile.sqlhash &&
			f.operation == sqlFile.operation) {
			return true
		}
	}

	return false

}

// Track interface methods
func (t Track) Id() string {
	return t.id
}

func (t Track) Filename() string {
	return t.filename
}

func (t Track) SQL() string {
	return t.sql
}

func (t Track) SQLHash() string {
	return t.sqlhash
}

func (t Track) Operation() string {
	return t.operation
}

func (t Track) Property(p string) bool {

	if p == "failOnError" {
		return t.failOnError
	} else if p == "multiThread" {
		return t.multiThread
	}

	return false

}

func (t Track) IsEqual(es1 ExecutableSQL) bool {

	track, ok := es1.(Track)

	if ok {
		if (t.id == track.id && t.filename == track.filename && 
			t.sql == track.sql && t.sqlhash == track.sqlhash &&
			t.operation == track.operation && t.failOnError == track.failOnError &&
			t.multiThread == track.multiThread) {
			return true
		}
	}

	return false

}
