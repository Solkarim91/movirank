## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

Open [http://localhost:3000/movies](http://localhost:3000/movies) with your browser to view the home page. 

You can start editing the page by modifying `pages/movies/index.tsx`. 

## Running the GraphQL server

From your terminal, move into the `api` directory & run the command `go run server.go`, in order to run the server.

Open [http://localhost:8080/graphql](http://localhost:8080/graphql) in your browser to use the GraphQL playground environment. 

## Seeding the database

Go to `server.go` & uncomment lines 61-65:

` 
  // for _, seed := range seeds.SeedMovies() {
	// 	if err := seed.Run(db); err != nil {
	// 		log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
	// 	}
	// }
`

Then restart the GraphQL server by following the previous step. Running the server with these lines uncommented will seed the database. By default these lins should remain commented out unless needing to seed the DB. 
