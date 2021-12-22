package main

import (
	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/internal/handler"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	lambda.Start(handler.New(client))
}
