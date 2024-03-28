package main

import (
	"database/sql"
	"fmt"
	"log"
    b64 "encoding/base64"
    hex "encoding/hex"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open SQLite database (creates if not exists)
	db, err := sql.Open("sqlite3", "./grafana.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to SQLite database")
        
    // grab the usernames, passwords and salts from the downloaded db
    rows, err := db.Query("select email,password,salt from user")
    if err != nil {
        return
    }
    defer rows.Close()

    for rows.Next() {
        var email string
        var password string
        var salt string
        err = rows.Scan(&email, &password, &salt)
        if err != nil {
            fmt.Println("Failed to extract hashes :(")
            continue
        }

        decoded_hash, _ := hex.DecodeString(password)
        hash64 := b64.StdEncoding.EncodeToString([]byte(decoded_hash))
        salt64 := b64.StdEncoding.EncodeToString([]byte(salt))
        fmt.Println("sha256:10000:" + salt64 + ":" + hash64 + "\n")
    }
}
