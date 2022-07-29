package endpoint

import "github.com/gin-gonic/gin"

type Endpoint interface {
	CreatePreSignedFile(*gin.Context)
	FinishPreSignedFile(*gin.Context)
	DeleteFile(*gin.Context)
	ListFiles(*gin.Context)
	GetFile(*gin.Context)
}

type endpoint struct {
}

func New() Endpoint {
	return &endpoint{}
}
