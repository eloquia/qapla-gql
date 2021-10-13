package main

import (
	"log"
	"net/http"
	"os"
	"qaplagql/graph"
	"qaplagql/graph/generated"
	"qaplagql/graph/model"
	"qaplagql/inmem"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:4200", "*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	/* - - - - - - - - - - - - - - - - - - - - - - - - - -
		Using dependency injection:
		1. Create inmem or Postgres datasources
		2. Create services composed of those datasources
		3. Provide services to service
	- - - - - - - - - - - - - - - - - - - - - - - - - - */

	usersMap := make(map[string]*model.User)
	projectMap := make(map[string]*model.Project)
	projectUser := make(map[string][]string)

	userServiceInmem := inmem.NewUserServiceInmem(
		usersMap,
		projectMap,
		projectUser,
	)

	err := userServiceInmem.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	projectServiceInmem := inmem.NewProjectServiceInmem(
		usersMap,
		projectMap,
		projectUser,
	)
	err = projectServiceInmem.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	authServiceInmem := &inmem.AuthServiceInmem{
		UserSevice: userServiceInmem,
	}

	/* - - - - - - - - -
			Set up server
	- - - - - - - - - - */

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				UserService:    userServiceInmem,
				ProjectService: projectServiceInmem,
				AuthService:    authServiceInmem,
			}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "*"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
