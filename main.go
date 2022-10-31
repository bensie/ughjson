package main

import (
	"encoding/json"
	"strings"
)

type User struct {
	Name string `json:"name"`
}

func (u User) firstName() string {
	return strings.Split(u.Name, " ")[0]
}

func (u User) viewerCanRead(viewer *User) bool {
	if viewer.Name == u.Name {
		return true
	}
	return false
}

func (u User) MarshalJSON() ([]byte, error) {
	type alias User
	return json.Marshal(struct {
		alias
		FirstName string `json:"firstName"`
	}{
		alias:     alias(u),
		FirstName: u.firstName(),
	})
}

type UserWithViewerPermissions struct {
	*User
	Viewer *User
}

func (u UserWithViewerPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		User
		ViewerCanRead bool `json:"viewerCanRead"`
	}{
		User:          *u.User,
		ViewerCanRead: u.viewerCanRead(u.Viewer),
	})
}
