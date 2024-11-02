package publishers

import (
	"context"
	"siap_app/internal/app/entity/publishers"
)

type publisherUC interface {
	CreatePublisher(ctx context.Context, input publishers.PublisherRequest) error
	GetPublisherById(ctx context.Context, id int) (publishers.PublisherResponse, error)
	DeletePublisher(ctx context.Context, id int) error
	UpdatePublisher(ctx context.Context, input publishers.RequestUpdatePublisher) error
}