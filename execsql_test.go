package main

import (
	"testing"
)

func TestIsEqual(t *testing.T) {

	var track1 Track
	track1.id = "01"
	track1.filename = "test.sql"
	track1.sql = "create table"
	track1.sqlhash = "dda60e10809a1d64b79c9ef05ed53e1b"
	track1.operation = "v"
	track1.failOnError = false
	track1.multiThread = false

	// id different
	var track2 Track
	track2.id = "02"
	track2.filename = "test.sql"
	track2.sql = "create table"
	track2.sqlhash = "dda60e10809a1d64b79c9ef05ed53e1b"
	track2.operation = "v"
	track2.failOnError = false
	track2.multiThread = false
	
	if (track1.IsEqual(track2)) {
		t.Errorf("IsEqualEs error: expected false, got true")
	}

	// filename different
	track2.id = "01"
	track2.filename = "trackdb.sql"
	
	if (track1.IsEqual(track2)) {
		t.Errorf("IsEqualEs error: expected false, got true")
	}

	// sql different
	track2.filename = "test.sql"
	track2.sql = "drop table"

	if (track1.IsEqual(track2)) {
		t.Errorf("IsEqualEs error: expected false, got true")
	}

	// sqlhash different
	track2.sql = "trackdb.sql"
	track2.sqlhash = "60e10809a1d64b79c9ef05ed53e1b"

	if (track1.IsEqual(track2)) {
		t.Errorf("IsEqualEs error: expected false, got true")
	}

	// operation different
	track2.sql = "trackdb.sql"
	track2.sqlhash = "60e10809a1d64b79c9ef05ed53e1b"

	if (track1.IsEqual(track2)) {
		t.Errorf("IsEqualEs error: expected false, got true")
	}


}
