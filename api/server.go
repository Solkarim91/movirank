package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/solkarim91/movirank/api/app/database"
	"github.com/solkarim91/movirank/api/app/repository"
	"github.com/solkarim91/movirank/api/graph"
)

const defaultPort = "8080"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
    godotenv.Load()
    config := &database.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        Password: os.Getenv("DB_PASS"),
        User:     os.Getenv("DB_USER"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
        DBName:   os.Getenv("DB_NAME"),
    }

    db, err := database.NewConnection(config)

    if err != nil {
        panic(err)
    }

    database.Migrate(db)

    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    repo := repository.NewMovieService(db)

    // UNCOMMENT BELOW TO SEED DATABASE ON SERVER START:
    
    // for _, seed := range seeds.SeedMovies() {
	// 	if err := seed.Run(db); err != nil {
	// 		log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
	// 	}
	// }

    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
        MovieRepository: repo,
    }}))

    http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))    
    http.Handle("/", CorsMiddleware(srv)) 

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
