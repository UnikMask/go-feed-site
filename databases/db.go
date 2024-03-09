package databases

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"os"
)

var db *sql.DB

func ExecuteFile(fp string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	queryString, err := ReadFile(fp)
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

func ExecutePreparedStatement(fp string, args ...string) error {
	contents, err := ReadFile(fp)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare(contents)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(args)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(fp string) (string, error) {
	contents, err := os.ReadFile(fp)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func OpenDatabase(fp string) error {
	var err error
	db, err = sql.Open("sqlite3", fp)
	if err != nil {
		return err
	}
	return nil
}

func CloseDatabase() {
	db.Close()
}
