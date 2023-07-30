package account

import (
	"github.com/gin-gonic/gin"
)

type Repository interface {
	IAutoGen
}

type repository struct {
	autoGen
}

func (receiver repository) SelectByAge() {}

func New(ctx *gin.Context) Repository {
	return &repository{autoGen{ctx}}
}

func (receiver repository) name() {
	//acc := New(receiver.ctx)
}
