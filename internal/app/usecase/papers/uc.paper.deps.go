package papers

import (
	"context"
	"siap_app/internal/app/entity/notification"
	"siap_app/internal/app/entity/papers"
	"siap_app/internal/app/entity/publishers"
)

type paperRepo interface {
	CreatePaper(ctx context.Context, input papers.RequestPaperInsert) error
	DeletePaper(ctx context.Context, id int, userID int) error
	GetPaperById(ctx context.Context, id int) (papers.ResponsePaper, error)
	GetDetailPaperUserById(ctx context.Context, input papers.PaginationPaper) ([]papers.ResponsePaperDetail, int64, error)
	GetDetailPaperById(ctx context.Context, id int) (papers.ResponsePaperDetail, error)
	UpdatePaper(ctx context.Context, input papers.RequestPaperUpdate, userID int) error
	AssignPaper(ctx context.Context, input papers.RequestPaperAssign, userID int) error
	AssignPaperPublisher(ctx context.Context, input papers.RequestPaperAssignPublisher, userID int) error
	ApprovalPaper(ctx context.Context, input papers.EntityApprovalPaper, userID string) error
}

type publisherRepo interface {
	GetPublisherById(ctx context.Context, id int) (publishers.PublisherResponse, error)
}

type notificationRepo interface {
	CreateLogNotif(ctx context.Context, input notification.SentRequestNotification) error
}
