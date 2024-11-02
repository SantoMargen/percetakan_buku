package papers

import "time"

type ResponsePaper struct {
	ID               string    `json:"id"`
	UniqueID         string    `json:"unique_id"`
	UserID           string    `json:"user_id"`
	Title            string    `json:"title"`
	Authors          string    `json:"authors"`
	CoAuthors        string    `json:"co_authors"`
	PublicationDate  time.Time `json:"publication_date"`
	Journal          string    `json:"journal"`
	Volume           int       `json:"volume"`
	Issue            int       `json:"issue"`
	PageRange        string    `json:"page_range"`
	DOI              string    `json:"doi"`
	Abstract         string    `json:"abstract"`
	Keywords         string    `json:"keywords"`
	ResearchType     string    `json:"research_type"`
	FundingInfo      string    `json:"funding_info"`
	Affiliations     string    `json:"affiliations"`
	FullTextLink     string    `json:"full_text_link"`
	Language         string    `json:"language"`
	ReviewStatus     string    `json:"review_status"`
	License          string    `json:"license"`
	Notes            string    `json:"notes"`
	CreatedAt        time.Time `json:"created_at"`
	UpdateAt         time.Time `json:"updated_at"`
	ApprovalPosition string    `json:"approval_position"`
	ApprovalList     string    `json:"approval_list"`
}
