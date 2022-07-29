package iknore

import "context"

type CreatePreSignedFileRequest struct {
	Photo     *CreatePreSignedFileRequestPhoto `json:"photo"`
	Video     *CreatePreSignedFileRequestVideo `json:"video"`
	Audio     *CreatePreSignedFileRequestAudio `json:"audio"`
	FileSize  int                              `json:"file_size"`
	ExpiredAt int64                            `json:"expired_at"`
}

type CreatePreSignedFileRequestPhoto struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type CreatePreSignedFileRequestVideo struct {
	Height    int `json:"height"`
	Width     int `json:"width"`
	Duration  int `json:"duration"`
	Framerate int `json:"framerate"`
}

type CreatePreSignedFileRequestAudio struct {
	Bitrate  int `json:"bitrate"`
	Duration int `json:"duration"`
}

type CreatePreSignedFileResponse struct {
	ID        string `json:"id"`
	UploadURL string `json:"upload_url"`
	ExpiredAt int64  `json:"expired_at"`
}

// CreatePreSignedFile
func (c *client) CreatePreSignedFile(ctx context.Context, req *CreatePreSignedFileRequest) (*CreatePreSignedFileResponse, error) {
	return nil, nil
}

type FinishPreSignedFileRequest struct {
	ID string
}

type FinishPreSignedFileResponse struct {
}

// FinishPreSignedFile
func (c *client) FinishPreSignedFile(ctx context.Context, req *FinishPreSignedFileRequest) (*FinishPreSignedFileResponse, error) {
	return nil, nil
}
