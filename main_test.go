package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MarshalJSON(t *testing.T) {
	u := User{
		Name: "Franky 4 Fingers",
	}
	j, _ := json.Marshal(u)
	assert.Contains(t, string(j), `"firstName":`)

	uv := UserWithViewerPermissions{
		User:   &u,
		Viewer: &u,
	}
	j, _ = json.Marshal(uv)
	assert.Contains(t, string(j), `"viewerCanRead":`)
}
