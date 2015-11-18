package client

import "github.com/ory-am/ladon/guard/operator"

type AuthorizeRequest struct {
	Resource   string `json:"string"`
	Token      string
	Permission string            `json:"permission"`
	Context    *operator.Context `json:"context"`
}

type Client interface {
	IsAllowed(ar *AuthorizeRequest) (bool, error)
	IsAuthenticated(token string) (bool, error)
}
