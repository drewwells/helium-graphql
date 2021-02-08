package main

import (
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/schema"
	"github.com/graphql-go/graphql/testutil"
)

var tests = []struct {
	Query     string
	Schema    graphql.Schema
	Expected  interface{}
	Variables map[string]interface{}
}{
	// {
	// 	Schema: TodoSchema,
	// 	Query:  `query{todoList{id,text,done}}`,
	// 	Expected: &graphql.Result{
	// 		Data: map[string]interface{}{
	// 			"todoList": []interface{}{},
	// 		},
	// 	},
	// },
	{
		Schema: schema.Schema,
		Query:  `query{oracle{price,block}}`,
		Expected: &graphql.Result{
			Data: map[string]interface{}{"oracle": map[string]interface{}{"block": 710040, "price": 326680500}},
		},
	},
}

func TestQL(t *testing.T) {
	for _, test := range tests {
		params := graphql.Params{
			Schema:         test.Schema,
			RequestString:  test.Query,
			VariableValues: test.Variables,
		}
		result := graphql.Do(params)
		if len(result.Errors) > 0 {
			t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
		}
		if !reflect.DeepEqual(result, test.Expected) {
			t.Logf("\ngot:  %#v\nwanted: %#v", result, test.Expected)
			t.Fatalf("wrong result, query: %v, graphql result diff: %v", test.Query, testutil.Diff(test.Expected, result))
		}

	}
}
