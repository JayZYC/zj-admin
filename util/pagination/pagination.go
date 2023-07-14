package pagination

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type (
	PageRequest struct {
		CurrPage int `form:"pageNum" json:"pageNum"`
		PageSize int `form:"pageSize" json:"pageSize"`
	}

	PageResponse struct {
		Total int64       `json:"total" validate:"required"`
		List  interface{} `json:"list" validate:"required"`
	}
)

// BindPage 分页
func BindPage(c *gin.Context) (PageRequest, error) {
	var page PageRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		return page, err
	}
	err := page.check()
	return page, err
}

func (p *PageRequest) check() error {
	if p.CurrPage < 1 {
		return errors.New("页数最小为1")
	}
	if p.PageSize <= 1 {
		return errors.New("条数最小为1")
	}
	return nil
}

func (p *PageRequest) GetOffset() int {
	first := (p.CurrPage - 1) * p.PageSize
	return first
}
