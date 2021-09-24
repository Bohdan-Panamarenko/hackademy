// Package trades provides an SQLite based trades database.
package main

// Your main or test packages require this import so
// the sql package is properly initialized.
// _ "github.com/mattn/go-sqlite3"

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	insertSQL = `
INSERT INTO blockchain (
	time, data, prevblockhash, hash
) VALUES (
	?, ?, ?, ?
)
`

	schemaSQL = `
CREATE TABLE IF NOT EXISTS blockchain (
    time INTEGER,
    data VARCHAR(256),
    prevblockhash CHAR(64),
    hash CHAR(64)
);
`
)

// DB is a database of stock trades.
type DB struct {
	sql  *sql.DB
	stmt *sql.Stmt
}

// NewDB constructs a Trades value for managing stock trades in a
// SQLite database. This API is not thread safe.
func NewDB(dbFile string) (*DB, error) {
	sqlDB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	// if _, err = sqlDB.Exec(schemaSQL); err != nil {
	// 	return nil, err
	// }

	stmt, err := sqlDB.Prepare(insertSQL)
	if err != nil {
		return nil, err
	}

	db := DB{
		sql:  sqlDB,
		stmt: stmt,
	}
	return &db, nil
}

func (db *DB) getPrevHash() string {
	row := db.sql.QueryRow(`
	SELECT hash FROM blockchain ORDER BY time DESC LIMIT 1 
	`)
	var prevHash string
	row.Scan(&prevHash)

	return prevHash
}

// Add stores a block into the buffer.
func (db *DB) Add(block *Block) error {

	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(db.stmt).Exec(block.Timestamp, block.Data, block.PrevBlockHash, block.Hash)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Flush inserts pending trades into the database.
// func (db *DB) Flush() error {
// 	tx, err := db.sql.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	for _, trade := range db.buffer {
// 		_, err := tx.Stmt(db.stmt).Exec(trade.Time, trade.Symbol, trade.Price, trade.IsBuy)
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 	}

// 	db.buffer = db.buffer[:0]
// 	return tx.Commit()
// }

func (db *DB) Close() error {
	defer func() {
		db.stmt.Close()
		db.sql.Close()
	}()

	return nil
}

func resetDB() error {
	err := os.Remove("blockchain.db")
	if err != nil {
		return err
	}

	sqlDB, err := sql.Open("sqlite3", "blockchain.db")
	if err != nil {
		return err
	}

	if _, err = sqlDB.Exec(schemaSQL); err != nil {
		return err
	}

	stmt, err := sqlDB.Prepare(insertSQL)
	if err != nil {
		return err
	}

	tx, err := sqlDB.Begin()
	if err != nil {
		return err
	}

	block := NewGenesisBlock()

	_, err = tx.Stmt(stmt).Exec(block.Timestamp, block.Data, block.PrevBlockHash, block.Hash)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (db *DB) getRecords() ([]Block, error) {

	rows, err := db.sql.Query(`
	SELECT * FROM blockchain ORDER BY time DESC
	`)
	if err != nil {
		return nil, err
	}

	records := make([]Block, 0)
	for rows.Next() {
		var record Block
		if err := rows.Scan(&record.Timestamp, &record.Data, &record.PrevBlockHash, &record.Hash); err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}

	defer rows.Close()

	return records, nil
}
