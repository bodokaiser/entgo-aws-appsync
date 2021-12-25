package handler

import (
	"context"
	"encoding/json"
	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/internal/resolver"
	"fmt"
	"log"
)

type Action string

const (
	ActionMigrate Action = "migrate"

	ActionTodos      = "todos"
	ActionTodoByID   = "todoById"
	ActionAddTodo    = "addTodo"
	ActionRemoveTodo = "removeTodo"
)

type Event struct {
	Action Action          `json:"action"`
	Input  json.RawMessage `json:"input"`
}

type Handler struct {
	client *ent.Client
}

func New(c *ent.Client) *Handler {
	return &Handler{
		client: c,
	}
}

func (h *Handler) Handle(ctx context.Context, e Event) (interface{}, error) {
	log.Printf("action: %s", e.Action)
	log.Printf("payload: %s", e.Input)

	switch e.Action {
	case ActionMigrate:
		return nil, h.client.Schema.Create(ctx)
	case ActionTodos:
		input := resolver.TodosInput{}
		return resolver.Todos(ctx, h.client, input)
	case ActionTodoByID:
		input := resolver.TodoByIDInput{}
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionTodoByID, err)
		}
		return resolver.TodoByID(ctx, h.client, input)
	case ActionAddTodo:
		input := resolver.AddTodoInput{}
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionAddTodo, err)
		}
		return resolver.AddTodo(ctx, h.client, input)
	case ActionRemoveTodo:
		input := resolver.RemoveTodoInput{}
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionRemoveTodo, err)
		}
		return resolver.RemoveTodo(ctx, h.client, input)
	}

	return nil, fmt.Errorf("invalid action %q", e.Action)
}
