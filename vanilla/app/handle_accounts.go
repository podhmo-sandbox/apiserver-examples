package app

import (
	"net/http"

	"github.com/podhmo/apiserver-examples/vanilla/app/mapper"
)

// handleAccountList retrieves all the accounts.
func (s *Server) handleAccountList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts := s.db.GetAccounts()
		res := make([]*mapper.AccountTiny, len(accounts))
		for i, account := range accounts {
			res[i] = mapper.ToAccountTiny(&account)
		}

		w.Header().Set("Content-Type", "/application/json")
		if err := NewJSONEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
