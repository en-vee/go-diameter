package sm

import "strings"

func parseDialNetwork(addr string) (network, host string) {
	switch {
	case strings.HasPrefix(addr, "sctp://"):
		return "sctp", strings.TrimPrefix(addr, "sctp://")
	case strings.HasPrefix(addr, "tcp://"):
		return "tcp", strings.TrimPrefix(addr, "tcp://")
	default:
		return "tcp", addr
	}
}
