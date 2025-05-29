package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreatePurchase(ctx *gin.Context) {

}

func (server *Server) DatatablePurchase(c *gin.Context) {

	rowsResult, err := server.store.PurchaseQueries.Datatable(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusAccepted, rowsResult)
}
