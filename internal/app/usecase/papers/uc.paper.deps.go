package papers

import (
	"context"
	"siap_app/internal/app/entity/papers"
)

type paperRepo interface {
	CreatePaper(ctx context.Context, input papers.RequestPaperInsert) error
	DeletePaper(ctx context.Context, id int, userID int) error
	GetPaperById(ctx context.Context, id int) (papers.ResponsePaper, error)
	UpdatePaper(ctx context.Context, input papers.RequestPaperUpdate, userID int) error
}
