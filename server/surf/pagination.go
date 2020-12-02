package surf

const defaultLimit = int(100)

type paginationParams struct {
	Page   int `json:"page"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type paginationResult struct {
	TotalPage int `json:"totalPage"`
}
