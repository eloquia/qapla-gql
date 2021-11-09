package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"qaplagql/graph"
	"qaplagql/graph/generated"
	"qaplagql/graph/model"
	"qaplagql/seequell"
	"qaplagql/services"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
		1. Create inmem/MongoDB datasource
		2. Create services composed of that datasource
		3. Provide services to server
	- - - - - - - - - - - - - - - - - - - - - - - - - - */
	var userService services.UserService
	var projectService services.ProjectService
	var authService services.AuthService
	var meetingService services.MeetingService
	var tagService services.TagService

	datasourceType := os.Args[1]
	switch datasourceType {
	case "inmem":
		log.Fatalf("In-memory not currently supported")
		// log.Printf("Starting in-memory data services")
		// userService, projectService, authService, meetingService, tagService, err = initializeInmemServices()
		// if err != nil {
		// 	log.Fatalf("Unable to initialize inmem services")
		// }
	case "mongodb", "mongo":
		log.Fatalf("MongoDB not currently supported")
	case "postgres", "postgresql":
		log.Printf("Starting Postgres data services")
		userService, projectService, authService, meetingService, tagService, err = initializeSqlServices()
		if err != nil {
			log.Fatalf("Unable to initialize postgres services: %+v", err)
		}
	default:
		log.Fatalf("Unknown datasource type: %+s", datasourceType)
	}

	/* - - - - - - - - -
			Set up server
	- - - - - - - - - - */

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				UserService:    userService,
				ProjectService: projectService,
				AuthService:    authService,
				MeetingService: meetingService,
				TagService:     tagService,
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

/* - - - - - - - - - - - - - - - - -
								MongoDB
- - - - - - - - - - - - - - - - - */

func initializeSqlServices() (services.UserService, services.ProjectService, services.AuthService, services.MeetingService, services.TagService, error) {
	// dbHost := os.Getenv("QAPLA_POSTGRES_DB_HOST")
	dbPort := os.Getenv("QAPLA_POSTGRES_DB_PORT")
	dbUser := os.Getenv("QAPLA_POSTGRES_DB_USER")
	dbName := os.Getenv("QAPLA_POSTGRES_DB_NAME")
	dbPassword := os.Getenv("QAPLA_POSTGRES_DB_PASSWORD")
	sslMode := os.Getenv("QAPLA_POSTGRES_DB_SSLMODE")

	dataSourceName := fmt.Sprintf("port=%s host=db user=%s dbname=%s password=%s sslmode=%s", dbPort, dbUser, dbName, dbPassword, sslMode)
	log.Printf("dataSourceName: %s", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error with ping: %+v", err)
	}

	queries := seequell.New(db)

	userService := seequell.NewUserService(db, queries)
	projectService := seequell.NewProjectService(db, queries)
	authService := seequell.NewAuthService(db, queries)
	meetingService := seequell.NewMeetingService(db, queries)
	tagService := seequell.NewTagService(db, queries)

	return userService, projectService, authService, meetingService, tagService, nil
}

/* - - - - - - - - - - - - - - - - -
							In-memory
- - - - - - - - - - - - - - - - - */

// func initializeInmemServices() (services.UserService, services.ProjectService, services.AuthService, services.MeetingService, services.TagService, error) {
// 	usersMap := initializeUsers()
// 	projectMap := initializeProjects()
// 	projectUserMap := make(map[string][]string)
// 	projectUser := make([]string, 1)
// 	projectUser[0] = "1"
// 	projectUserMap["1"] = projectUser
// 	meetingMap := initializeMeetings()
// 	meetingUserMap := make(map[string][]string)
// 	meetingProjectMap := make(map[string][]string)
// 	tagMap := initializeTags()
// 	meetingItemMap := make(map[string]*model.MeetingItem)
// 	meetingNoteMap := make(map[string]*model.MeetingNote)

// 	userServiceInmem := inmem.NewUserServiceInmem(
// 		usersMap,
// 		projectMap,
// 		projectUserMap,
// 	)

// 	projectServiceInmem := inmem.NewProjectServiceInmem(
// 		usersMap,
// 		projectMap,
// 		projectUserMap,
// 	)

// 	authServiceInmem := &inmem.AuthServiceInmem{
// 		UserSevice: userServiceInmem,
// 	}

// 	meetingServiceInmem := inmem.NewMeetingServiceInmem(
// 		meetingMap,
// 		usersMap,
// 		projectMap,
// 		projectUserMap,
// 		meetingUserMap,
// 		meetingProjectMap,
// 		tagMap,
// 		meetingItemMap,
// 		meetingNoteMap,
// 	)

// 	tagServiceInmem := inmem.NewTagServiceInmem(tagMap)

// 	return userServiceInmem, projectServiceInmem, authServiceInmem, meetingServiceInmem, tagServiceInmem, nil
// }

func initializeUsers() map[string]*model.UserDetails {
	usersMap := make(map[string]*model.UserDetails)
	usersMap["1"] = &model.UserDetails{
		ID:        1,
		FirstName: "Dale",
		LastName:  "Chang",
		Email:     "dale@eloquia.io",
	}

	return usersMap
}

// func initializeProjects() map[string]*model.Project {
// 	projectMap := make(map[string]*model.Project)
// 	projectMap["1"] = &model.Project{
// 		ID:          1,
// 		Name:        "Test Project",
// 		Description: "Test project for testing purposes",
// 		Slug:        "test-project",
// 	}

// 	return projectMap
// }

// func initializeMeetings() map[string]*model.MeetingDetails {
// 	meetingMap := make(map[string]*model.MeetingDetails)

// 	var users []*model.User
// 	dale := &model.User{
// 		ID:        "1",
// 		FirstName: "Dale",
// 		LastName:  "Chang",
// 	}
// 	users = append(users, dale)

// 	var tags []*model.MeetingNoteTag
// 	tag := &model.MeetingNoteTag{
// 		ID:   "1",
// 		Text: "Good",
// 	}
// 	tags = append(tags, tag)
// 	var notes []*model.MeetingNote
// 	note := &model.MeetingNote{
// 		ID:     "1",
// 		About:  dale,
// 		Author: dale,
// 		Text:   "Test status",
// 		Tags:   tags,
// 	}
// 	notes = append(notes, note)
// 	var meetingItems []*model.MeetingItem
// 	attendanceReason := ""
// 	meetingItem := &model.MeetingItem{
// 		ID:                      "1",
// 		Personnel:               dale,
// 		PlannedAttendanceStatus: "Attending",
// 		ActualAttendanceStatus:  "Attending",
// 		AttendanceReason:        &attendanceReason,
// 		Notes:                   notes,
// 	}
// 	meetingItems = append(meetingItems, meetingItem)

// 	meetingMap["1"] = &model.MeetingDetails{
// 		ID:              "1",
// 		Name:            "Test Meeting",
// 		StartTime:       time.Now(),
// 		DurationMinutes: 30,
// 		People:          users,
// 		MeetingItems:    meetingItems,
// 	}

// 	return meetingMap
// }

// func initializeTags() map[string]*model.MeetingNoteTag {
// 	tagMap := make(map[string]*model.MeetingNoteTag)
// 	tagMap["1"] = &model.MeetingNoteTag{
// 		ID:   "1",
// 		Text: "Good",
// 	}
// 	tagMap["2"] = &model.MeetingNoteTag{
// 		ID:   "2",
// 		Text: "Bad",
// 	}
// 	tagMap["3"] = &model.MeetingNoteTag{
// 		ID:   "3",
// 		Text: "Breakthrough",
// 	}
// 	tagMap["4"] = &model.MeetingNoteTag{
// 		ID:   "4",
// 		Text: "Status",
// 	}

// 	return tagMap
// }
