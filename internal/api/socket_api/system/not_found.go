package system

import etp "github.com/elum-utils/go-etp"

// NotFoundHandler handles an incoming event without a registered route.
func NotFoundHandler(*etp.Context) error {
	return etp.ErrRouteNotFound
}
