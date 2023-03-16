package clients

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/mainflux/mainflux/users"
	"github.com/mainflux/migrations/tools/configs"
)

type service struct {
	url          string
	token        string
	refreshToken string
}

func (s *service) CreateUser(user users.User) error {
	return nil
}

func (s *service) Login(admin configs.AdminAccount) error {
	cred := map[string]interface{}{
		"identity": admin.ID,
		"secret":   admin.Password,
	}

	b, err := json.Marshal(cred)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, s.url, bytes.NewReader(b))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	s.token, s.refreshToken, err = getTokenFromBody(resData)
	return err
}

func getTokenFromBody(data []byte) (tok string, refreshTok string, err error) {
	type tokenRes struct {
		Token        string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	var token tokenRes

	if err = json.Unmarshal(data, &token); err != nil {
		return "", "", err
	}
	return token.Token, token.RefreshToken, nil
}
