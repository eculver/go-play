package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	rsDB       = "----"
	rsHostPort = "----"
	rsUser     = "----"
	rsPass     = "----"

	s3Bucket    = "----"
	s3AccessKey = "----"
	s3SecretKey = "----"

	TYPE_BOOL  = "boolean"
	TYPE_TEXT  = "varchar"
	TYPE_FLOAT = "float"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("filename required")
	}

	// get filename as first arg
	filename := os.Args[1]
	file, ioErr := os.Open(filename)
	if ioErr != nil { // For read access.
		log.Fatalf("file does not exist: %v", ioErr)
	}

	// setup DB connection
	db, dbErr := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s", rsUser, rsPass, rsHostPort, rsDB))
	if dbErr != nil {
		log.Fatalf("Couldn't connect to DB: %v", dbErr)
	}

	// fmt.Println("connected")

	// table name = filename
	pathParts := strings.Split(filename, "/")
	tableName := strings.Replace(pathParts[len(pathParts)-1], ".csv", "", -1)
	dropTableSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	// fmt.Println(dropTableSQL)
	if _, dropErr := db.Exec(dropTableSQL); dropErr != nil {
		log.Fatalf("DROP TABLE ERROR: %v", dropErr)
	}

	// read csv, generate table
	r := csv.NewReader(file)
	columnNames := []string{}
	columnDefs := make(map[string]string)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("CSV READ error: %v", err)
		}

		// done if we've populated all column defs
		if len(columnDefs) == len(record) {
			break
		}

		// first row is column names
		if len(columnNames) == 0 {
			columnNames = record
			continue
		}

		for idx, v := range record {
			if v != "" {
				if _, err := strconv.ParseBool(v); err == nil {
					columnDefs[columnNames[idx]] = TYPE_BOOL
					continue
				}
				if _, err := strconv.ParseFloat(v, 64); err == nil {
					columnDefs[columnNames[idx]] = TYPE_FLOAT
					continue
				}
				columnDefs[columnNames[idx]] = TYPE_TEXT
			}
		}
	}

	// fmt.Printf("got column names: %s\n", columnNames)
	// fmt.Printf("got column defs: %s\n", columnDefs)

	// generate column parts of the CREATE TABLE SQL
	cols := []string{}
	for colName, colType := range columnDefs {
		cols = append(cols, fmt.Sprintf("%s %s", colName, colType))
	}
	colDefsSQL := strings.Join(cols, ",")
	createTableSQL := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, colDefsSQL)

	// fmt.Println("CREATE TABLE SQL")
	fmt.Println(createTableSQL)

	// fmt.Println("table dropped")

	// issue COPY command
}
