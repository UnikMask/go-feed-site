package main

import (
	"database/sql"

	"fmt"
	"os"
)

func executeFile(fp string, db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    queryString, err := readFile(fp)
    fmt.Println(queryString)
    if err != nil {
        return err
    }
    _, err = tx.Exec(queryString)
    if err != nil {
        return err
    }
    err = tx.Commit()
    if err != nil {
        return err
    }
    return nil
}

func readFile(fp string) (string, error) {
	contents, err := os.ReadFile(fp)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

