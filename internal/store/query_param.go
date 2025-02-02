package store

import "github.com/gin-gonic/gin"

type CategoryQueryParam struct {
	Page     int    `json:"page" form:"page" validate:"gte=1"`
	Size     int    `json:"size" form:"size" validate:"gte=1,lte=100"`
	Name     string `json:"name" form:"name" validate:"omitempty,min=1,max=100"`
	Sort     string `json:"sort" form:"sort" validate:"omitempty,oneof=name created_at"`
	SortType string `json:"sort_type" form:"sort_type" validate:"omitempty,oneof=ASC DESC"`
	Offset   int
	Limit    int
}

func (q *CategoryQueryParam) Parse(c *gin.Context) error {
	if err := c.ShouldBindQuery(&q); err != nil {
		return err
	}

	q.Limit = q.Size
	q.Offset = (q.Page - 1) * q.Size
	if q.Sort == "" {
		q.Sort = "created_at"
	}
	if q.SortType == "" {
		q.SortType = "DESC"
	}

	return nil
}
