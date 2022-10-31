// @/server.go
package main

import (
		"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/common"
		"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/generated"
	
		"github.com/rs/cors"
    resolvers 		"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/resolvers"
    "log"
    "net/http"
    "os"
    "encoding/json"


    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
)
type Ping struct {
  Status int 
  Rssult string
}

const defaultPort = "8081"
func pingHandler(w http.ResponseWriter, r *http.Request) {
  ping := Ping{http.StatusOK, "ok"}

  res, err := json.Marshal(ping)

  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(res)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

		// common.InitDb()
    db, err := common.InitDb()
    if err != nil {
        log.Fatal(err)
    }
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowCredentials: true,
		})

		// // handler = c.Handler(handler)
		// mux := http.NewServeMux()


		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

    customCtx := &common.CustomContext{
        Database: db,
    }
		
		corsSrv := c.Handler(srv)

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		
    http.Handle("/query", common.CreateContext(customCtx, corsSrv))
    http.HandleFunc("/ping", pingHandler)

		// http.Handle("/query", c.Handler(customCtx, srv))
		// http.Handle("/query", c.Handler(srv))


    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
  

}