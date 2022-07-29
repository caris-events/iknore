package iknore

import "context"

type Client interface {
	CreatePreSignedFile(context.Context, CreatePreSignedFileRequest) (CreatePreSignedFileResponse, error)
	FinishPreSignedFile(context.Context, FinishPreSignedFileRequest) (FinishPreSignedFileResponse, error)
}

type client struct {
	conf *Config
}

type Config struct {
	ServerURL string
	Key       string
}

func NewClient(config *Config) Client {
	return &client{}
}
