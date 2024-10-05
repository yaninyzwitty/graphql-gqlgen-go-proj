package graph

import (
	"database/sql"

	"github.com/yaninyzwitty/graphql-ggqlen-go-proj/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	DB    *sql.DB
}
