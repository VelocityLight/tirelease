package entity

type ListOption struct {
	// offset
	Page    int64 `json:"page,omitempty" form:"page"`
	PerPage int64 `json:"per_page,omitempty" form:"per_page"`
	Offset  int64 `json:"offset,omitempty" form:"offset"`

	// order by
	OrderBy string `json:"order_by,omitempty" form:"order_by"`
	Order   string `json:"order,omitempty" form:"order"`
}

type ListResponse struct {
	TotalCount int64 `json:"total_count"`
	TotalPage  int64 `json:"total_page"`
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
}

func (option *ListOption) CalcOffset() {
	if option.Page == 0 || option.PerPage == 0 {
		option.Offset = 0
	}
	option.Offset = (option.Page - 1) * option.PerPage
}

func (response *ListResponse) CalcTotalPage() {
	if response.PerPage == 0 {
		return
	}
	response.TotalPage = response.TotalCount / response.PerPage
	if response.TotalCount%response.PerPage != 0 {
		response.TotalPage = response.TotalPage + 1
	}
}
