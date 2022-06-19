package orders

import (
	"encoding/json"
	"github.com/MahmoudMekki/MDS-task/pkg/models"
	"github.com/MahmoudMekki/MDS-task/pkg/rabbit"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog/log"
	"net/http"
)

func BulkUpload(ctx *gin.Context) {
	filePtr, err := ctx.FormFile("file")
	if err != nil || filePtr.Size <= 0 {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	file, err := filePtr.Open()
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while reading file"})
		return
	}
	defer file.Close()
	var orders []models.OrderCSV
	err = gocsv.Unmarshal(file, &orders)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while reading file"})
		return
	}
	msg, err := json.Marshal(orders)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while reading file"})
		return
	}
	err = rabbit.Produce(models.OrdersMQTopic, msg)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while reading file"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "file is being processed"})
}
