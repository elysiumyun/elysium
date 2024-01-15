package controller

import (
	"net/http"
	"time"

	"github.com/elysiumyun/elysium/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func EntryQuery(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"total":     0,
		"status":    http.StatusOK,
		"timestamp": time.Now().UnixMilli(),
		"data":      []string{},
	}, "entry query success")
}
