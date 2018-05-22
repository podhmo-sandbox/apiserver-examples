package app

import (
	"fmt"
	"time"

	"github.com/podhmo/apiserver-examples/noframework/store"
)

const (
	createdBy         = "test@goa.design"
	accountNameFormat = "Account #%d"
	bottleNameFormat  = "Bottle #%d in account #%d"
)

var (
	createdAt = time.Now()
	kind      = "wine"
	sweetness = 1
	country   = "usa"
	region    = "ca"
	review    = "review"
	rating    = 4
	varietal  = "pinot noir"
	vineyard  = "vineyard"
	vintage   = 2012
	color     = "red"
)

func createAccount(db *store.DB) *store.AccountModel {
	a := db.NewAccount()
	a.Name = fmt.Sprintf(accountNameFormat, a.ID)
	a.CreatedAt = createdAt
	a.CreatedBy = createdBy
	db.SaveAccount(a)
	return &a
}
