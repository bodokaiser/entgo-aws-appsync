package handler

import (
	"context"
	"encoding/json"
	"entgo-aws-appsync/ent"
	"fmt"
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
	switch e.Action {
	case ActionMigrate:
		return nil, h.client.Schema.Create(ctx)
	case ActionTodos:
		return nil, nil
	case ActionTodoByID:
		return nil, nil
	case ActionAddTodo:
		return nil, nil
	case ActionRemoveTodo:
		return nil, nil
	}

	return nil, fmt.Errorf("invalid action %q", e.Action)
}
