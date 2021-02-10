package schema

import (
	"reflect"
	"testing"

	"github.com/drewwells/helium-graphql/mock"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

var tests = []struct {
	Query     string
	Expected  interface{}
	Variables map[string]interface{}
}{
	{
		Query: `query{oracle{price,block}}`,
		Expected: &graphql.Result{
			Data: map[string]interface{}{"oracle": map[string]interface{}{"block": 42, "price": 42}},
		},
	},
	{
		Query: `query{stats{token_supply}}`,
		Expected: &graphql.Result{
			Data: map[string]interface{}{
				"stats": map[string]interface{}{"token_supply": 66042860.58752121}},
		},
	},
}

func TestQL(t *testing.T) {

	_, lis := mock.NewServer(t)
	defer lis.Close()

	for _, test := range tests {
		schema, err := SchemaWithOpts(WithURL("http://" + lis.Addr().String()))
		if err != nil {
			t.Fatal(err)
		}

		params := graphql.Params{
			Schema:         schema,
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
