package role

import "github.com/gin-gonic/gin"

type Repository interface{}

type repository struct {
	ctx *gin.Context
}
