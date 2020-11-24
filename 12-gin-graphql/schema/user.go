package schema

import (
	"gin-http-server/model"

	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserType",
	Description: "User Type",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "user id",
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "user email",
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var getAllUser = graphql.Field{
	Name:        "GetAllUser",
	Description: "GetAllUser",
	Type:        graphql.NewList(userType),
	Args:        graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		return model.FindAllUsers()
	},
}

var createUser = graphql.Field{
	Name:        "createUser",
	Description: "createUser",
	Type:        userType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		name := p.Args["name"].(string)
		email := p.Args["email"].(string)

		user := &model.User{
			Name:  name,
			Email: email,
		}

		return user, model.CreateUser(user)
	},
}
