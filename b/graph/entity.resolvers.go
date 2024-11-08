package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"main/graph/model"
)

// FindSimplePostByID is the resolver for the findSimplePostByID field.
func (r *entityResolver) FindSimplePostByID(ctx context.Context, id string) (*model.SimplePost, error) {
	panic(fmt.Errorf("not implemented: FindSimplePostByID - findSimplePostByID"))
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
