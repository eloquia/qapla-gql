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

	usersMap := initializeUsers()
	projectMap := initializeProjects()
	projectUserMap := make(map[string][]string)
	projectUser := make([]string, 1)
	projectUser[0] = "1"
	projectUserMap["1"] = projectUser
	meetingMap := make(map[string]*model.Meeting)
	meetingUserMap := make(map[string][]string)
	meetingProjectMap := make(map[string][]string)

	userServiceInmem := inmem.NewUserServiceInmem(
		usersMap,
		projectMap,
		projectUserMap,
	)

	projectServiceInmem := inmem.NewProjectServiceInmem(
		usersMap,
		projectMap,
		projectUserMap,
	)

	authServiceInmem := &inmem.AuthServiceInmem{
		UserSevice: userServiceInmem,
	}

	meetingServiceInmem := inmem.NewMeetingServiceInmem(
		meetingMap,
		usersMap,
		projectMap,
		projectUserMap,
		meetingUserMap,
		meetingProjectMap,
	)

	/* - - - - - - - - -
			Set up server
	- - - - - - - - - - */

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				UserService:    userServiceInmem,
				ProjectService: projectServiceInmem,
				AuthService:    authServiceInmem,
				MeetingService: meetingServiceInmem,
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

func initializeUsers() map[string]*model.User {
	usersMap := make(map[string]*model.User)
	usersMap["1"] = &model.User{
		ID:         "1",
		FirstName:  "Dale",
		LastName:   "Chang",
		GoesBy:     "Dale",
		MiddleName: "",
		Email:      "dale@eloquia.io",
		Password:   "notagoodpassword",
		Ethnicity:  "East Asian",
		Position:   "Software Engineer",
	}

	return usersMap
}

func initializeProjects() map[string]*model.Project {
	projectMap := make(map[string]*model.Project)
	projectMap["1"] = &model.Project{
		ID:          "1",
		Name:        "Test Project",
		Description: "Test project for testing purposes",
		Slug:        "test-project",
	}

	return projectMap
}
