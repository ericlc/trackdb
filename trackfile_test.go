package main

import (
	"strings"
	"testing"
)

/*

functions:

- OpenSQLFile
- factoryExecutableSQL
- checkFilename
- findTrackHeaders
- checkHeaderInit
- checkHeaderAtt
- checkHeaderOp
- checkHeader
- trackProperties

Test cases

correct___
incorrect___

*/

func TestCheckTrackHeader(t *testing.T) {

	var correctTrackHeaders = []string{
		"--track:02 multiThread:false (v)",
		"--track:02 failOnError:false (v)",
		"--track:02 multiThread:true (v)",
		"--track:02 failOnError:true (v)",
		"--track:02 multiThread:false failOnError:false (v)",
		"--track:02 multiThread:true failOnError:true (v)",
	}

	var incorrectTrackHeaders = []string{
		"--track:02",
		" --track:02",
		" --track :02",
		" --track :0 2",
		"--track:02 multiThread:test (v)",
		"--track:02 failOnError:test (v)",
		"--track:02 multiThread:test2 (v)",
		"--track:02 failOnError:test2 (v)",
		"--track:02 multiThread :false failOnError:false (v)",
		"--track:02 multiThread :true failOnError:true (v)",
	}

	for _, trackHeader := range correctTrackHeaders {
		err := checkTrackHeader(trackHeader)
		if err != nil {
			t.Errorf("track header %v error: expected no error, got %v", trackHeader, err)	
		}
	}

	for _, trackHeader := range incorrectTrackHeaders {
		err := checkTrackHeader(trackHeader)
		if err == nil {
			t.Errorf("track header %v error: expected error, got nil", trackHeader)	
		}
	}

}

func TestCheckHeaderAtt(t *testing.T) {

	var correctAtts = []string{
		"multiThread:true",
		"failOnError:true",
		"multiThread:false",
		"failOnError:false",
	}
	
	var incorrecAtts = []string{
		"anotherParameter:true",
		"otherParameter:true",
	}


	for _, att := range correctAtts {
		err := checkHeaderAtt(att)
		if err != nil {
			t.Errorf("header attribute %v error: expected no error, got %v", att, err)
		}
	}
	for _, att := range incorrecAtts {
		err := checkHeaderAtt(att)
		if err == nil {
			t.Errorf("header attribute %v error: expected error, got nil", att)
		}
	}

}

func TestCheckHeaderOp(t *testing.T) {

	var incorrectOperations = []string{
		"v",
		"r",
		"((v))",
		"((r))",
		`\(v\)`,
		`\(r\)`,
		"(vvv)",
		"(rrr)",
		"(a)",
		"(b)",
		"(vr)",
		"(rv)",
		" (v) ",
		" (r) ",
	}

	var correctOperations = []string{
		"(v)",
		"(r)",
	}

	for _, operation := range correctOperations {
		err := checkHeaderOp(operation)
		if err != nil {
			t.Errorf("operation %v error: expected no error, got %v", operation, err)
		}
	}
	
	for _, operation := range incorrectOperations {
		err := checkHeaderOp(operation)
		if err == nil {
			t.Errorf("operation %v error: expected error, got nil", operation)
		}
	}

}

func TestCheckHeaderInit(t *testing.T) {

	var correctHeaders = []string{
		"--track:01",
		"--track:0",
		"--track:0123456789",
	}

	var incorrectHeaders = []string{
		" --track:01",
		"--track:",
	}

	for _, header := range correctHeaders {
		err := checkHeaderInit(header)
		if err != nil {
			t.Errorf("track header %v error: expected no error, got %v", header, err)
		}
	}
	for _, header := range incorrectHeaders {
		err := checkHeaderInit(header)
		if err == nil {
			t.Errorf("track header %v error: expected no error, got %v", header, err)
		}
	}

}

func TestFindTrackHeaders(t *testing.T) {

	testCases := []struct {
		trackHeader string
		trackReturn bool
	}{
		{"--track:01 (r)\n", true},
		{"--track:02 (r)\n", true},
		{"  --track:02 (r)\n", true},
		{" --track:02 (r)\n", true},
		{"-track:02 (r)\n", false},
		{"  -track:02 (r)\n", false},
		{" --track: 02 (r)\n", true},
		{"---track: 02 (r)\n", false},
		{"--track:\n", true},
		{"--track\n", true},
		{"--track \ntesting ...", true},
	}

	var b strings.Builder
	var total int

	for k, test := range testCases {
		b.WriteString(test.trackHeader)
		if testCases[k].trackReturn {
			total += 1
		}
	}

	content := b.String()

	tracks := findTrackHeaders(&content)

	if len(tracks) != total {
		t.Errorf("find tracks error: expected %v tracks, got %v", len(tracks), total)
	}

}

func TestCheckFilename(t *testing.T) {

	correctFilename := []string{"V01__create_table.sql",
		"v01__create_table.sql",
		"R01__alter__table.sql",
		"r01__alter__table.sql",
		"R0000000000__alter__table.sql",
		"R0123456789__alter__table.sql",
		"r01__a_alter__table.sql",
        "V01__CREATE_TABLE.SQL",
        "R01__ALTER__TABLE.SQL",
        "R01__ALTER__TABLE.SQL",
        "R0000000000__ALTER__TABLE.SQL",
        "R0123456789__ALTER__TABLE.SQL",
        "R01__A_ALTER__TABLE.SQL",
	}

	for _, filename := range(correctFilename) {
		_, err := checkFilename(filename)
		if (err != nil) {
			t.Errorf("validate filename %v error: expected no error, got %v", filename, err)
		}		
	}
	
	incorrectFilename := []string{"V__xxxxx.sql",
		" v01__create_table.sql",
		"    R01__alter__table.sql",
		"   r01__alter__table.sql",
		"a01__alter__table.sql",
		"a01_alter__table.sql",
		"a_alter__table.sql",
		"r01234567890__alter__table.sql",
		"CCY.Autorizador.Serv.exe",
		"r01___alter__table.sql",
		"v01__alter__table.sql2",
        " V01__CREATE_TABLE.SQL",
        "    R01__ALTER__TABLE.SQL",
        "   R01__ALTER__TABLE.SQL",
        "A01__ALTER__TABLE.SQL",
        "A01_ALTER__TABLE.SQL",
        "A_ALTER__TABLE.SQL",
        "R01234567890__ALTER__TABLE.SQL",
        "CCY.AUTORIZADOR.SERV.EXE",
        "R01___ALTER__TABLE.SQL",
        "V01__ALTER__TABLE.SQL2",
	}
	for _, filename := range(incorrectFilename) {
		_, err := checkFilename(filename)
		if (err == nil) {
			t.Errorf("validate filename %v error: expected error, got nil", filename)
		}		
	}

}
