package app

import (
	"net/http"
	"time"

	"github.com/podhmo/apiserver-examples/vanilla/renderer"
	"github.com/podhmo/apiserver-examples/vanilla/store"
)

// ToAccount builds an account media type from an account model.
func ToAccount(account *store.AccountModel) *Account {
	return &Account{
		ID:        account.ID,
		Name:      account.Name,
		CreatedAt: account.CreatedAt,
		CreatedBy: account.CreatedBy,
	}
}

// ToAccountTiny builds an account media type with tiny view from an account model.
func ToAccountTiny(account *store.AccountModel) *AccountTiny {
	return &AccountTiny{
		ID:   account.ID,
		Name: account.Name,
	}
}

// Account :
type Account struct {
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	CreatedBy string    `form:"created_by" json:"created_by" xml:"created_by"`
	ID        int       `form:"id" json:"id" xml:"id"`
	Name      string    `form:"name" json:"name" xml:"name"`
}

// AccountTiny :
type AccountTiny struct {
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// handleAccountList retrieves all the accounts.
func (s *Server) handleAccountList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts := s.db.GetAccounts()
		res := make([]*AccountTiny, len(accounts))
		for i, account := range accounts {
			res[i] = ToAccountTiny(&account)
		}
		renderer.JSON(w, res)
	}
}
