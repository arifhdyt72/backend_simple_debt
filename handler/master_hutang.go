package handler

import (
	"backend_test_debt/helper"
	"backend_test_debt/initializer"
	masterhutang "backend_test_debt/master_hutang"
	"backend_test_debt/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type masterHutangHandler struct {
	masterHutangRepository masterhutang.Repository
}

func NewMasterHutangHandler(masterHutangRepository masterhutang.Repository) *masterHutangHandler {
	return &masterHutangHandler{
		masterHutangRepository: masterHutangRepository,
	}
}

func (h *masterHutangHandler) CreateMasterHutang(c *gin.Context) {
	var input masterhutang.MasterHutangInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("create master hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var masterHutang models.MasterHutang
	masterHutang.NamaHutang = input.NamaHutang
	masterHutang.JumlahMaksimal = input.JumlahMaksimal
	masterHutang.JatuhTempo = input.JatuhTempo

	result, err := h.masterHutangRepository.Create(masterHutang)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("create master hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully created master hutang", http.StatusCreated, "success", result)
	c.JSON(http.StatusCreated, response)
}

func (h *masterHutangHandler) GetAllData(c *gin.Context) {
	input := c.Request.URL.Query()
	var data masterhutang.DatatableInput

	page, _ := strconv.Atoi(input.Get("page"))
	first, _ := strconv.Atoi(input.Get("first"))
	rows, _ := strconv.Atoi(input.Get("rows"))
	pageCount, _ := strconv.Atoi(input.Get("pageCount"))

	data.Filters = input.Get("filters")
	data.Page = page
	data.First = first
	data.Rows = rows
	data.PageCount = pageCount
	data.SortField = input.Get("sortField")
	data.SortOrder = input.Get("sortOrder")

	listMasterHutang, err := h.masterHutangRepository.GetAllData(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get all master hutang failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := h.masterHutangRepository.GetCountData(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get all master hutang failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result := gin.H{"data": listMasterHutang, "countData": count}
	response := helper.ApiResponse("Successfully to get customers", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *masterHutangHandler) UpdateMasterHutang(c *gin.Context) {
	var input masterhutang.MasterHutangInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("update master hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var masterHutang models.MasterHutang
	initializer.DB.First(&masterHutang, int(input.ID))
	masterHutang.NamaHutang = input.NamaHutang
	masterHutang.JumlahMaksimal = input.JumlahMaksimal
	masterHutang.JatuhTempo = input.JatuhTempo

	result, err := h.masterHutangRepository.Update(masterHutang)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("update master hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully updated master hutang", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *masterHutangHandler) DeleteMasterHutang(c *gin.Context) {
	var input masterhutang.InputId

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("delete master hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result, err := h.masterHutangRepository.DeleteData(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("delete master hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully deleted master hutang", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *masterHutangHandler) GetAllMasterHutang(c *gin.Context) {
	var masterHutangs []models.MasterHutang
	err := initializer.DB.Find(&masterHutangs).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("get all master hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully get all master hutang", http.StatusOK, "success", masterHutangs)
	c.JSON(http.StatusOK, response)
}
