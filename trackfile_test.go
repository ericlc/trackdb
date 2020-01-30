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

func TestValidateParam(t *testing.T) {

	var params = map[string]bool{
		"multiThread:true":      true,
		"failOnError:true":      true,
		"multiThread:false":     true,
		"failOnError:false":     true,
		"anotherParameter:true": false,
		"otherParameter:true":   false,
	}

	for key, value := range params {

		if check := validateParam(key); check != value {
			t.Errorf("validate param %v error: expected %v, got %v", key, value, check)
		}

	}

}

func TestCheckHeaderOp(t *testing.T) {

	var operations = map[string]bool{
		"(v)":   true,
		"(r)":   true,
		"v":     false,
		"r":     false,
		"((v))": false,
		"((r))": false,
		`\(v\)`: false,
		`\(r\)`: false,
		"(vvv)": false,
		"(rrr)": false,
		"(a)":   false,
		"(b)":   false,
		"(vr)":  false,
		"(rv)":  false,
		" (v) ": false,
		" (r) ": false,
	}

	for key, value := range operations {

		if check := validateOperation(key); check != value {
			t.Errorf("validate operation %v error: expected %v, got %v", key, value, check)
		}

	}

}

func TestValidateTrackHeader(t *testing.T) {

	var headers = map[string]bool{
		"--track:01":         true,
		"--track:0":          true,
		" --track:01":        false,
		"--track:":           false,
		"--track:0123456789": true,
	}

	for key, value := range headers {

		if check := validateTrackHeader(key); check != value {
			t.Errorf("validate track header %v error: expected %v, got %v", key, value, check)
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
