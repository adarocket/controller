package helpers

import "adarocket/controller/config"

func Contains(a []config.Node, Ticker, UUID string) bool {
	for _, n := range a {
		if n.UUID == UUID && n.Ticker == Ticker {
			return true
		}
	}
	return false
}
