package papers

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/papers"
	"strconv"
	"time"
)

func (uc *UseCase) CreatePaper(ctx context.Context, input papers.RequestPaper, userID int) error {

	newPaper := papers.RequestPaperInsert{
		UniqueID: time.Now().Format("20060102150405"),
		UserID:   userID,
		Paper:    input,
	}

	return uc.paperRepo.CreatePaper(ctx, newPaper)
}

func (uc *UseCase) GetPaperById(ctx context.Context, ID int) (papers.ResponsePaper, error) {
	data := papers.ResponsePaper{}
	getPaperResponse, err := uc.paperRepo.GetPaperById(ctx, ID)

	if err != nil {
		return data, err
	}

	return getPaperResponse, nil
}

func (uc *UseCase) DeletePaper(ctx context.Context, id int, userId int) error {

	var (
		errNotAllowed int
	)

	getPaperResponse, err := uc.paperRepo.GetPaperById(ctx, id)
	getUserCreated, _ := strconv.Atoi(getPaperResponse.UserID)
	getReviewStatus, _ := strconv.Atoi(getPaperResponse.UserID)

	if err != nil {
		return err
	}

	statusAllowDelete := [4]int{1}

	for _, value := range statusAllowDelete {
		if value == getReviewStatus {
			errNotAllowed++
		}
	}

	if getUserCreated != userId {
		return fmt.Errorf("NOT AUTHORIZED")
	} else if errNotAllowed > 0 {
		return fmt.Errorf("PAPER ON PROCESS")
	}

	return uc.paperRepo.DeletePaper(ctx, id, userId)

}

func (uc *UseCase) UpdatePaper(ctx context.Context, input papers.RequestPaperUpdate, userID int) error {

	var (
		errNotAllowed int
	)

	getPaperResponse, err := uc.paperRepo.GetPaperById(ctx, input.ID)
	getUserCreated, _ := strconv.Atoi(getPaperResponse.UserID)
	getReviewStatus, _ := strconv.Atoi(getPaperResponse.UserID)

	if err != nil {
		return err
	}

	statusAllowUpdate := [4]int{2, 3}
	fmt.Println(getReviewStatus)

	for _, value := range statusAllowUpdate {
		if value == getReviewStatus {
			errNotAllowed++
		}
	}
	fmt.Println(errNotAllowed)

	if getUserCreated != userID {
		return fmt.Errorf("NOT AUTHORIZED")
	} else if errNotAllowed > 0 {
		return fmt.Errorf("PAPER ON PROCESS")
	}
	return uc.paperRepo.UpdatePaper(ctx, input, userID)
}
