package model

import (
	"context"
	"net/url"
)

type MessageParams struct {
	Success string
	Error   string
}

func NewMessageParams(params url.Values) MessageParams {
	return MessageParams{
		Success: params.Get("success"),
		Error:   params.Get("error"),
	}
}

func GetMessageParams(ctx context.Context) MessageParams {
	if params, ok := ctx.Value("message").(MessageParams); ok {
		return params
	}
	return NewMessageParams(nil)
}
