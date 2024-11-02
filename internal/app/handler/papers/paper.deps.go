package papers

import (
	"context"
	"siap_app/internal/app/entity/papers"
)

type paperUC interface {
	CreatePaper(ctx context.Context, input papers.RequestPaper, userId int) error
	DeletePaper(ctx context.Context, id int, userId int) error
	GetPaperById(ctx context.Context, id int) (papers.ResponsePaper, error)
	UpdatePaper(ctx context.Context, input papers.RequestPaperUpdate, userId int) error
}
