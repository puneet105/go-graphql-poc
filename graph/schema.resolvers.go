package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-graphql-poc/graph/generated"
	"go-graphql-poc/graph/model"
	"math/rand"
	"strings"
)

func (r *mutationResolver) Add(ctx context.Context, familyTree model.FamilyTreeInput) (*model.FamilyTree, error) {
	tree := &model.FamilyTree{
		ID: fmt.Sprintf("F%d", rand.Int()),
		Name: familyTree.Name,
		Age: familyTree.Age,
		Address: familyTree.Address,
		Contact: familyTree.Contact,
	}
	r.FamilyTrees = append(r.FamilyTrees, tree)
	return tree, nil
}

func (r *mutationResolver) Delete(ctx context.Context, name string) (*model.FamilyTree, error) {
	for index, x := range r.FamilyTrees{
		if x.Name == name{
			r.FamilyTrees = append(r.FamilyTrees[:index], r.FamilyTrees[index+1:]...)
			return x, nil
		}
	}
	return nil,nil
}

func (r *queryResolver) FamilyDetails(ctx context.Context) ([]*model.FamilyTree, error) {
	return r.FamilyTrees, nil
}

func (r *queryResolver) Search(ctx context.Context, name string) (*model.FamilyTree, error) {
	memberName := strings.ToLower(name)
	for _, x := range r.FamilyTrees{
		if strings.Contains(strings.ToLower(x.Name),memberName){
			return x, nil
		}
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
