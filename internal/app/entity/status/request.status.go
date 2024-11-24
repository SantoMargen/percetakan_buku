package status

type PaginationStatus struct {
	Page   int           `json:"page"`
	Size   int           `json:"size"`
	Filter *FilterStatus `json:"filter"`
}

type FilterStatus struct {
	IDStatus string `json:"id_status"`
}
