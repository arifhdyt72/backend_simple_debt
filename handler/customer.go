package handler

import (
	"backend_test_debt/customer"
	"backend_test_debt/helper"
	"backend_test_debt/initializer"
	"backend_test_debt/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerRepository customer.Repository
}

func NewCustomerHandler(customerRepository customer.Repository) *customerHandler {
	return &customerHandler{
		customerRepository: customerRepository,
	}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	var input customer.CustomerInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("create customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var customer models.Customer
	truest := true
	customer.Name = input.Name
	customer.Email = input.Email
	customer.PhoneNumber = input.PhoneNumber
	customer.Address = input.Address
	customer.Status = &truest

	result, err := h.customerRepository.Create(customer)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("create customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully created customer", http.StatusCreated, "success", result)
	c.JSON(http.StatusCreated, response)
}

func (h *customerHandler) GetAllData(c *gin.Context) {
	input := c.Request.URL.Query()
	var data customer.DatatableInput

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

	listCustomer, err := h.customerRepository.GetAllData(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get all customer failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := h.customerRepository.GetCountData(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get all customer failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result := gin.H{"data": listCustomer, "countData": count}
	response := helper.ApiResponse("Successfully to get customers", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	var input customer.CustomerInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("update customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var customer models.Customer
	initializer.DB.First(&customer, int(input.ID))
	customer.Name = input.Name
	customer.Email = input.Email
	customer.PhoneNumber = input.PhoneNumber
	customer.Address = input.Address

	result, err := h.customerRepository.Update(customer)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("update customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully updated customer", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) DeleteCustomer(c *gin.Context) {
	var input customer.InputId

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("delete customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result, err := h.customerRepository.DeleteData(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("delete customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully deleted customer", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) GetAllCustomer(c *gin.Context) {
	var customers []models.Customer
	err := initializer.DB.Find(&customers).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("get all customer failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully get all customer", http.StatusOK, "success", customers)
	c.JSON(http.StatusOK, response)
}
