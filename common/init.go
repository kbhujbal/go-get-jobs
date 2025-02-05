package common

import (
	"sync"

	"github.com/go-resty/resty/v2"
)

var (
	clientOnce sync.Once
	clientO    *resty.Client
)

func GetClient() *resty.Client {
	clientOnce.Do(func() {
		clientO = resty.New()
	})
	return clientO
}

func init() {
	checkDuplicatesComapnies()
	checkAndInitWorkdayCompaniesList()
}
