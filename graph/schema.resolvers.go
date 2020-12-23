package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlGenTutorial/connection"
	"gqlGenTutorial/dataaccess/repositories/linkrepositories"
	"gqlGenTutorial/graph/generated"
	"gqlGenTutorial/graph/model"
	"gqlGenTutorial/models/linkmodel"
	"gqlGenTutorial/settings"
	"log"
	"os"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var modelLink = linkmodel.LinkModel{}
	modelLink.Address = input.Address
	modelLink.Title = input.Title

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	// Pedimos una conexion a la base de datos POSTGRES
	db := connection.OpenConnection(settings.PostgresDB)

	//Creamos el modelo
	linkrepositories.CreateLink(&modelLink, db)

	//Devolvemos  la info del modelo creado
	return &model.Link{ID: modelLink.IDLink, Title: modelLink.Title, Address: modelLink.Address}, nil

	// panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	dummyLink := model.Link{
		Title:   "our dummy link",
		Address: "https://address.org",
		User:    &model.User{Name: "admin"},
	}
	links = append(links, &dummyLink)
	return links, nil

	// panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
