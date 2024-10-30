package publisher

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/publishers"
)

func (uc *UseCase) CreatePublisher(ctx context.Context, input publishers.PublisherRequest) error {
	countDuplikat, err := uc.publishRepo.FindByName(ctx, input.Name)

	if err != nil {
		return fmt.Errorf("failled get publisher name")
	}

	if countDuplikat > 0 {
		return fmt.Errorf("publisher has been registered")
	}

	return uc.publishRepo.CreatePublisher(ctx, input)
}

func (uc *UseCase) GetPublisherById(ctx context.Context, ID int) (publishers.PublisherResponse, error) {
	data := publishers.PublisherResponse{}
	getCategoryResponse, err := uc.publishRepo.GetPublisherById(ctx, ID)

	if err != nil {
		return data, err
	}

	return getCategoryResponse, nil
}

func (uc *UseCase) UpdatePublisher(ctx context.Context, input publishers.RequestUpdatePublisher) error {

	return uc.publishRepo.UpdatePublisher(ctx, input)
}

func (uc *UseCase) DeletePublisher(ctx context.Context, id int) error {
	getCategoryResponse, err := uc.publishRepo.GetPublisherById(ctx, id)

	if err != nil {
		return fmt.Errorf("failled get publisher")
	}

	if getCategoryResponse.Name == "" {
		return fmt.Errorf("fetch publisher failed")
	}

	return uc.publishRepo.DeletePublisher(ctx, id)

}
