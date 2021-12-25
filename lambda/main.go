package main

import (
	"database/sql"
	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/internal/handler"
	"log"
	"os"

	entsql "entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening database connection: %v", err)
	}

	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))
	defer client.Close()

	lambda.Start(handler.New(client).Handle)
}
