package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlGenTutorial/authentication"
	"gqlGenTutorial/connection"
	"gqlGenTutorial/dataaccess/repositories/linkrepositories"
	"gqlGenTutorial/dataaccess/repositories/userrepositories"
	"gqlGenTutorial/graph/generated"
	"gqlGenTutorial/graph/model"
	"gqlGenTutorial/models/linkmodel"
	"gqlGenTutorial/models/usermodel"
	"gqlGenTutorial/settings"
	"gqlGenTutorial/util"
	"log"
	"os"
)

//CreateLink :
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	//Auth
	user := authentication.ForContext(ctx)

	//Si falla mandamos error
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	//Comenzamos a crear link
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
	ok, msg := linkrepositories.CreateLink(&modelLink, db)
	if !ok {
		return &model.Link{}, msg
	}

	//Devolvemos  la info del modelo creado
	return &model.Link{ID: modelLink.IDLink, Title: modelLink.Title, Address: modelLink.Address}, nil

	// panic(fmt.Errorf("not implemented"))
}

//DeleteLink :
func (r *mutationResolver) DeleteLink(ctx context.Context, input model.DeleteLink) (string, error) {
	//Auth
	user := authentication.ForContext(ctx)

	//Si falla mandamos error
	if user == nil {
		return "", fmt.Errorf("access denied")
	}

	var modelLink = linkmodel.LinkModel{}
	modelLink.IDLink = input.ID

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	// Pedimos una conexion a la base de datos POSTGRES
	db := connection.OpenConnection(settings.PostgresDB)

	//Creamos el modelo
	ok, msg := linkrepositories.DeleteLink(&modelLink, db)
	if !ok {
		return "error", msg
	}

	//Devolvemos  la info del modelo creado
	return "success", nil

	// panic(fmt.Errorf("not implemented"))
}

//CreateUser :
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	//Auth
	user := authentication.ForContext(ctx)

	//Si falla mandamos error
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	var userModel = usermodel.UserModel{}
	userModel.UserName = input.Username
	userModel.Password, _ = util.HashPassword(input.Password)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	// Pedimos una conexion a la base de datos POSTGRES
	db := connection.OpenConnection(settings.PostgresDB)

	//Creamos el modelo
	ok, msg := userrepositories.CreateUser(&userModel, db)
	if !ok {
		return &model.User{}, msg
	}

	//Validamos la autenticacion
	token, err := authentication.GenerateToken(userModel.UserName)

	//Devolvemos  la info del modelo creado
	return &model.User{ID: token, Username: userModel.UserName, Password: userModel.Password}, nil

	// panic(fmt.Errorf("not implemented"))
}

//Login :
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	modelUser := usermodel.UserModel{}
	modelUser.UserName = input.Username
	modelUser.Password = input.Password

	// Pedimos una conexion a la base de datos POSTGRES
	db := connection.OpenConnection(settings.PostgresDB)

	//Validamos la autenticacion
	return authentication.Authenticate(modelUser, db)
	// panic(fmt.Errorf("not implemented"))
}

//RefreshToken :
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := authentication.ParseToken(input.Token)

	if err != nil {
		return "", fmt.Errorf("Access Denied")
	}

	token, err := authentication.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

//Links :
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	dummyLink := model.Link{
		Title:   "our dummy link",
		Address: "https://address.org",
		User:    &model.User{Username: "admin"},
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
