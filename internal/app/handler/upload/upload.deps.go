package upload

import (
	"context"
	"siap_app/internal/app/entity/upload"
)

type uploadFileUC interface {
	UploadFile(ctx context.Context, input []upload.RequestUpload) error
}
