package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ch3lo/graphql-demo/graph/model"
)

func (r *casoResolver) Orden(ctx context.Context, obj *model.Caso) (*model.Orden, error) {
	for _, v := range ordenes {
		if v.ID == obj.OrdenId {
			return v, nil
		}
	}

	return nil, nil
}

func (r *casoResolver) Representante(ctx context.Context, obj *model.Caso) (*model.Representante, error) {
	for _, v := range representantes {
		if v.ID == obj.RepresentanteId {
			return v, nil
		}
	}

	return nil, nil
}

func (r *mutationResolver) CrearCaso(ctx context.Context, input model.NuevoCaso) (*model.Caso, error) {
	caso := &model.Caso{
		ID:              fmt.Sprintf("caso-%d", rand.Intn(1000)),
		FechaCreacion:   time.Now().String(),
		RepresentanteId: input.RepID,
		OrdenId:         input.OrdenID,
	}

	casos = append(casos, caso)

	r.mutex.Lock()
	for _, ch := range r.CasoPublishedChannel {
		ch <- caso
	}
	r.mutex.Unlock()

	return caso, nil
}

func (r *ordenResolver) Producto(ctx context.Context, obj *model.Orden) ([]*model.Producto, error) {
	prods := []*model.Producto{}

	for _, v := range obj.ProductosIds {
		for _, p := range productos {
			if v == p.ID {
				prods = append(prods, p)
			}
		}
	}
	return prods, nil
}

func (r *queryResolver) Casos(ctx context.Context) ([]*model.Caso, error) {
	return casos, nil
}

func (r *subscriptionResolver) CasoCreado(ctx context.Context) (<-chan *model.Caso, error) {
	token := fmt.Sprintf("rdm-%d", rand.Intn(1000))
	casoEvent := make(chan *model.Caso, 1)
	go func() {
		<-ctx.Done()
	}()

	if r.CasoPublishedChannel == nil {
		r.CasoPublishedChannel = map[string]chan *model.Caso{}
	}

	r.mutex.Lock()
	r.CasoPublishedChannel[token] = casoEvent
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.CasoPublishedChannel, token)
		r.mutex.Unlock()
	}()

	return casoEvent, nil
}

// Caso returns CasoResolver implementation.
func (r *Resolver) Caso() CasoResolver { return &casoResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Orden returns OrdenResolver implementation.
func (r *Resolver) Orden() OrdenResolver { return &ordenResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type casoResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type ordenResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
