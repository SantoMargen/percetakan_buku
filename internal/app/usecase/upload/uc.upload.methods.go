package upload

import (
	"context"
	"siap_app/internal/app/entity/upload"
)

func (uc *UseCase) UploadFile(ctx context.Context, file []upload.RequestUpload) error {
	return uc.uploadRepo.UploadFile(ctx, file)
}
