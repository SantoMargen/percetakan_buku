package papers

type RequestPaper struct {
	Title        string `json:"title"`
	Authors      string `json:"authors"`
	CoAuthors    string `json:"co_authors"`
	Journal      string `json:"journal"`
	Volume       int    `json:"volume"`
	Issue        int    `json:"issue"`
	PageRange    string `json:"page_range"`
	DOI          string `json:"doi"`
	Abstract     string `json:"abstract"`
	Keywords     string `json:"keywords"`
	ResearchType string `json:"research_type"`
	FundingInfo  string `json:"funding_info"`
	Affiliations string `json:"affiliations"`
	FullTextLink string `json:"full_text_link"`
	Language     string `json:"language"`
	License      string `json:"license"`
	Notes        string `json:"notes"`
	URLPaper     string `json:"url_paper"`
}

type RequestPaperInsert struct {
	UniqueID string
	UserID   int
	Paper    RequestPaper
}

type RequestPaperById struct {
	ID int `json:"id"`
}

type RequestPaperUpdate struct {
	RequestPaperById
	RequestPaperInsert
}

type RequestPaperAssign struct {
	PublisherID    int            `json:"publisher_id"`
	PaperID        int            `json:"paper_id"`
	UserID         int            `json:"user_id"`
	ApprovalPosisi int            `json:"approval_posisi"` //skip this
	ApprovalList   []ApprovalList `json:"approval_list"`
}

type ApprovalList struct {
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	RoleName     string `json:"role_name"`
	ApprovalType string `json:"approval_type"`
	EntryTime    string `json:"entry_time"`
	EntryNote    string `json:"entry_note"`
}

type RequestPaperAssignPublisher struct {
	PublisherID int `json:"publisher_id"`
	PaperID     int `json:"paper_id"`
	UserID      int `json:"user_id"`
}

type RequestApprovalPaper struct {
	PaperID      int    `json:"paper_id"`
	Approval     string `json:"approval"`
	NoteApproval string `json:"note"`
}

type EntityApprovalPaper struct {
	RequestApprovalPaper
	ApprovalPosition string
	ApprovalList     []ApprovalList
	CatatanTolakan   []ApprovalList
	Status           string
}
