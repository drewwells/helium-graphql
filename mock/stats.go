package mock

import (
	"fmt"
	"net/http"

	"github.com/drewwells/helium-graphql/typs"
)

func init() {
	// https://docs.helium.com/api/blockchain/stats/
	Handle(typs.StatsURL, func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, `{
    "data": {
        "block_times": {
            "last_day": {
                "avg": 70.45921696574226,
                "stddev": 68.83988201467606
            },
            "last_hour": {
                "avg": 66.16666666666667,
                "stddev": 4.492131485499826
            },
            "last_month": {
                "avg": 58.57145439754135,
                "stddev": 54.42394685110655
            },
            "last_week": {
                "avg": 66.77767965559113,
                "stddev": 78.40328359329104
            }
        },
        "challenge_counts": {
            "active": 3226,
            "last_day": 33084
        },
        "counts": {
            "blocks": 644074,
            "challenges": 14189829,
            "cities": 2458,
            "consensus_groups": 16884,
            "countries": 55,
            "hotspots": 14402,
            "transactions": 37813658
        },
        "election_times": {
            "last_day": {
                "avg": 2981.2413793103447,
                "stddev": 1693.238289178722
            },
            "last_hour": {
                "avg": 1997.0,
                "stddev": null
            },
            "last_month": {
                "avg": 2247.3307291666665,
                "stddev": 821.206067231938
            },
            "last_week": {
                "avg": 2608.6293103448274,
                "stddev": 1048.715520504349
            }
        },
        "fees": {
            "last_day": {
                "staking": 611000000,
                "transaction": 35090000
            },
            "last_month": {
                "staking": 9306000000,
                "transaction": 703230000
            },
            "last_week": {
                "staking": 2775000000,
                "transaction": 196880000
            }
        },
        "state_channel_counts": {
            "last_day": {
                "num_dcs": 3160324,
                "num_packets": 532567
            },
            "last_month": {
                "num_dcs": 100340091,
                "num_packets": 16827145
            },
            "last_week": {
                "num_dcs": 23064634,
                "num_packets": 3876268
            }
        },
        "token_supply": 66042860.58752121
    }
}`)
	})
}
