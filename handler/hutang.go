package handler

import (
	"backend_test_debt/helper"
	"backend_test_debt/hutang"
	"backend_test_debt/initializer"
	"backend_test_debt/models"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

type hutangHandler struct {
	hutangRepository hutang.Repository
}

func NewHutangHandler(hutangRepository hutang.Repository) *hutangHandler {
	return &hutangHandler{
		hutangRepository: hutangRepository,
	}
}

func (h *hutangHandler) CreateHutang(c *gin.Context) {
	var input hutang.HutangInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("create hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	dateTransaksi, err := time.Parse("02/01/2006", input.TglTransaksi)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("create hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var hutang models.Hutang
	hutang.CustomerId = uint(input.CustomerId)
	hutang.MasterHutangId = uint(input.MasterHutangId)
	hutang.TglTransaksi = dateTransaksi
	hutang.JumlahHutang = input.JumlahHutang

	result, err := h.hutangRepository.Create(hutang)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("create hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully created hutang", http.StatusCreated, "success", result)
	c.JSON(http.StatusCreated, response)
}

func (h *hutangHandler) GetAllData(c *gin.Context) {
	input := c.Request.URL.Query()
	var data hutang.DatatableInput

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

	listHutang, err := h.hutangRepository.GetAllData(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get all hutang failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := h.hutangRepository.GetCountData(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get all hutang failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result := gin.H{"data": listHutang, "countData": count}
	response := helper.ApiResponse("Successfully to get all hutang", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *hutangHandler) UpdateHutang(c *gin.Context) {
	var input hutang.HutangInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("update hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	dateTransaksi, err := time.Parse("02/01/2006", input.TglTransaksi)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("update hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var hutang models.Hutang
	initializer.DB.First(&hutang, int(input.ID))

	hutang.CustomerId = uint(input.CustomerId)
	hutang.MasterHutangId = uint(input.MasterHutangId)
	hutang.TglTransaksi = dateTransaksi
	hutang.JumlahHutang = input.JumlahHutang

	result, err := h.hutangRepository.Update(hutang)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("update hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully updated hutang", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *hutangHandler) DeleteHutang(c *gin.Context) {
	var input hutang.InputId

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("delete hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result, err := h.hutangRepository.DeleteData(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("delete hutang failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully deleted hutang", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *hutangHandler) ExportHutang(c *gin.Context) {
	listData, err := h.hutangRepository.GetAllReport()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("export report failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	xlsx := excelize.NewFile()

	styleFillBorder, err := xlsx.NewStyle(`{
		"fill": {
			"type": "pattern",
			"color": ["#00FF00"],
			"pattern": 1
		},
		"border" : [
			{ "type" : "left" , "color" : "000000" , "style" : 1 },
			{ "type" : "top" , "color" : "000000" , "style" : 1 },
			{ "type" : "bottom" , "color" : "000000" , "style" : 1 },
			{ "type" :"right" , "color" : "000000" , "style": 1 }
		]
	}`)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("export report failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	styleBorder, err := xlsx.NewStyle(`{
		"border" : [
			{ "type" : "left" , "color" : "000000" , "style" : 1 },
			{ "type" : "top" , "color" : "000000" , "style" : 1 },
			{ "type" : "bottom" , "color" : "000000" , "style" : 1 },
			{ "type" :"right" , "color" : "000000" , "style": 1 }
		]
	}`)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("export report failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	sheet1Name := "Rekapitulasi hutang"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Nama customer")
	xlsx.SetCellValue(sheet1Name, "B1", "Master hutang")
	xlsx.SetCellValue(sheet1Name, "C1", "Tanggal transaksi")
	xlsx.SetCellValue(sheet1Name, "D1", "Tanggal jatuh tempo")
	xlsx.SetCellValue(sheet1Name, "E1", "Jumlah hutang")
	xlsx.SetCellValue(sheet1Name, "F1", "Total dibayar")
	xlsx.SetCellValue(sheet1Name, "G1", "Sisa tagihan")
	xlsx.SetCellValue(sheet1Name, "H1", "Status")
	xlsx.SetCellStyle(sheet1Name, "A1", "A1", styleFillBorder)
	xlsx.SetCellStyle(sheet1Name, "B1", "B1", styleFillBorder)
	xlsx.SetCellStyle(sheet1Name, "C1", "C1", styleFillBorder)
	xlsx.SetCellStyle(sheet1Name, "D1", "D1", styleFillBorder)
	xlsx.SetCellStyle(sheet1Name, "E1", "E1", styleFillBorder)
	xlsx.SetCellStyle(sheet1Name, "F1", "F1", styleFillBorder)
	xlsx.SetCellStyle(sheet1Name, "G1", "G1", styleFillBorder)
	xlsx.SetCellStyle(sheet1Name, "H1", "H1", styleFillBorder)

	err = xlsx.AutoFilter(sheet1Name, "A1", "B1", "")
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("export report failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	rowMaster := 2
	for _, data := range listData {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", rowMaster), data.Customer.Name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", rowMaster), data.MasterHutang.NamaHutang)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", rowMaster), data.TglTransaksi.Format("02/01/2006"))
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", rowMaster), data.TglJatuhTempo.Format("02/01/2006"))
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", rowMaster), data.JumlahHutang)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", rowMaster), data.TotalDibayar)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", rowMaster), data.SisaTagihan)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", rowMaster), data.Status)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("A%d", rowMaster), fmt.Sprintf("A%d", rowMaster), styleBorder)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("B%d", rowMaster), fmt.Sprintf("B%d", rowMaster), styleBorder)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("C%d", rowMaster), fmt.Sprintf("C%d", rowMaster), styleBorder)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("D%d", rowMaster), fmt.Sprintf("D%d", rowMaster), styleBorder)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("E%d", rowMaster), fmt.Sprintf("E%d", rowMaster), styleBorder)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("F%d", rowMaster), fmt.Sprintf("F%d", rowMaster), styleBorder)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("G%d", rowMaster), fmt.Sprintf("G%d", rowMaster), styleBorder)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("H%d", rowMaster), fmt.Sprintf("H%d", rowMaster), styleBorder)
		if len(data.Pembayaran) > 0 {
			rowMaster = rowMaster + 1
			for _, detail := range data.Pembayaran {
				xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", rowMaster), detail.TglTransaksi.Format("02/01/2006"))
				xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", rowMaster), detail.TotalDibayar)
				xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", rowMaster), detail.Status)
				xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("C%d", rowMaster), fmt.Sprintf("C%d", rowMaster), styleBorder)
				xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("F%d", rowMaster), fmt.Sprintf("F%d", rowMaster), styleBorder)
				xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("H%d", rowMaster), fmt.Sprintf("H%d", rowMaster), styleBorder)
				rowMaster = rowMaster + 1
				continue
			}
		} else {
			rowMaster = rowMaster + 1
			continue
		}
	}

	filename := fmt.Sprintf("rekapitulasi-hutang-%d.xlsx", time.Now().Unix())
	path := fmt.Sprintf("report/%s", filename)
	savePath := fmt.Sprintf("./report/%s", filename)

	err = xlsx.SaveAs(savePath)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to Export Report", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fileUploaded, err := os.Open(savePath)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("export report failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	defer fileUploaded.Close()

	baseUrl := os.Getenv("BASE_URL")
	fullUrl := fmt.Sprintf("%s%s", baseUrl, path)
	dataResult := gin.H{"url_report": fullUrl}
	response := helper.ApiResponse("successfully to get report", http.StatusOK, "success", dataResult)
	c.JSON(http.StatusOK, response)
}
