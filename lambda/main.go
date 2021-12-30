package main

import (
	"database/sql"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/jackc/pgx/v4/stdlib"

	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/internal/handler"
)

func main() {
	// open the daatabase connection using the pgx driver
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening database connection: %v", err)
	}

	// initiate the ent database client for the Postgres database
	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))
	defer client.Close()

	// register our event handler to lissten on Lambda events
	lambda.Start(handler.New(client).Handle)
}
