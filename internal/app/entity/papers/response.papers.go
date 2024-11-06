package papers

import "time"

type ResponsePaper struct {
	ID              string    `json:"id"`
	UniqueID        string    `json:"unique_id"`
	UserID          string    `json:"user_id"`
	Title           string    `json:"title"`
	Authors         string    `json:"authors"`
	CoAuthors       string    `json:"co_authors"`
	PublicationDate time.Time `json:"publication_date"`
	Journal         string    `json:"journal"`
	Volume          int       `json:"volume"`
	Issue           int       `json:"issue"`
	PageRange       string    `json:"page_range"`
	DOI             string    `json:"doi"`
	Abstract        string    `json:"abstract"`
	Keywords        string    `json:"keywords"`
	ResearchType    string    `json:"research_type"`
	FundingInfo     string    `json:"funding_info"`
	Affiliations    string    `json:"affiliations"`
	FullTextLink    string    `json:"full_text_link"`
	Language        string    `json:"language"`
	ReviewStatus    string    `json:"review_status"`
	License         string    `json:"license"`
	Notes           string    `json:"notes"`
	FlagAssign      string    `json:"flag_assign"`
	URLPaper        string    `json:"url"`
	EntryNamePaper  string    `json:"entry_name_paper"`
	EntryUserPaper  string    `json:"entry_user_paper"`
	CreatedAt       time.Time `json:"created_at"`
	UpdateAt        time.Time `json:"updated_at"`
}

type ResponsePaperDetail struct {
	Paper                ResponsePaper
	AssignPaperPublisher AssignPaperToPublisher
	ApprovalSubmission   ApprovalSubmissionPaper
	Publisher            Publisher
}

type ApprovalSubmissionPaper struct {
	ApprovalPosition        string `json:"approval_position"`
	ApprovalList            string `json:"approval_list"`
	CatatanReject           string `json:"catatan_tolakan"`
	EntryUserAssignApproval string `json:"entry_user_assign_approval"`
	EntryNameAssignApproval string `json:"entry_name_assign_approval"`
	EntryTimeAssignApproval string `json:"entry_time_assign_approval"`
}

type Publisher struct {
	PublishersID            string `json:"publishers_id"`
	Name                    string `json:"name"`
	Address                 string `json:"address"`
	Phone                   string `json:"phone"`
	Email                   string `json:"email"`
	Website                 string `json:"website"`
	FoundedYear             int    `json:"founded_year"`
	Country                 string `json:"country"`
	ContactPerson1          string `json:"contact_person_1"`
	ContactPerson2          string `json:"contact_person_2"`
	Fax                     string `json:"fax"`
	SocialMediaFBLinks      string `json:"social_fb_links"`
	SocialMediaTwitterLinks string `json:"social_twitter_links"`
	SocialMediaWebLinks     string `json:"social_web_links"`
	JoinDate                string `json:"join_date"`
	EntryTimePublisher      string `json:"entry_time_publihser"`
	EntryNamePublisher      string `json:"entry_name_publisher"`
	EntryUserPublisher      string `json:"entry_user_publisher"`
}

type AssignPaperToPublisher struct {
	EntryNameAssignPublisher string `json:"entry_name_to_publisher"`
	EntryUserAssignPublisher string `json:"entry_user_to_publisher"`
	EntryTimeAssignPublisher string `json:"entry_time_to_publisher"`
}
