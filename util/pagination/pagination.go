package pagination

import (
	"errors"
	"fmt"
	"strconv"
)

type (
	PageRequest struct {
		page int
		size int
	}

	PageResponse struct {
		Total int64       `json:"total" validate:"required"`
		List  interface{} `json:"list" validate:"required"`
	}
)

func (p *PageRequest) Size() int {
	return p.size
}

func (p *PageRequest) Page() int {
	return p.page
}

func NewPageRequest(page, size string) (*PageRequest, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, fmt.Errorf("页数转换为整数失败: %v", err)
	}

	if pageInt < 1 {
		return nil, errors.New("页数最小为1")
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return nil, fmt.Errorf("条数转换为整数失败: %v", err)
	}

	if sizeInt < 1 {
		return nil, errors.New("条数最小为1")
	}

	return &PageRequest{
		page: pageInt,
		size: sizeInt,
	}, nil
}

func (p *PageRequest) GetOffset() int {
	return (p.page - 1) * p.size
}
