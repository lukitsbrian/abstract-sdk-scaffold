package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers blog-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/blog/title", createTitleHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/blog/title", listTitleHandler(cliCtx, "blog")).Methods("GET")
		r.HandleFunc("/blog/title/{key}", getTitleHandler(cliCtx, "blog")).Methods("GET")
		r.HandleFunc("/blog/title", setTitleHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/blog/title", deleteTitleHandler(cliCtx)).Methods("DELETE")

		
}
