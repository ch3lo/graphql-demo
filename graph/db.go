package graph

import (
	"time"

	"github.com/ch3lo/graphql-demo/graph/model"
)

var casos []*model.Caso = []*model.Caso{
	{"caso-1", "orden-1", "rep-1", time.Now().String()},
	{"caso-2", "orden-2", "rep-2", time.Now().String()},
}
var representantes []*model.Representante = []*model.Representante{
	{"rep-1", "Homer"},
	{"rep-2", "Bender"},
}
var ordenes []*model.Orden = []*model.Orden{
	{"orden-1", []string{"prod-1", "prod-2"}},
	{"orden-2", []string{"prod-3"}},
}
var productos []*model.Producto = []*model.Producto{
	{"prod-1", 1, true},
	{"prod-2", 2, true},
	{"prod-3", 3, false},
}
