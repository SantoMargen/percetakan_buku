package paper

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"siap_app/internal/app/entity/papers"

	"github.com/pkg/errors"
)

func (r *repository) CreatePaper(ctx context.Context, input papers.RequestPaperInsert) error {

	var (
		countCek int
	)
	errCekFile := r.db.QueryRowContext(ctx, queryCekFileExist, input.Paper.UniqueID).Scan(&countCek)

	if errCekFile != nil {
		return errors.Wrap(errCekFile, "failed to create paper")
	}

	if countCek == 0 {
		return errors.Wrap(errors.Errorf("file not found "), " failed to create paper")
	}

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
		input.Paper.License,
		input.Paper.Notes,
		input.Paper.UniqueID,
	)

	if err != nil {
		return errors.Wrap(err, "failed to create paper")
	}

	return nil
}

func (r *repository) GetPaperById(ctx context.Context, paperID int) (papers.ResponsePaper, error) {
	var paper papers.ResponsePaper
	err := r.db.QueryRowContext(ctx, querySelectPaper, paperID).Scan(
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
		&paper.License,
		&paper.Notes,
		&paper.CreatedAt,
		&paper.UpdateAt,
		&paper.FlagAssign,
		&paper.URLPaper,
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
		input.Paper.License,
		input.Paper.Notes,
		input.Paper.UniqueID,
		input.ID,
	)

	if err != nil {
		return errors.Wrap(err, "failed to create paper")
	}

	return nil
}

func (r *repository) AssignPaper(ctx context.Context, input papers.RequestPaperAssign, userID int) error {

	approvalJSON, errParse := json.Marshal(input.ApprovalList)

	if errParse != nil {
		return errors.Wrap(errParse, "failed parse approval list")
	}

	_, err := r.db.ExecContext(ctx, queryAssignPaper,
		input.PaperID,
		input.PublisherID,
		input.ApprovalPosisi,
		string(approvalJSON),
		userID,
	)

	if err != nil {
		return errors.Wrap(err, "failed to assign paper to publisher")
	}

	return nil
}

func (r *repository) AssignPaperPublisher(ctx context.Context, input papers.RequestPaperAssignPublisher, userID int) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "start transaction error")
	}
	defer func() {
		if getError := recover(); getError != nil {
			_ = tx.Rollback()
			log.Println("Recovered in AssignPaperPublisher:", getError)
		}
	}()

	prepareInsert, err := tx.PrepareContext(ctx, queryAssignPaperPublisher)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "prepare insert statement error")
	}
	defer prepareInsert.Close()

	if _, err := prepareInsert.ExecContext(ctx, input.PaperID, input.PublisherID, userID); err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "failed to execute insert statement")
	}

	prepareUpdate, err := tx.PrepareContext(ctx, queryUpdateFlagSubmission)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "prepare update statement error")
	}
	defer prepareUpdate.Close()

	if _, err := prepareUpdate.ExecContext(ctx, input.PaperID); err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "failed to execute update statement")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit transaction error")
	}

	return nil
}

func (r *repository) ApprovalPaper(ctx context.Context, input papers.EntityApprovalPaper, userID string) error {

	var (
		count int
	)

	getNewApprovalList, err := json.Marshal(input.ApprovalList)
	if err != nil {
		return errors.Wrap(err, "failled get task approval")
	}

	getNewApprovaRejectNote, err := json.Marshal(input.CatatanTolakan)
	if err != nil {
		return errors.Wrap(err, "failled get task reject note")
	}

	errCheckTask := r.db.QueryRowContext(ctx, queryCheckTask,
		userID,
		input.PaperID,
	).Scan(&count)

	if errCheckTask != nil {
		return errors.Wrap(errCheckTask, "failled get task approval")
	} else {
		if count == 0 {
			return errors.Errorf("no task to approval")
		} else {
			_, err := r.db.ExecContext(ctx, queryApprovalPaper,
				input.ApprovalPosition,
				getNewApprovalList,
				getNewApprovaRejectNote,
				input.Status,
				input.PaperID,
				userID,
			)

			if err != nil {
				return errors.Wrap(err, "failed execute approval paper")
			} else {
				return nil
			}

		}
	}

}

func (r *repository) GetDetailPaperById(ctx context.Context, paperID int) (papers.ResponsePaperDetail, error) {
	var paper papers.ResponsePaperDetail

	err := r.db.QueryRowContext(ctx, queryDetailPaper, paperID).Scan(
		&paper.Paper.ID,
		&paper.Paper.UserID,
		&paper.Paper.UniqueID,
		&paper.Paper.Title,
		&paper.Paper.Authors,
		&paper.Paper.CoAuthors,
		&paper.Paper.PublicationDate,
		&paper.Paper.Journal,
		&paper.Paper.Volume,
		&paper.Paper.Issue,
		&paper.Paper.PageRange,
		&paper.Paper.DOI,
		&paper.Paper.Abstract,
		&paper.Paper.Keywords,
		&paper.Paper.ResearchType,
		&paper.Paper.FundingInfo,
		&paper.Paper.Affiliations,
		&paper.Paper.FullTextLink,
		&paper.Paper.Language,
		&paper.Paper.License,
		&paper.Paper.Notes,
		&paper.Paper.CreatedAt,
		&paper.Paper.UpdateAt,
		&paper.Paper.FlagAssign,
		&paper.Publisher.Name,
		&paper.Publisher.Address,
		&paper.Publisher.Phone,
		&paper.Publisher.Email,
		&paper.Publisher.Website,
		&paper.Publisher.FoundedYear,
		&paper.Publisher.Country,
		&paper.Publisher.ContactPerson1,
		&paper.Publisher.ContactPerson2,
		&paper.Publisher.Fax,
		&paper.Publisher.SocialMediaFBLinks,
		&paper.Publisher.SocialMediaTwitterLinks,
		&paper.Publisher.SocialMediaWebLinks,
		&paper.Publisher.JoinDate,
		&paper.Publisher.EntryUserPublisher,
		&paper.Publisher.EntryNamePublisher,
		&paper.Publisher.EntryTimePublisher,
		&paper.ApprovalSubmission.ApprovalPosition,
		&paper.ApprovalSubmission.ApprovalList,
		&paper.ApprovalSubmission.CatatanReject,
		&paper.ApprovalSubmission.EntryUserAssignApproval,
		&paper.ApprovalSubmission.EntryNameAssignApproval,
		&paper.ApprovalSubmission.EntryTimeAssignApproval,
		&paper.AssignPaperPublisher.EntryUserAssignPublisher,
		&paper.AssignPaperPublisher.EntryNameAssignPublisher,
		&paper.AssignPaperPublisher.EntryTimeAssignPublisher,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return paper, errors.Wrap(err, "paper not found")
		}
		return paper, errors.Wrap(err, "failed to get paper by id")
	}

	return paper, nil
}
