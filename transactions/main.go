package main

import (
	"database/sql"
	"fmt"
)

func main() {
	var err error

	defer func() {
		if err != nil {
			fmt.Printf("aww.. err is not nil: %v", err)
		}
	}()

	err = fmt.Errorf("umm....2)")

	return

}

// For the sake of this exercise, we assume conn is
// already initialized.
var conn *sql.DB

// CreateDBV1 creates a database named dbname, a user with
// name dbuser and password dbpass, then grants privileges on
// database to user.
//
//  Returns error if any of these failed.
func CreateDBV1(dbname, dbuser, dbpass string) error {
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s CHARSET utf8;", dbname))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create database: %v", err)
	}

	_, err = conn.Exec(fmt.Sprintf("CREATE USER '%s' IDENTIFIED BY '%s';", dbuser, dbpass))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create user %q: %v", dbuser, err)
	}

	_, err = conn.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%';", dbuser, dbpass))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to grant privileges on %q to %q: %v", dbname, dbuser, err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("committing transaction failed: %v", err)
	}

	return nil
}

// CreateDBV2 meh
func CreateDBV2(dbname, dbuser, dbpass string) error {
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	queries := []string{
		fmt.Sprintf("CREATE DATABASE %s CHARSET utf8;", dbname),
		fmt.Sprintf("CREATE USER '%s' IDENTIFIED BY '%s';", dbuser, dbpass),
		fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%';", dbuser, dbpass),
	}

	for _, q := range queries {
		_, err = conn.Exec(q)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create database: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("committing transaction failed: %v", err)
	}

	return nil
}

// CreateDBV3 asd
func CreateDBV3(dbname, dbuser, dbpass string) error {
	var (
		err error
		tx  *sql.Tx
	)

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	tx, err = conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s CHARSET utf8;", dbname))
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}

	_, err = conn.Exec(fmt.Sprintf("CREATE USER '%s' IDENTIFIED BY '%s';", dbuser, dbpass))
	if err != nil {
		return fmt.Errorf("failed to create user %q: %v", dbuser, err)
	}

	_, err = conn.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%';", dbuser, dbpass))
	if err != nil {
		return fmt.Errorf("failed to grant privileges on %q to %q: %v", dbname, dbuser, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("committing transaction failed: %v", err)
	}

	return nil
}
