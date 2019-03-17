package gqlapollotracing_test

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/realtsiry/gqlgen-contrib/gqlapollotracing"
)

var es graphql.ExecutableSchema

func Example() {
	handler := handler.GraphQL(
		es,
		handler.RequestMiddleware(gqlapollotracing.RequestMiddleware()),
		handler.Tracer(gqlapollotracing.NewTracer()),
	)
	http.Handle("/query", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
