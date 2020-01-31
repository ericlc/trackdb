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

*/

func TestValidateTrack(t *testing.T) {

	var tracks = map[string]bool{
		"--track:02":                                          false,
		" --track:02":                                         false,
		" --track :02":                                        false,
		" --track :0 2":                                       false,
		"--track:02 multiThread:false (v)":                    true,
		"--track:02 failOnError:false (v)":                    true,
		"--track:02 multiThread:true (v)":                     true,
		"--track:02 failOnError:true (v)":                     true,
		"--track:02 multiThread:false failOnError:false (v)":  true,
		"--track:02 multiThread:true failOnError:true (v)":    true,
		"--track:02 multiThread:test (v)":                     false,
		"--track:02 failOnError:test (v)":                     false,
		"--track:02 multiThread:test2 (v)":                    false,
		"--track:02 failOnError:test2 (v)":                    false,
		"--track:02 multiThread :false failOnError:false (v)": false,
		"--track:02 multiThread :true failOnError:true (v)":   false,
	}

	for key, value := range tracks {

		if check := validateTrack(key); check != value {
			t.Errorf("validate track %v error: expected %v, got %v", key, value, check)
		}

	}

}

func TestCheckHeaderAtt(t *testing.T) {

	var attsCorrect = []string{
		"multiThread:true",
		"failOnError:true",
		"multiThread:false",
		"failOnError:false",
	}
	
	var attsIncorrect = []string{
		"anotherParameter:true",
		"otherParameter:true",
	}


	for _, att := range attsCorrect {

		err := checkHeaderAtt(att)

		if err != nil {
			t.Errorf("header attribute %v error: expected no error, got %v", att, err)
		}

	}
	for _, att := range attsIncorrect {

		err := checkHeaderAtt(att)

		if err == nil {
			t.Errorf("header attribute %v error: expected error, got nil", att, err)
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
		if err != nil {
			t.Errorf("operation %v error: expected error, got nil", operation, err)
		}
	}

}

func TestCheckHeaderInit(t *testing.T) {

	var headers = map[string]bool{
		"--track:01":         true,
		"--track:0":          true,
		" --track:01":        false,
		"--track:":           false,
		"--track:0123456789": true,
	}

	var headersCorrect = []string{
		"--track:01",
		"--track:0",
		"--track:0123456789",
	}

	var headersIncorrect = []string{
		" --track:01",
		"--track:",
	}

	for _, header := range headersCorrect {
		err := checkHeaderInit(header)
		if err != nil {
			t.Errorf("track header %v error: expected no error, got %v", header, err)
		}
	}
	for _, header := range headersCorrect {
		err := checkHeaderInit(header)
		if err != nil {
			t.Errorf("track header %v error: expected no error, got %v", header, err)
		}
	}

}

func TestFindTracks(t *testing.T) {

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

	tracks := findTracks(&content)

	if len(tracks) != total {
		t.Errorf("find tracks error: expected %v tracks, got %v", len(tracks), total)
	}

}

func TestCheckFilename(t *testing.T) {

	filenameCorrect := []string{"V01__create_table.sql",
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

	for _, filename := range(filenameCorrect) {
		_, err := checkFilename(filename)
		if (err != nil) {
			t.Errorf("validate filename %v error: expected no error, got %v", filename, err)
		}		
	}
	
	filenameIncorrect := []string{"V__xxxxx.sql",
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
	for _, filename := range(filenameIncorrect) {
		_, err := checkFilename(filename)
		if (err == nil) {
			t.Errorf("validate filename %v error: expected error, got nil", filename)
		}		
	}

}
