package adapter

import (
	"testing"
)

func TestAuthorizationLaunch(t *testing.T) {
	tests := []struct {
		authorization string
		launch        string
		ok            bool
	}{
		{
			authorization: " launch-params ",
			launch:        "launch-params",
			ok:            true,
		},
		{authorization: "", ok: false},
		{authorization: "   ", ok: false},
		{authorization: "launch params", launch: "launch params", ok: true},
	}

	for _, test := range tests {
		launch, ok := authorizationLaunch(test.authorization)
		if ok != test.ok || launch != test.launch {
			t.Errorf(
				"authorizationLaunch(%q) = (%q, %t), want (%q, %t)",
				test.authorization,
				launch,
				ok,
				test.launch,
				test.ok,
			)
		}
	}
}
