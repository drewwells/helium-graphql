package schema

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/drewwells/helium-graphql/typs"
	"github.com/graphql-go/graphql"
)

// Stats https://docs.helium.com/api/blockchain/stats/
type Stats struct {
	// BlockTimes      StatsAggregate `json:"block_times,omitempty"`
	// ChallengeCounts struct {
	// 	Active  int
	// 	LastDay int `json:"last_day,omitempty"`
	// } `json:"challenge_counts,omitempty"`
	// Counts struct {
	// 	Blocks          int
	// 	Challenges      int
	// 	Cities          int
	// 	ConsensusGroups int `json:"consensus_groups,omitempty"`
	// 	Countries       int
	// 	HotSpots        int
	// 	Transactions    int
	// }
	// ElectionTimes StatsAggregate `json:"election_times,omitempty"`
	// Fees          struct {
	// 	LastDay   StatsFeesRate `json:"last_day"`
	// 	LastMonth StatsFeesRate `json:"last_month"`
	// 	LastWeek  StatsFeesRate `json:"last_week"`
	// }
	// StateChannelCounts struct {
	// 	LastDay   StateChannelCountsRate `json:"last_day,omitempty"`
	// 	LastMonth StateChannelCountsRate `json:"last_month,omitempty"`
	// 	LastWeek  StateChannelCountsRate `json:"last_week,omitempty"`
	// } `json:"state_channel_counts,omitempty"`
	TokenSupply float64 `json:"token_supply"`
}

type StateChannelCountsRate struct {
	NumDCs     int `json:"num_dcs,omitempty"`
	NumPackets int `json:"num_packets,omitempty"`
}

type StatsFeesRate struct {
	Staking     int
	Transaction int
}

// StatsRates tracks average and standard deviation
type StatsRate struct {
	Avg    float64 `json:"avg,omitempty"`
	Stddev float64 `json:"stddev,omitempty"`
}

type StatsAggregate struct {
	LastDay   StatsRate `json:"last_day,omitempty"`
	LastHour  StatsRate `json:"last_hour,omitempty"`
	LastMonth StatsRate `json:"last_month,omitempty"`
}

func StatsField(c *Config) *graphql.Field {
	return &graphql.Field{
		Type:        statsType,
		Description: "",
		Resolve:     resolveStats(c),
	}
}

var statsType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Stats",
	Fields: graphql.BindFields(Stats{}),
})

func resolveStats(c *Config) func(p graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), c.APITimeout)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", c.Endpoint(typs.StatsURL), nil)
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

		var m map[string]Stats
		err = json.Unmarshal(bs, &m)
		return m["data"], err
	}
}
