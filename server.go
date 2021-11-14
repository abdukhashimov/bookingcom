package main

import (
	"fmt"
	"log"
	"os"

	"abdukhashimov/mybron.uz/graph"
	"abdukhashimov/mybron.uz/graph/generated"
	"abdukhashimov/mybron.uz/logger"
	"abdukhashimov/mybron.uz/services"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

// Defining the Graphql handler
func graphqlHandler(queries *sqlc.Queries) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file

	h := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{
				Log:      logger.New("info", "graphql-test"),
				Services: services.NewServices(queries),
			},
		}),
	)

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
	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"booking",
		"disable",
	)

	postgresConn, err := sql.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(postgresConn)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up Gin
	r := gin.New()
	r.POST("/query", graphqlHandler(queries))
	r.GET("/", playgroundHandler())
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	r.Run()
}
