package handler

import (
	"backend_test_debt/helper"
	"backend_test_debt/models"
	"backend_test_debt/pembayaran"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type pembayaranHandler struct {
	pembayaranRepository pembayaran.Repository
}

func NewPembayaranHandler(pembayaranRepository pembayaran.Repository) *pembayaranHandler {
	return &pembayaranHandler{
		pembayaranRepository: pembayaranRepository,
	}
}

func (h *pembayaranHandler) CreatePembayaran(c *gin.Context) {
	var input pembayaran.PembayaranInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to upload attachment", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	dateTransaksi, err := time.Parse("02/01/2006", input.TglTransaksi)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("create pembayaran failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to upload file 1", http.StatusUnprocessableEntity, "error", data)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	titleText := strings.ReplaceAll(strings.ToLower(file.Filename), " ", "-")
	path := fmt.Sprintf("images/%d-%s", time.Now().Unix(), titleText)
	// Create the file with the desired permissions
	out, err := os.OpenFile("./"+path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to create file 2", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	defer out.Close()

	err = c.SaveUploadedFile(file, "./"+path)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to upload file 3", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Set file permissions to 0644
	err = os.Chmod(path, 0755)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to set file permissions", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var pembayaran models.Pembayaran
	pembayaran.HutangId = uint(input.HutangId)
	pembayaran.TotalDibayar = input.TotalDibayar
	pembayaran.TglTransaksi = dateTransaksi
	pembayaran.BuktiPembayaran = path

	result, err := h.pembayaranRepository.Create(pembayaran)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to create pembayaran", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully created pembayaran", http.StatusCreated, "success", result)
	c.JSON(http.StatusCreated, response)
}
