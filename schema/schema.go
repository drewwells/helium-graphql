package schema

import (
	"math/rand"

	"github.com/graphql-go/graphql"
)

var TodoList []Todo

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// define custom GraphQL ObjectType `todoType` for our Golang struct `Todo`
// Note that
// - the fields in our todoType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

// define schema, with our rootQuery and rootMutation
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
	// Mutation: rootMutation,
})

type OptFn func(c Config) error

// WithURl sets the remote URL to resolve data with
func WithURL(hostport string) func(c Config) error {
	return func(c Config) error {
		c.hostport = hostport
		return nil
	}
}

type Config struct {
	hostport string
}

func SchemaWithOpts(fns ...OptFn) (graphql.Schema, error) {
	empty := graphql.Schema{}
	c := Config{}
	for _, fn := range fns {
		if err := fn(c); err != nil {
			return empty, err
		}
	}

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation,
	})
}
