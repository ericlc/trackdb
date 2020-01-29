package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func validateCli(flagSet *flag.FlagSet) (missing []string, err error) {

	check := func(flag *flag.Flag) {
		if flag.Value.String() == "" {
			missing = append(missing, flag.Name)
		}
	}

	flagSet.VisitAll(check)

	if len(missing) > 0 {
		return missing, errors.New("Missing required parameter(s)")
	}

	return missing, nil

}

func flagParse() *flag.FlagSet {

	flag.String("host", "", "hostname of database server or ip address")
	flag.String("port", "", "port of database connection")
	flag.String("dbname", "", "database name")
	flag.String("dbtype", "", "database type: mysql, postgres, db2, sqlserver, oracle")
	flag.String("u", "", "username of database connection")
	flag.String("p", "", "password of database connection")
	flag.String("filename", "", "filename")
	flag.String("fencoding", "UTF-8", "file encoding: UTF-8, ISO8859-1 ...")
	flag.Bool("domain", false, "use windows user to connect to SQL Server")

	flag.Parse()

	return flag.CommandLine

}

func main() {

	commandLine := flagParse()

	missing, err := validateCli(commandLine)

	if err != nil {
		fmt.Printf("%v: %v\n", err, strings.Join(missing, ", "))
		fmt.Printf("Usage of %v:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

}
