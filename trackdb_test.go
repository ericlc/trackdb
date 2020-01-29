package main

import (
	"flag"
	"strings"
	"testing"
)

func TestFlagParse(t *testing.T) {

	parameters := [...]string{
		"host", "port", "dbname", "dbtype",
		"u", "p", "filename", "fencoding", "domain",
	}

	var mainFlags []string

	var missingMain []string
	var missingTest []string

	check := func(flag *flag.Flag) {

		mainFlags = append(mainFlags, flag.Name)

		for _, value := range parameters {
			if flag.Name == value {
				return
			}
		}

		if !strings.HasPrefix(flag.Name, "test") {
			missingTest = append(missingTest, flag.Name)
		}
	}

	commandLine := flagParse()
	commandLine.VisitAll(check)

	for _, parameter := range parameters {

		found := false

		for _, main := range mainFlags {

			if !strings.HasPrefix(main, "test") {
				if parameter == main {
					found = true
					break
				}
			}
		}

		if !found {
			missingMain = append(missingMain, parameter)
		}
	}

	if len(missingTest) == 0 && len(missingMain) == 0 {
		t.Logf("flag parse success: all flag(s) in main file match flag(s) in test file")
	}

	if len(missingMain) != 0 {
		t.Errorf("flag parse error: missing flag(s) in main file: %v", strings.Join(missingMain, ", "))
	}

	if len(missingTest) != 0 {
		t.Errorf("flag parse error: missing flag(s) in test file: %v", strings.Join(missingTest, ", "))
	}

}

func TestValidateCli(t *testing.T) {

	var commandLine = flag.NewFlagSet("test", flag.ExitOnError)

	commandLine.String("f", "test.sql", "filename")

	missing, err := validateCli(commandLine)

	if err == nil {
		t.Logf("cli validation success: no missing parameter(s) as expected")
	} else {
		t.Errorf("cli validation error: expected no missing parameter(s), got: missing parameter(s): %v", strings.Join(missing, ", "))
	}

	commandLine.String("host", "", "hostname of database server or ip address")
	missing, err = validateCli(commandLine)

	if err != nil {
		t.Logf("cli validation success: missing parameter(s) as expected")
	} else {
		t.Errorf("cli validation error: expected at least one missing parameter, got: no missing parameter(s)")
	}

}
