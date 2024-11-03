package publishers

import (
	"context"
	"database/sql"
	"fmt"
	"siap_app/internal/app/entity/publishers"
	"strconv"

	"github.com/pkg/errors"
)

func (r *repository) GetPublisherAll(ctx context.Context, input publishers.PublisherPagination) ([]publishers.PublisherResponse, int64, error) {

	var (
		dataPublisherList []publishers.PublisherResponse
		offset            int
		query             string
		countQuery        string
		total             int64
	)

	offset = (input.Page - 1) * input.Size

	query = "SELECT " + columnSelectPublisher + " FROM publishers WHERE 1=1"
	countQuery = "SELECT COUNT(*) FROM publishers WHERE 1=1"

	var args []interface{}

	if input.Filter != nil {
		if input.Filter.Name != "" {
			query += " AND name = $" + strconv.Itoa(len(args)+1)
			countQuery += " AND name = $" + strconv.Itoa(len(args)+1)
			args = append(args, input.Filter.Name)
		}
		if input.Filter.Phone != "" {
			query += " AND phone LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND phone LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.Phone+"%")
		}
		if input.Filter.Email != "" {
			query += " AND email LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND email LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.Email+"%")
		}
		if input.Filter.ContactPerson != "" {
			query += " AND contact_person_1 LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND contact_person_1 LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.ContactPerson+"%")
		}
		if input.Filter.FlagStatus != 0 {
			query += " AND flag_status = $" + strconv.Itoa(len(args)+1)
			countQuery += " AND flag_status = $" + strconv.Itoa(len(args)+1)
			args = append(args, input.Filter.FlagStatus)
		}

	}

	query += " ORDER BY publisher_id ASC LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, input.Size, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataPublisher publishers.PublisherResponse
		if err := rows.Scan(
			&dataPublisher.PublishersID,
			&dataPublisher.Name,
			&dataPublisher.Address,
			&dataPublisher.Phone,
			&dataPublisher.Email,
			&dataPublisher.Website,
			&dataPublisher.FoundedYear,
			&dataPublisher.Country,
			&dataPublisher.ContactPerson1,
			&dataPublisher.ContactPerson2,
			&dataPublisher.Fax,
			&dataPublisher.SocialMediaFBLinks,
			&dataPublisher.SocialMediaTwitterLinks,
			&dataPublisher.Website,
			&dataPublisher.JoinDate,
			&dataPublisher.EntryUser,
			&dataPublisher.EntryName,
			&dataPublisher.EntryTime,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		dataPublisherList = append(dataPublisherList, dataPublisher)
	}

	err = r.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count items: %w", err)
	}

	return dataPublisherList, total, nil
}

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
		&publisher.EntryUser,
		&publisher.EntryName,
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

func (r *repository) GetTaskPublisherAll(ctx context.Context, input publishers.PublisherPagination) ([]publishers.TaskPublisherResponse, int64, error) {

	var (
		dataPublisherList []publishers.TaskPublisherResponse
		offset            int
		query             string
		countQuery        string
		total             int64
	)

	offset = (input.Page - 1) * input.Size

	query = queryTaskPublisher + " WHERE 1=1"
	countQuery = queryCountTaskPublisher

	var args []interface{}

	if input.Filter != nil {
		if input.Filter.Name != "" {
			query += " AND publishers.name = $" + strconv.Itoa(len(args)+1)
			countQuery += " AND publishers.name = $" + strconv.Itoa(len(args)+1)
			args = append(args, input.Filter.Name)
		}
		if input.Filter.Phone != "" {
			query += " AND publishers.phone LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND publishers.phone LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.Phone+"%")
		}
		if input.Filter.Email != "" {
			query += " AND publishers.email LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND publishers.email LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.Email+"%")
		}
		if input.Filter.ContactPerson != "" {
			query += " AND publishers.contact_person_1 LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND publishers.contact_person_1 LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.ContactPerson+"%")
		}

	}

	query += " ORDER BY publishers.publisher_id ASC LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, input.Size, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataPublisher publishers.TaskPublisherResponse
		if err := rows.Scan(
			&dataPublisher.PublishersID,
			&dataPublisher.Name,
			&dataPublisher.Address,
			&dataPublisher.Phone,
			&dataPublisher.Email,
			&dataPublisher.Website,
			&dataPublisher.FoundedYear,
			&dataPublisher.Country,
			&dataPublisher.ContactPerson1,
			&dataPublisher.ContactPerson2,
			&dataPublisher.Fax,
			&dataPublisher.SocialMediaFBLinks,
			&dataPublisher.SocialMediaTwitterLinks,
			&dataPublisher.Website,
			&dataPublisher.JoinDate,
			&dataPublisher.AssignById,
			&dataPublisher.AssignByName,
			&dataPublisher.AssignByDate,
			&dataPublisher.FlagAssignStatus,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		dataPublisherList = append(dataPublisherList, dataPublisher)
	}

	err = r.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count items: %w", err)
	}

	return dataPublisherList, total, nil
}
