package pinboard

import (
	"encoding/json"
)

type UserResource struct {
	token string
}

func NewUserResource(token string) UserResource {
	return UserResource{
		token: token,
	}
}

// userResponse holds the response from /user/secret and
// /user/api_token
type userResponse struct {
	Result string `json:"result"`
}

// Secret returns the user's secret RSS key (for viewing private
// feeds).
func (r UserResource) Secret() (string, error) {
	resp, err := get(userSecret, r.token, nil)
	if err != nil {
		return "", err
	}

	var ur userResponse
	err = json.Unmarshal(resp, &ur)
	if err != nil {
		return "", err
	}

	return ur.Result, nil
}

// APIToken returns the user's API token (for making API calls
// without a password).
func (r UserResource) APIToken() (string, error) {
	resp, err := get(userAPIToken, r.token, nil)
	if err != nil {
		return "", err
	}

	var ur userResponse
	err = json.Unmarshal(resp, &ur)
	if err != nil {
		return "", err
	}

	return ur.Result, nil
}
