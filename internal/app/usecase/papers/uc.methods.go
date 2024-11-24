package papers

import (
	"context"
	"encoding/json"
	"fmt"
	"siap_app/internal/app/entity/notification"
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

func (uc *UseCase) GetPaperById(ctx context.Context, ID int) (papers.ResponsePaperDetail, error) {
	data := papers.ResponsePaperDetail{}
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
	getUserCreated, _ := strconv.Atoi(getPaperResponse.Paper.UserID)
	getReviewStatus, _ := strconv.Atoi(getPaperResponse.Paper.UserID)

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
	getUserCreated, _ := strconv.Atoi(getPaperResponse.Paper.UserID)
	getReviewStatus := getPaperResponse.Status.IdStatus

	if err != nil {
		return err
	}

	statusAllowUpdate := [4]int{2, 3}

	for _, value := range statusAllowUpdate {
		if value == getReviewStatus {
			errNotAllowed++
		}
	}

	if getUserCreated != userID {
		return fmt.Errorf("NOT AUTHORIZED")
	} else if errNotAllowed > 0 {
		return fmt.Errorf("PAPER ON PROCESS")
	}
	return uc.paperRepo.UpdatePaper(ctx, input, userID)
}

func (uc *UseCase) AssignPaper(ctx context.Context, input papers.RequestPaperAssign, userID int, fullName string) error {
	var (
		errNotAllowed int
	)

	getPaperResponse, err := uc.paperRepo.GetPaperById(ctx, input.PaperID)

	if err != nil {
		return err
	}
	getReviewStatus, _ := strconv.Atoi(getPaperResponse.Paper.UserID)
	getPublisherResponse, err := uc.publisherRepo.GetPublisherById(ctx, input.PublisherID)

	if err != nil {
		return err
	}

	statusAllowAssign := [4]int{0}

	for _, value := range statusAllowAssign {
		if value == getReviewStatus {
			errNotAllowed++
		}
	}
	if getPublisherResponse.Name == "" {
		return fmt.Errorf("PUBLISHER NOT FOUND")
	} else if errNotAllowed > 0 {
		return fmt.Errorf("PAPER ON PROCESS")
	}

	sentNotification := notification.RequestNotification{
		KeyNotif:    input.PaperID,
		DescNotif:   "",
		TitleNotif:  "Approval Paper " + getPaperResponse.Paper.UniqueID,
		Receiver:    input.ApprovalList[0].UserID,
		UrlRedirect: "http://localhost:8080/paper-id",
	}

	dataSentNotification := notification.SentRequestNotification{
		RequestNotification: sentNotification,
		Sender:              userID,
	}
	input.ApprovalPosisi = input.ApprovalList[0].UserID

	uc.notificationRepo.CreateLogNotif(ctx, dataSentNotification)

	return uc.paperRepo.AssignPaper(ctx, input, userID, fullName)

}

func (uc *UseCase) AssignPaperPublisher(ctx context.Context, input papers.RequestPaperAssignPublisher, userID int) error {

	getPaperResponse, err := uc.paperRepo.GetPaperById(ctx, input.PaperID)

	if err != nil {

		return err
	}

	getPublisherResponse, err := uc.publisherRepo.GetPublisherById(ctx, input.PublisherID)

	if err != nil {
		return err
	}

	if getPaperResponse.Paper.FlagAssign == "1" {

		return fmt.Errorf("PAPER HAS BEEN ASSIGN TO PUBLISHER")

	} else if getPublisherResponse.Name == "" {

		return fmt.Errorf("PUBLISHER NOT FOUND")

	}

	return uc.paperRepo.AssignPaperPublisher(ctx, input, userID)

}

func (uc *UseCase) ApprovalPaper(ctx context.Context, input papers.EntityApprovalPaper, userID string) error {
	var (
		approvalList   []papers.ApprovalList
		noteRejectList []papers.ApprovalList
		getEntityPaper papers.EntityApprovalPaper
	)

	formattedTime := time.Now().Format("2006-01-02 15:04:05")

	getPaperResponse, err := uc.paperRepo.GetPaperById(ctx, input.PaperID)
	if err != nil {
		return fmt.Errorf("error fetching paper details: %w", err)
	}

	if err := parseRejectionNotes(getPaperResponse.Paper.CatatanReject, &noteRejectList); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(getPaperResponse.Paper.ApprovalList), &approvalList); err != nil {
		return fmt.Errorf("failed to parse approval list: %w", err)
	}

	if getPaperResponse.Paper.FlagAssign == "1" {
		return fmt.Errorf("paper has been assigned to publisher")
	}

	if err := handleApprovals(approvalList, input, userID, formattedTime, noteRejectList, &getEntityPaper); err != nil {
		return err
	}

	if getEntityPaper.PaperID != 0 {

		resApproval := uc.paperRepo.ApprovalPaper(ctx, getEntityPaper, userID)

		if resApproval == nil {

			convertUserId, _ := strconv.Atoi(userID)

			sendNotification := func(title, approval, receiverID string) {
				convertPositionId, _ := strconv.Atoi(receiverID)

				sentNotification := notification.RequestNotification{
					KeyNotif:    input.PaperID,
					DescNotif:   approval,
					TitleNotif:  title + " " + getPaperResponse.Paper.UniqueID,
					Receiver:    convertPositionId,
					UrlRedirect: "http://localhost:8080/paper-id",
				}

				dataSentNotification := notification.SentRequestNotification{
					RequestNotification: sentNotification,
					Sender:              convertUserId,
				}

				uc.notificationRepo.CreateLogNotif(ctx, dataSentNotification)
			}

			if input.Approval == "approve" && getEntityPaper.ApprovalPosition != "0" {
				sendNotification("Approval Paper", input.Approval, getEntityPaper.ApprovalPosition)
			} else if input.Approval == "reject" {
				sendNotification("Reject Paper", input.Approval, getPaperResponse.Paper.UserID)
			}

		}

		return resApproval
	}

	return fmt.Errorf("no task approval record")
}

func (uc *UseCase) GetListPapers(ctx context.Context, input papers.PaginationPaper) ([]papers.ResponsePaperDetail, int64, error) {

	resp, total, err := uc.paperRepo.GetListPapers(ctx, input)

	if err != nil {
		return nil, 0, fmt.Errorf("error get data paper user by id : %w", err)
	}

	return resp, total, nil

}

func parseRejectionNotes(rejectionNotes string, noteRejectList *[]papers.ApprovalList) error {
	if rejectionNotes != "" {
		if err := json.Unmarshal([]byte(rejectionNotes), noteRejectList); err != nil {
			return fmt.Errorf("failed to parse reject note list: %w", err)
		}
	}
	return nil
}

func handleApprovals(approvalList []papers.ApprovalList, input papers.EntityApprovalPaper, userID, formattedTime string, noteRejectList []papers.ApprovalList, getEntityPaper *papers.EntityApprovalPaper) error {
	var (
		getRequestApproval papers.RequestApprovalPaper
	)
	userIdLast := len(approvalList) - 1
	for i, value := range approvalList {

		parseStringUserID := strconv.Itoa(value.UserID)
		if parseStringUserID != userID {
			continue
		}

		approvalList[i].EntryTime = formattedTime
		approvalList[i].EntryNote = input.NoteApproval
		approvalList[i].ApprovalType = "approve"

		noteReject := papers.ApprovalList{
			Name:         approvalList[i].Name,
			RoleName:     approvalList[i].RoleName,
			ApprovalType: input.Approval,
			EntryTime:    approvalList[i].EntryTime,
			EntryNote:    approvalList[i].EntryNote,
		}
		noteRejectList = append(noteRejectList, noteReject)

		getRequestApproval = papers.RequestApprovalPaper{
			PaperID:      input.PaperID,
			Approval:     input.Approval,
			NoteApproval: input.NoteApproval,
		}

		if input.Approval == "approve" {
			if userIdLast == i {
				getEntityPaper.ApprovalPosition = "0"
				getEntityPaper.Status = "5"
			} else {
				getEntityPaper.ApprovalPosition = strconv.Itoa(approvalList[i+1].UserID)
				getEntityPaper.Status = "3"
			}

		} else {
			getEntityPaper.ApprovalPosition = "0"
			getEntityPaper.Status = "4"
			getEntityPaper.CatatanTolakan = noteRejectList
		}

		getEntityPaper.RequestApprovalPaper = getRequestApproval
		getEntityPaper.ApprovalList = approvalList
		getEntityPaper.PaperID = input.PaperID

		return nil
	}

	return fmt.Errorf("user ID not found in approval list")
}
