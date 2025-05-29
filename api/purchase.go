package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luke_design_pattern/internal"
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

type DeletePRCParams struct {
	PrcNumber string `uri:"prc_number" binding:"required"`
}

func (server *Server) DeletePurchase(c *gin.Context) {
	// uri
	// purchase/delete/:prc_number
	var params DeletePRCParams

	err := c.ShouldBindUri(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := internal.DeletePurchaseItemsTxParams{
		PurchaseNumber: params.PrcNumber,
	}

	log.Printf("check param %s", params.PrcNumber)
	err = server.store.DeletePurchaseTx(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusAccepted, fmt.Sprintf("succes delete purchase number: %s", params.PrcNumber))
}
