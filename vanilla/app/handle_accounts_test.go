package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/podhmo/apiserver-examples/vanilla/store"
)

func TestListAccount(t *testing.T) {
	var (
		db = store.NewDB()
		s  = NewServer(db)
	)

	cases := map[string]struct{ NumAccounts int }{
		"empty":  {0},
		"single": {1},
		"many":   {5},
	}

	for k, tc := range cases {
		db.Reset()
		for i := 0; i < tc.NumAccounts; i++ {
			createAccount(db)
		}

		req, err := http.NewRequest(http.MethodGet, "/cellar/accounts", nil)
		if err != nil {
			t.Fatal("unexpected", err)
		}
		w := httptest.NewRecorder()
		s.ServeHTTP(w, req)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status is %d but %d", http.StatusOK, resp.StatusCode)
		}

		var accounts []AccountTiny
		decoder := json.NewDecoder(w.Result().Body)
		if err := decoder.Decode(&accounts); err != nil {
			t.Fatal("unexpected", err)
		}
		if len(accounts) != tc.NumAccounts {
			t.Errorf("%s: invalid number of accounts, expected %d, got %d", k, tc.NumAccounts, len(accounts))
		}
		for i, a := range accounts {
			id := i + 1
			if a.ID != id {
				t.Errorf("%s: invalid account ID at index %d, expected %v, got %v", k, i, id, a.ID)
			}
			if a.Name != fmt.Sprintf(accountNameFormat, id) {
				t.Errorf("%s: invalid account name at index %d, expected %+v, got %+v", k, i, fmt.Sprintf(accountNameFormat, id), a.Name)
			}
		}
	}
}
