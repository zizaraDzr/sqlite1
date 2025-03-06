package main

import (
	"database/sql"
	"log"
	"sqlite/grpcapi"
	"sqlite/jsonapi"
	"sync"

	"github.com/alexflint/go-arg"
	_ "modernc.org/sqlite"
)

var args struct {
	DbPath   string `arg:"env:SQLITE_DB"`
	BindJson string `arg:"env:SQLITET_BIND_JSON"`
	BindGrpc string `arg:"env:SQLITE_BIND_GRPC"`
}

func main() {
	arg.MustParse(&args)

	if args.DbPath == "" {
		args.DbPath = "list.db"
	}
	if args.BindJson == "" {
		args.BindJson = ":3000"
	}
	if args.BindGrpc == "" {
		args.BindGrpc = ":8081"
	}

	log.Printf("using database '%v'\n", args.DbPath)
	db, err := sql.Open("sqlite", args.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// mdb.TryCreate(db)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Printf("starting JSON API server...\n")
		jsonapi.Serve(db, args.BindJson)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Printf("starting gRPC API server...\n")
		grpcapi.Serve(db, args.BindGrpc)
		wg.Done()
	}()

	wg.Wait()
}
