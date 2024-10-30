package publishers

import (
	"context"
	"database/sql"
	"siap_app/internal/app/entity/publishers"

	"github.com/pkg/errors"
)

func (r *repository) CreatePublisher(ctx context.Context, input publishers.PublisherRequest) error {
	_, err := r.db.ExecContext(ctx, queryCreatePublisher,
		input.Name,
		input.Address,
		input.Phone,
		input.Email,
		input.Website,
		input.FoundedYear,
		input.Country,
		input.ContactPerson1,
		input.ContactPerson2,
		input.Fax,
		input.SocialMediaFBLinks,
		input.SocialMediaTwitterLinks,
		input.SocialMediaWebLinks,
		input.JoinDate)

	if err != nil {
		return errors.Wrap(err, "failed to create publishers")
	}

	return nil
}

func (r *repository) GetPublisherById(ctx context.Context, publisherId int) (publishers.PublisherResponse, error) {
	var publisher publishers.PublisherResponse
	err := r.db.QueryRowContext(ctx, queryPublishersById, publisherId).Scan(
		&publisher.PublishersID,
		&publisher.Name,
		&publisher.Address,
		&publisher.Phone,
		&publisher.Email,
		&publisher.Website,
		&publisher.FoundedYear,
		&publisher.Country,
		&publisher.ContactPerson1,
		&publisher.ContactPerson2,
		&publisher.Fax,
		&publisher.SocialMediaFBLinks,
		&publisher.SocialMediaTwitterLinks,
		&publisher.SocialMediaWebLinks,
		&publisher.JoinDate,
		&publisher.EntryTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return publisher, errors.Wrap(err, "publisher not found")
		}
		return publisher, errors.Wrap(err, "failed to get publisher by id")
	}

	return publisher, nil
}

func (r *repository) UpdatePublisher(ctx context.Context, input publishers.RequestUpdatePublisher) error {
	_, err := r.db.ExecContext(ctx, queryUpdatePublisher,
		input.Name,
		input.Address,
		input.Phone,
		input.Email,
		input.Website,
		input.FoundedYear,
		input.Country,
		input.ContactPerson1,
		input.ContactPerson2,
		input.Fax,
		input.SocialMediaFBLinks,
		input.SocialMediaTwitterLinks,
		input.SocialMediaWebLinks,
		input.JoinDate,
		input.ID,
	)

	if err != nil {
		return errors.Wrap(err, "failed to update publishers")
	}

	return nil
}

func (r *repository) DeletePublisher(ctx context.Context, ID int) error {
	_, err := r.db.ExecContext(ctx, queryDeletePublisher,
		ID,
	)

	if err != nil {
		return errors.Wrap(err, "failed to delete publishers")
	}

	return nil
}

func (r *repository) FindByName(ctx context.Context, name string) (int, error) {
	var countData int
	err := r.db.QueryRowContext(ctx, queryPublishersByName, name).Scan(
		&countData,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return countData, errors.Wrap(err, "publisher not found")
		}
		return countData, errors.Wrap(err, "failed to get publisher by name")
	}

	return countData, nil
}
