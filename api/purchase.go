package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luke_design_pattern/internal"
)

// CreatePurchase godoc
// @Summary      Create a new Purchase
// @Description  Create a new Purchase for detail
// @Tags         purchase
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "purchase ID"
// @Success      200  {object}  map[string]string
// @Router       /purchase/create [post]
func (server *Server) CreatePurchase(c *gin.Context) {

	ctx := c.Request.Context()

	var reqPrcHistory internal.CreatePurchaseHistoryParams
	var err error
	var respPrcHistory internal.PurchaseHistory

	err = c.ShouldBindJSON(&reqPrcHistory)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	// generate automatically
	reqPrcHistory.PurchaseNumber = internal.GenerateRandomTrxNumber(reqPrcHistory.CustomerID)
	log.Printf("check req %v", reqPrcHistory)

	respPrcHistory, err = server.store.CreatePurchaseHistoryTx(ctx, reqPrcHistory)
	if err != nil {
		c.JSON(http.StatusConflict, errorResponse(err, http.StatusConflict))
		return
	}
	c.JSON(http.StatusAccepted, successResponse(
		"success create purchase history",
		respPrcHistory,
		http.StatusAccepted,
	))

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

type DeletePurchaseSuccessResponse struct {
	Message    string                               `json:"message"`
	Data       internal.DeletePurchaseItemsTxParams `json:"data"`
	StatusCode int                                  `json:"status_code" example:"200"`
}

// DeletePurchase godoc
// @Summary      Delete Purchase
// @Description  Deletion proccess can only be done if it is not completed
// @Tags         purchase
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "purchase ID"
// @Success      200  {object}  DeletePurchaseSuccessResponse "Success delete purchase"
// @Failure 	 400 {object} ErrorResponse "Bad Request"
// @Router       /purchase/delete/{id} [delete]
func (server *Server) DeletePurchase(c *gin.Context) {
	// uri
	// purchase/delete/:prc_number
	var params DeletePRCParams

	err := c.ShouldBindUri(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}
	args := internal.DeletePurchaseItemsTxParams{
		PurchaseNumber: params.PrcNumber,
	}

	log.Printf("check param %s", params.PrcNumber)
	err = server.store.DeletePurchaseTx(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	resp := successResponse(
		fmt.Sprintf("succes delete purchase number: %s", params.PrcNumber),
		internal.DeletePurchaseItemsTxParams{
			PurchaseNumber: params.PrcNumber,
		},
		http.StatusAccepted,
	)
	c.JSON(http.StatusAccepted, resp)
}
