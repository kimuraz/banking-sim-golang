package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetOffsetLimit(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "15"))

	if size > 50 {
		size = 50
	}

	offset := (page - 1) * size

	return offset, size
}
