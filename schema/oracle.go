package schema

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/drewwells/helium-graphql/typs"
	"github.com/graphql-go/graphql"
)

type Oracle struct {
	Timestamp *time.Time `json:"timestamp"`
	Price     int        `json:"price"`
	Block     int        `json:"block"`
}

func (c *Config) OracleField() *graphql.Field {
	return &graphql.Field{
		Type:        oracleType,
		Description: "Oracle Price",
		Resolve:     resolveOracle(c),
	}
}

var oracleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Oracle",
	Fields: graphql.Fields{
		"timestamp": &graphql.Field{
			Type: graphql.DateTime,
		},
		"price": &graphql.Field{
			Type: graphql.Int,
		},
		"block": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var root = typs.Root

var timeout = time.Duration(5 * time.Second)

func resolveOracle(c *Config) func(p graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), timeout)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", c.Endpoint(typs.OracleURL), nil)
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var m map[string]Oracle
		err = json.Unmarshal(bs, &m)
		return m["data"], err
	}
}
