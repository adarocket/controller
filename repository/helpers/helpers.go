package helpers

import (
	"github.com/adarocket/controller/repository/config"
)

func Contains(a []config.Node, Ticker, UUID string) bool {
	for _, n := range a {
		if n.UUID == UUID {
			return true
		}
	}
	return false
}
