package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"errors"
	"strings"
)

func OpenSQLFile(filename string) ([]ExecutableSQL, error) {

	fileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	content := string(fileContent)

	es, err := factoryExecutableSQL(filename, &content)

	if err != nil {
		return nil, err
	}

	return es, nil


}

func factoryExecutableSQL(filename string, content *string) ([]ExecutableSQL, error) {
	
	var es []ExecutableSQL

	trackHeaders := findTracks(content)
	
	if trackHeaders != nil {
		// validate and execute tracks
		for _, trackHeader := range trackHeaders {
			test := validateTrack(trackHeader)
			fmt.Println(test)
		}
	} else {
		// validateFilename
		fileProperties, err := validateFilename(filename)
		if err != nil {
			return nil, err
		}
		// execute file
		var sqlFile SQLFile
		sqlFile.id = fileProperties["id"]
		sqlFile.filename = filename
		sqlFile.sql = *content
		sqlFile.sqlhash = ""
		sqlFile.operation = fileProperties["operation"]
		es = append(es, sqlFile)
	}

	return es, nil

}

func validateFilename(filename string) (map[string]string, error) {

	// validate filename and returns file operation and id

	// filename pattern:
	// first letter: V or R + number + underscore + underscore + 
	// name of file + extension .sql
	// example:
	// V02__create_table.sql

	fileProperties := make(map[string]string)

	reg := regexp.MustCompile("(?i)^([v|r])([0-9]{1,10})__[^_](.+)?\\.sql$")

	if (reg.MatchString(filename)) {
		properties := reg.FindAllStringSubmatch(filename, -1)
		fileProperties["operation"] = strings.ToUpper(properties[0][1])
		fileProperties["id"] = properties[0][2]
		return fileProperties, nil
	} else {
		return fileProperties, errors.New("Invalid filename of a file without tracks")
	}

} 


func findTracks(content *string) []string {

	reg := regexp.MustCompile("(?m)^ *--track.*$")
	return reg.FindAllString(*content, -1)

}

func validateTrackHeader(header string) bool {

	reg := regexp.MustCompile("^--track:[0-9]{1,10}")
	return reg.MatchString(header)

}

func validateParam(param string) bool {

	var validAttributes = map[string]string{
		"multiThread": "true|false",
		"failOnError": "true|false",
	}

	reg := regexp.MustCompile(":")
	attribute := reg.Split(param, -1)

	if len(attribute) == 2 {
		if _, ok := validAttributes[attribute[0]]; ok {
			reg = regexp.MustCompile(validAttributes[attribute[0]])
			return reg.MatchString(attribute[1])
		}
	}

	return false

}

func validateOperation(operation string) bool {

	reg := regexp.MustCompile("^\\((v|r)\\)$")
	return reg.MatchString(operation)

}

func validateTrack(track string) bool {

	// clean last spaces
	reg := regexp.MustCompile(" +$")
	track = reg.ReplaceAllString(track, "")

	reg = regexp.MustCompile(" +")
	params := reg.Split(track, -1)

	if len(params) < 2 {
		return false
	}

	for k, param := range params {

		if k == 0 {
			if !validateTrackHeader(param) {
				return false
			}
		} else if k == len(params)-1 {
			// validate operation
			if !validateOperation(param) {
				return false
			}
		} else {
			if !validateParam(param) {
				return false
			}
		}

	}

	return true

}

func trackProperties(trackHeader string) map[string]string {

	var properties = make(map[string]string)
	return properties	
}
