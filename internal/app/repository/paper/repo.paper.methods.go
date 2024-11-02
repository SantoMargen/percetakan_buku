package paper

import (
	"context"
	"database/sql"
	"siap_app/internal/app/entity/papers"

	"github.com/pkg/errors"
)

func (r *repository) CreatePaper(ctx context.Context, input papers.RequestPaperInsert) error {

	_, err := r.db.ExecContext(ctx, queryInsertPaper,
		input.UniqueID,
		input.UserID,
		input.Paper.Title,
		input.Paper.Authors,
		input.Paper.CoAuthors,
		input.Paper.Journal,
		input.Paper.Volume,
		input.Paper.Issue,
		input.Paper.PageRange,
		input.Paper.DOI,
		input.Paper.Abstract,
		input.Paper.Keywords,
		input.Paper.ResearchType,
		input.Paper.FundingInfo,
		input.Paper.Affiliations,
		input.Paper.FullTextLink,
		input.Paper.Language,
		1,
		input.Paper.License,
		input.Paper.Notes,
	)

	if err != nil {
		return errors.Wrap(err, "failed to create paper")
	}

	return nil
}

func (r *repository) GetPaperById(ctx context.Context, category_id int) (papers.ResponsePaper, error) {
	var paper papers.ResponsePaper
	err := r.db.QueryRowContext(ctx, querySelectPaper, category_id).Scan(
		&paper.ID,
		&paper.UserID,
		&paper.UniqueID,
		&paper.Title,
		&paper.Authors,
		&paper.CoAuthors,
		&paper.PublicationDate,
		&paper.Journal,
		&paper.Volume,
		&paper.Issue,
		&paper.PageRange,
		&paper.DOI,
		&paper.Abstract,
		&paper.Keywords,
		&paper.ResearchType,
		&paper.FundingInfo,
		&paper.Affiliations,
		&paper.FullTextLink,
		&paper.Language,
		&paper.ReviewStatus,
		&paper.License,
		&paper.Notes,
		&paper.CreatedAt,
		&paper.UpdateAt,
		&paper.ApprovalPosition,
		&paper.ApprovalList,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return paper, errors.Wrap(err, "paper not found")
		}
		return paper, errors.Wrap(err, "failed to get paper by id")
	}

	return paper, nil
}

func (r *repository) DeletePaper(ctx context.Context, id int, userID int) error {
	_, err := r.db.ExecContext(ctx, queryDeletePaper, id)

	if err != nil {
		return errors.Wrap(err, "failed to delete paper")
	}

	return nil
}

func (r *repository) UpdatePaper(ctx context.Context, input papers.RequestPaperUpdate, userID int) error {

	_, err := r.db.ExecContext(ctx, queryUpdatePaper,
		input.Paper.Title,
		input.Paper.Authors,
		input.Paper.CoAuthors,
		input.Paper.Journal,
		input.Paper.Volume,
		input.Paper.Issue,
		input.Paper.PageRange,
		input.Paper.DOI,
		input.Paper.Abstract,
		input.Paper.Keywords,
		input.Paper.ResearchType,
		input.Paper.FundingInfo,
		input.Paper.Affiliations,
		input.Paper.FullTextLink,
		input.Paper.Language,
		1,
		input.Paper.License,
		input.Paper.Notes,
		input.ID,
	)

	if err != nil {
		return errors.Wrap(err, "failed to create paper")
	}

	return nil
}
