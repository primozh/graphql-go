package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	database "github.com/primozh/gin-graphql-postgres/db"
	"github.com/primozh/gin-graphql-postgres/graph"
	"github.com/primozh/gin-graphql-postgres/graph/generated"
)

const defaultPort = ":8080"

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	db, err := database.GetDatabase()
	if err != nil {
		panic(err)
	}

	database.RunMigrations(db)
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(defaultPort)
}
