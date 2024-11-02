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
