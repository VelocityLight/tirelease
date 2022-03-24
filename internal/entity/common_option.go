package entity

type ListOption struct {
	// offset
	Page    int `json:"page,omitempty" form:"page"`
	PerPage int `json:"per_page,omitempty" form:"per_page"`
	Offset  int `json:"offset,omitempty" form:"offset"`

	// order by
	OrderBy string `json:"order_by,omitempty" form:"order_by"`
	Order   string `json:"order,omitempty" form:"order"`
}

func (option *ListOption) CalcOffset() {
	if option.Page == 0 || option.PerPage == 0 {
		option.Offset = 0
	}
	option.Offset = (option.Page - 1) * option.PerPage
}
