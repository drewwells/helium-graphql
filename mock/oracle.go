package mock

import (
	"fmt"
	"net/http"

	"github.com/drewwells/helium-graphql/typs"
)

func init() {
	Handle(typs.OracleURL, func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, `{"data":{"timestamp":"2021-02-08T04:03:26.000000Z","price":42,"block":42}}`)
	})
}
