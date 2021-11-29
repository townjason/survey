package content

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Parameter string
	Context   *gin.Context
	UserId    int64
}

