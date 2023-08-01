package account

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewApi() *Handler {
	return &Handler{}
}

func (*Handler) Create(ctx *gin.Context) {

}

func (*Handler) Remove(ctx *gin.Context) {

}

func (*Handler) Modify(ctx *gin.Context) {

}

func (*Handler) ListPage(ctx *gin.Context) {

}

func (*Handler) Detail(ctx *gin.Context) {

}
