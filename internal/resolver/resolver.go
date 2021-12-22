package resolver

import (
	"context"

	"entgo-aws-appsync/ent"
)

type TodosInput struct{}

func Todos(ctx context.Context, client *ent.Client, input TodosInput) ([]ent.Todo, error) {
	return nil, nil
}

type TodoByIDInput struct {
	TodoID string `json:"todoId"`
}

func TodoByID(ctx context.Context, client *ent.Client, input TodoByIDInput) (*ent.Todo, error) {
	return nil, nil
}

type AddTodoInput struct {
	Title string `json:"title"`
}

type AddTodoOutput struct {
	Todo ent.Todo `json:"todo"`
}

func AddTodo(ctx context.Context, client *ent.Client, input AddTodoInput) (*AddTodoOutput, error) {
	return nil, nil
}

type RemoveTodoInput struct {
	TodoID string `json:"todoId"`
}

type RemoveTodoOutput struct {
	Todo ent.Todo `json:"todo"`
}

func RemoveTodo(ctx context.Context, client *ent.Client, input RemoveTodoInput) (*RemoveTodoOutput, error) {
	return nil, nil
}
