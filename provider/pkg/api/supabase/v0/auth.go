package v0

import (
	"context"
	"fmt"
	"net/http"
)

func WithAuthToken(token string) ClientOption {
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	})
}
