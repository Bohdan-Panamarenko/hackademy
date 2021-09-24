package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			panic("Not enough arguments")
		}
		add(os.Args[2])
	case "list":
		list()
	case "mine":
		if len(os.Args) < 3 {
			panic("Not enough arguments")
		}
		mine(os.Args[2])
	case "reset":
		resetDB()
	default:
		panic("Unknown command")
	}
}

func mine(difficulty string) error {
	diffic, _ := strconv.Atoi(difficulty)

	db, err := NewDB("blockchain.db")
	if err != nil {
		return err
	}

	for {
		block := NewBlock("Mining", db.getPrevHash())

		for !checkHash(block.Hash, diffic) {
			block.Timestamp = time.Now().UnixNano()
			block.setHash()
			fmt.Println(block)
		}

		err = db.Add(block)
		if err != nil {
			return err
		}
	}

	defer db.Close()
	return nil
}

func checkHash(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func add(data string) error {
	db, err := NewDB("blockchain.db")
	if err != nil {
		return err
	}

	block := NewBlock(data, db.getPrevHash())

	err = db.Add(block)
	if err != nil {
		return err
	}

	defer db.Close()
	return nil
}

func list() error {
	db, err := NewDB("blockchain.db")
	if err != nil {
		return err
	}

	records, _ := db.getRecords()
	if err != nil {
		return err
	}

	for _, record := range records {
		fmt.Println(record)
	}

	defer db.Close()
	return nil
}
