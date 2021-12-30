package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/internal/resolver"
)

// Action specifies the event type.
type Action string

// List of supported event actions.
const (
	ActionMigrate Action = "migrate"

	ActionTodos      = "todos"
	ActionTodoByID   = "todoById"
	ActionAddTodo    = "addTodo"
	ActionRemoveTodo = "removeTodo"
)

// Event is the argument of the event handler.
type Event struct {
	Action Action          `json:"action"`
	Input  json.RawMessage `json:"input"`
}

// Handler handles supported events.
type Handler struct {
	client *ent.Client
}

// Returns a new event handler.
func New(c *ent.Client) *Handler {
	return &Handler{
		client: c,
	}
}

// Handle implements the event handling by action.
func (h *Handler) Handle(ctx context.Context, e Event) (interface{}, error) {
	log.Printf("action %s with payload %s\n", e.Action, e.Input)

	switch e.Action {
	case ActionMigrate:
		return nil, h.client.Schema.Create(ctx)
	case ActionTodos:
		var input resolver.TodosInput
		return resolver.Todos(ctx, h.client, input)
	case ActionTodoByID:
		var input resolver.TodoByIDInput
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionTodoByID, err)
		}
		return resolver.TodoByID(ctx, h.client, input)
	case ActionAddTodo:
		var input resolver.AddTodoInput
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionAddTodo, err)
		}
		return resolver.AddTodo(ctx, h.client, input)
	case ActionRemoveTodo:
		var input resolver.RemoveTodoInput
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionRemoveTodo, err)
		}
		return resolver.RemoveTodo(ctx, h.client, input)
	}

	return nil, fmt.Errorf("invalid action %q", e.Action)
}
