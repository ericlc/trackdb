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

	trackHeaders := findTrackHeaders(content)
	
	if trackHeaders != nil {
		// validate and execute tracks
		for _, trackHeader := range trackHeaders {
			test := checkTrackHeader(trackHeader)
			fmt.Println(test)
		}
	} else {
		// validateFilename
		fileProperties, err := checkFilename(filename)
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

func checkFilename(filename string) (map[string]string, error) {

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


func findTrackHeaders(content *string) []string {

	reg := regexp.MustCompile("(?m)^ *--track.*$")
	return reg.FindAllString(*content, -1)

}

func checkHeaderInit(header string) error {

	regStr := "^--track:[0-9]{1,10}"

	reg := regexp.MustCompile(regStr)
	if reg.MatchString(header) {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid track header: %s", header))
	}

}

func checkHeaderAtt(att string) error {

	var validAttributes = map[string]string{
		"multiThread": "true|false",
		"failOnError": "true|false",
	}

	reg := regexp.MustCompile(":")
	attribute := reg.Split(att, -1)

	if len(attribute) != 2 {
		return errors.New(fmt.Sprintf("Invalid track parameter format: %s", att))
	}

	if _, ok := validAttributes[attribute[0]]; !ok {
		return errors.New(fmt.Sprintf("Invalid track parameter: %s", attribute[0]))
	}

	reg = regexp.MustCompile(validAttributes[attribute[0]])

	if !reg.MatchString(attribute[1]) {
		return errors.New(fmt.Sprintf("Invalid track parameter value: %s", attribute[1]))
	}

	return nil

}

func checkHeaderOp(operation string) error {

	reg := regexp.MustCompile("^\\((v|r)\\)$")
	
	if !reg.MatchString(operation) {
		return errors.New(fmt.Sprintf("Invalid track operation: %s", operation))
	}

	return nil

}

func checkTrackHeader(header string) error {
	
	var err error

	// clean last spaces
	reg := regexp.MustCompile(" +$")
	header = reg.ReplaceAllString(header, "")

	reg = regexp.MustCompile(" +")
	params := reg.Split(header, -1)

	if len(params) < 2 {
		return errors.New(fmt.Sprintf("Invalid track header (incomplete): %s", header))
	}

	for k, param := range params {

		if k == 0 {
			err = checkHeaderInit(param)
			if err != nil {
				return err
			}
		} else if k == len(params)-1 {
			// validate operation
			err = checkHeaderOp(param)
			if err != nil {
				return err
			}
		} else {
			err = checkHeaderAtt(param)
			if err != nil {
				return err
			}
		}

	}

	return nil

}

func trackProperties(trackHeader string) map[string]string {

	var properties = make(map[string]string)
	return properties	
}
