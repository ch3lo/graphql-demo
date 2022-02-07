package graph

import (
	"sync"

	"github.com/ch3lo/graphql-demo/graph/model"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CasoPublishedChannel map[string]chan *model.Caso
	mutex                sync.Mutex
}
