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

	trackSQLs, trackHeaders := findTracks(content)
	
	if trackHeaders != nil {
		// validate and return track headers
		for i, trackHeader := range trackHeaders {
			track, err := checkTrackHeader(trackHeader)
			if err != nil {
				return nil, err
			}
			track.filename = filename
			track.sql = trackSQLs[i+1]
			es = append(es, track)
		}
	} else {
		// checkFilename
		fileProperties, err := checkFilename(filename)
		if err != nil {
			return nil, err
		}
		// return file
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


func findTracks(content *string) ([]string, []string) {

	reg := regexp.MustCompile("(?m)^ *--track.*$")
	return reg.Split(*content, -1), reg.FindAllString(*content, -1)

}

func checkHeaderInit(header string) (*string, error) {

	regStr := "^--track:([0-9]{1,10})$"

	reg := regexp.MustCompile(regStr)
	if reg.MatchString(header) {
		id := reg.FindStringSubmatch(header)[1]
		return &id, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Invalid track header: %s", header))
	}

}

func checkHeaderAtt(att string) (map[string]string, error) {

	var validAttributes = map[string]string{
		"multiThread": "true|false",
		"failOnError": "true|false",
	}

	var attMap = make(map[string]string)

	reg := regexp.MustCompile(":")
	attribute := reg.Split(att, -1)

	if len(attribute) != 2 {
		return nil, errors.New(fmt.Sprintf("Invalid track parameter format: %s", att))
	}

	if _, ok := validAttributes[attribute[0]]; !ok {
		return nil, errors.New(fmt.Sprintf("Invalid track parameter: %s", attribute[0]))
	}

	reg = regexp.MustCompile(validAttributes[attribute[0]])

	if !reg.MatchString(attribute[1]) {
		return nil, errors.New(fmt.Sprintf("Invalid track parameter value: %s", attribute[1]))
	}

	attMap[attribute[0]] = attribute[1]

	return attMap, nil

}

func checkHeaderOp(operation string) (*string, error) {

	reg := regexp.MustCompile("^\\((v|r)\\)$")
	
	if !reg.MatchString(operation) {
		return nil, errors.New(fmt.Sprintf("Invalid track operation: %s", operation))
	}

	op := strings.ToUpper(reg.FindStringSubmatch(operation)[1])

	return &op, nil

}

func checkTrackHeader(header string) (*Track, error) {
	
	var track Track

	// clean last spaces
	reg := regexp.MustCompile(" +$")
	header = reg.ReplaceAllString(header, "")

	reg = regexp.MustCompile(" +")
	params := reg.Split(header, -1)

	if len(params) < 2 {
		return nil, errors.New(fmt.Sprintf("Invalid track header (incomplete): %s", header))
	}

	for k, param := range params {

		if k == 0 {
			id, err := checkHeaderInit(param)
			if err != nil {
				return nil, err
			}
			track.id = *id
		} else if k == len(params)-1 {
			// validate operation
			op, err := checkHeaderOp(param)
			if err != nil {
				return nil, err
			}
			track.operation = *op
		} else {
			attMap, err := checkHeaderAtt(param)
			if err != nil {
				return nil, err
			}
			for key, value := range(attMap) {
				err = track.SetProperty(key, value)
				if err != nil {
					return nil, err
				}
			}

		}

	}

	return &track, nil

}

func trackProperties(trackHeader string) map[string]string {

	var properties = make(map[string]string)
	return properties	
}
