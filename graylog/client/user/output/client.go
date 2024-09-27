package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) GetUserByUsername(
	ctx context.Context, username string,
) (map[string]interface{}, *http.Response, error) {
	if username == "" {
		return nil, nil, errors.New("username is required")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/users/" + username,
		ResponseBody: &body,
	})

	return body, resp, err
}
