package main

import (
    "database/sql"
	_"github.com/mattn/go-sqlite3"
    "./storage"
    "./path"
    "fmt"
    "io/ioutil"
    "flag"
)


func migrate(db *sql.DB, xml string) {
    user := path.GetUser(xml)

    user_id := storage.AddUser(db, user.Name)

    fmt.Printf("sqliteId: %d\n", user_id)
    fmt.Printf("username: %s\n", user.Name)
    fmt.Printf("malId: %d\n", user.MalUserId)
}


func main() {
    var outputFilename string
    flag.StringVar(&outputFilename, "output", "output.sqlite", "Output file")

    flag.Parse()

    args := flag.Args()
    if len(args) < 1 {
        panic("No input file specified")
    }

    inputFilename := args[0]
    xmlBytes, err := ioutil.ReadFile(inputFilename)
    if err != nil {
        panic(err)
    }
    xml := string(xmlBytes)

	db := storage.InitDB("output.sqlite")
    defer db.Close()
    storage.CreateSchema(db)



    migrate(db, xml)

    fmt.Printf("Done.\n")
}

