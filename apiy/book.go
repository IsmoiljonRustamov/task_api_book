package apiy

import (
	"books/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"),10,64) 
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse param",
		})
		return 
	}

	book, err := h.storage.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get book",
		})
		return 
	}
	ctx.JSON(http.StatusOK, book)
}

func (h *handler) CreateBook(ctx *gin.Context) {
	var b storage.Book
	err := ctx.ShouldBindJSON(&b)
	if err != nil { 
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message" : "failed to create book",
	})
	return

}
	book, err := h.storage.Create(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create book",
		})
		return
	}
	ctx.JSON(http.StatusOK, book)

}