package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var receipts = make(map[uuid.UUID]receiptDO)

func postReceiptsProcessHandler(c *gin.Context) {
	var newReceipt receiptDTO

	if err := c.BindJSON(&newReceipt); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var receiptUUID, err = uuid.NewRandom()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resReceipt, err := newReceipt.toReceiptDO()
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	receipts[receiptUUID] = resReceipt
	c.IndentedJSON(http.StatusCreated, receiptCreateResponse{Id: receiptUUID.String()})
}

func getReceiptPointsByIdHandler(c *gin.Context) {
	stringUUID := c.Param("id")
	if stringUUID == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	receiptUuid, err := uuid.Parse(stringUUID)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	receipt, ok := receipts[receiptUuid]
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, receiptGetPointsResponse{Points: receipt.points})
}

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.POST("/receipts/process", postReceiptsProcessHandler)
	r.GET("/receipts/:id/points", getReceiptPointsByIdHandler)
	err := r.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
