package upload

import (
	"context"
	"siap_app/internal/app/entity/upload"
)

type uploadRepo interface {
	UploadFile(ctx context.Context, file []upload.RequestUpload) error
}
