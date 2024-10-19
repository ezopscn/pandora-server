package dto

// 分页信息
type Pagination struct {
	PageNumber   uint  `json:"pageNumber" `  // 页码
	PageSize     uint  `json:"pageSize"`     // 每页数据量
	Total        int64 `json:"total"`        // 数据量
	NoPagination bool  `json:"noPagination"` // 不分页，默认 false，则分页
}

// 数据分页返回格式
type PaginationResponse struct {
	Pagination Pagination  `json:"pagination"`
	List       interface{} `json:"list"`
}

// 分页数据设置
const (
	MaxPageSize     uint = 100 // 每次请求最大的数据量，为了数据安全性
	DefaultPageSize uint = 30  // 默认每页数据量
)

// 分页查询
func (p *Pagination) GetPaginationLimitAndOffset(limit int, offset int) {
	if p.NoPagination {
		// 不分页，直接返回
		return
	} else {
		// 1.请求数据量不能小于 1，也不能大于最大限制的请求量，如果不合法，则使用默认值
		pageSize := p.PageSize
		if p.PageSize < 1 || p.PageSize > MaxPageSize {
			pageSize = DefaultPageSize
		}

		// 2.页码规则，页码小于 0，则返回首页
		pageNumber := p.PageNumber
		if p.PageNumber < 1 {
			pageNumber = 1
		}

		// 3.回填数据
		p.PageSize = pageSize
		p.PageNumber = pageNumber

		// 4.设置偏移量
		limit = int(pageSize)
		offset = int(pageSize * (pageNumber - 1))
		return
	}
}
