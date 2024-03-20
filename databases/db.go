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

func ExecutePreparedStatement(fp string, args ...any) error {
	contents, err := ReadFile(fp)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare(contents)
	if err != nil {
		return err
	}

	args_any := []any{}
	for _, arg := range args {
		args_any = append(args_any, arg)
	}
	_, err = stmt.Exec(args_any...)
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

func QueryRow(fp string, stmtArgs []any, scanElements []any) (bool, error) {
	contents, err := ReadFile(fp)
	if err != nil {
		return false, err
	}

	stmt, err := db.Prepare(contents)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	query := stmt.QueryRow(stmtArgs...)
	err = query.Scan(scanElements...)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

type Rows struct {
	c *sql.Rows
}

func (r Rows) ScanNext(writes ...any) (bool, error) {
	if !r.c.Next() {
		return false, r.c.Err()
	}
	err := r.c.Scan(writes...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r Rows) Close() {
	r.c.Close()
}

func Query(fp string, stmtArgs ...any) (Rows, error) {
	contents, err := ReadFile(fp)
	if err != nil {
		return Rows{}, err
	}

	stmt, err := db.Prepare(contents)
	if err != nil {
		return Rows{}, err
	}
	query, err := stmt.Query(stmtArgs...)
	if err != nil {
		return Rows{}, err
	}

	return Rows{c: query}, nil
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
