package resolver

import (
	"context"
	"fmt"
	"strconv"

	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/ent/todo"
)

// TodosInput is the input to the Todos query.
type TodosInput struct{}

// Todos queries all todos.
func Todos(ctx context.Context, client *ent.Client, input TodosInput) ([]*ent.Todo, error) {
	return client.Todo.
		Query().
		All(ctx)
}

// TodoByIDInput is the input to the TodoByID query.
type TodoByIDInput struct {
	ID string `json:"id"`
}

// TodoByID queries a single todo by its id.
func TodoByID(ctx context.Context, client *ent.Client, input TodoByIDInput) (*ent.Todo, error) {
	tid, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed parsing todo id: %w", err)
	}
	return client.Todo.
		Query().
		Where(todo.ID(tid)).
		Only(ctx)
}

// AddTodoInput is the input to the AddTodo mutation.
type AddTodoInput struct {
	Title string `json:"title"`
}

// AddTodoOutput is the output to the AddTodo mutation.
type AddTodoOutput struct {
	Todo *ent.Todo `json:"todo"`
}

// AddTodo adds a todo and returns it.
func AddTodo(ctx context.Context, client *ent.Client, input AddTodoInput) (*AddTodoOutput, error) {
	t, err := client.Todo.
		Create().
		SetTitle(input.Title).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating todo: %w", err)
	}
	return &AddTodoOutput{Todo: t}, nil
}

// RemoveTodoInput is the input to the RemoveTodo mutation.
type RemoveTodoInput struct {
	TodoID string `json:"todoId"`
}

// RemoveTodoOutput is the output to the RemoveTodo mutation.
type RemoveTodoOutput struct {
	Todo *ent.Todo `json:"todo"`
}

// RemoveTodo removes a todo and returns it.
func RemoveTodo(ctx context.Context, client *ent.Client, input RemoveTodoInput) (*RemoveTodoOutput, error) {
	t, err := TodoByID(ctx, client, TodoByIDInput{ID: input.TodoID})
	if err != nil {
		return nil, fmt.Errorf("failed querying todo with id %q: %w", input.TodoID, err)
	}
	err = client.Todo.
		DeleteOne(t).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed deleting todo with id %q: %w", input.TodoID, err)
	}
	return &RemoveTodoOutput{Todo: t}, nil
}
