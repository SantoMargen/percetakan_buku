package papers

import (
	"context"
	"siap_app/internal/app/entity/papers"
)

type paperUC interface {
	CreatePaper(ctx context.Context, input papers.RequestPaper, userId int) error
	DeletePaper(ctx context.Context, id int, userId int) error
	GetPaperById(ctx context.Context, id int) (papers.ResponsePaper, error)
	GetDetailPaperUserById(ctx context.Context, input papers.PaginationPaper) ([]papers.ResponsePaperDetail, int64, error)
	UpdatePaper(ctx context.Context, input papers.RequestPaperUpdate, userId int) error
	AssignPaper(ctx context.Context, input papers.RequestPaperAssign, userID int) error
	AssignPaperPublisher(ctx context.Context, input papers.RequestPaperAssignPublisher, userID int) error
	ApprovalPaper(ctx context.Context, input papers.EntityApprovalPaper, userID string) error
}
