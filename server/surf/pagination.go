package surf

const defaultLimit = int(100)

type paginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type paginationResult struct {
	TotalPages int `json:"totalPages"`
}
