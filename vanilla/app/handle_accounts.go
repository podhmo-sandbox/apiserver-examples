package app

import (
	"net/http"

	"github.com/podhmo/apiserver-examples/vanilla/app/mapper"
	"github.com/podhmo/apiserver-examples/vanilla/app/renderer"
)

// handleAccountList retrieves all the accounts.
func (s *Server) handleAccountList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts := s.db.GetAccounts()
		res := make([]*mapper.AccountTiny, len(accounts))
		for i, account := range accounts {
			res[i] = mapper.ToAccountTiny(&account)
		}
		renderer.JSON(w, res)
	}
}
