package confluence

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// UserResult is a struct to hold user data
type UserResult struct {
	Type           string `json:"type"`
	Username       string `json:"username"`
	UserKey        string `json:"userKey"`
	ProfilePicture struct {
		Path      string `json:"path"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		IsDefault bool   `json:"isDefault"`
	} `json:"profilePicture"`
	DisplayName string `json:"displayName"`
	Links       struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Self    string `json:"self"`
	} `json:"_links"`
}

func (w *Wiki) userEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(w.endPoint.String() + "/user")
}

// User gets a user based on username or key
func (w *Wiki) User(usernameOrKey string) (*UserResult, error) {
	userEndpoint, err := w.userEndpoint()
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	if len(usernameOrKey) == 32 {
		data.Set("key", usernameOrKey)
	} else {
		data.Set("username", usernameOrKey)
	}
	userEndpoint.RawQuery = data.Encode()

	req, err := http.NewRequest("GET", userEndpoint.String(), nil)
	if err != nil {
		return nil, err

	}
	res, err := w.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var result UserResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
