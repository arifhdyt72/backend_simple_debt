package main

import (
	"backend_test_debt/customer"
	"backend_test_debt/handler"
	"backend_test_debt/hutang"
	"backend_test_debt/initializer"
	masterhutang "backend_test_debt/master_hutang"
	"backend_test_debt/middleware"
	"backend_test_debt/pembayaran"
	"backend_test_debt/user"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	// INIT ENV VARIABLE
	initializer.LoadEnv()

	// CONNECT TO DATABASE
	initializer.ConnectDB()

	os.Setenv("TZ", "Asia/Jakarta")

	// INIT CUSTOM LOG
	initializer.InitLogger()
}

func main() {
	userRepositoy := user.NewRepository()
	jwtService := middleware.NewJWTService()
	userHandler := handler.NewUserHandler(userRepositoy, jwtService)

	customerRepository := customer.NewRepository()
	customerHandler := handler.NewCustomerHandler(customerRepository)

	masterHutangRepository := masterhutang.NewRepository()
	masterHutangHandler := handler.NewMasterHutangHandler(masterHutangRepository)

	hutangRepository := hutang.NewRepository()
	hutangHandler := handler.NewHutangHandler(hutangRepository)

	pembayaranRepository := pembayaran.NewRepository()
	pembayaranHandler := handler.NewPembayaranHandler(pembayaranRepository)

	rand.Seed(time.Now().UnixNano())
	// CREATE HTTP SERVER USING GIN
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api/v1")
	r.POST("/api/login", userHandler.LoginHandler)
	{
		api.Use(middleware.AuthMiddlware(jwtService))
		api.GET("/user", userHandler.GetAllData)
		api.POST("/user", userHandler.CreateUser)
		api.PUT("/user", userHandler.UpdateUser)
		api.DELETE("/user/:id", userHandler.DeleteUser)

		api.POST("/customer", customerHandler.CreateCustomer)
		api.GET("/customer", customerHandler.GetAllData)
		api.GET("/customer_all", customerHandler.GetAllCustomer)
		api.PUT("/customer", customerHandler.UpdateCustomer)
		api.DELETE("/customer/:id", customerHandler.DeleteCustomer)

		api.POST("/master_hutang", masterHutangHandler.CreateMasterHutang)
		api.GET("/master_hutang", masterHutangHandler.GetAllData)
		api.GET("/master_hutang_all", masterHutangHandler.GetAllMasterHutang)
		api.PUT("/master_hutang", masterHutangHandler.UpdateMasterHutang)
		api.DELETE("/master_hutang/:id", masterHutangHandler.DeleteMasterHutang)

		api.POST("/hutang", hutangHandler.CreateHutang)
		api.GET("/hutang", hutangHandler.GetAllData)
		api.GET("/hutang_report", hutangHandler.ExportHutang)
		api.PUT("/hutang", hutangHandler.UpdateHutang)
		api.DELETE("/hutang/:id", hutangHandler.DeleteHutang)

		api.POST("/pembayaran", pembayaranHandler.CreatePembayaran)
	}

	r.Static("/images", "images")
	r.Static("/report", "report")

	r.GET("/info", func(c *gin.Context) {
		routeInfo := r.Routes()
		fmt.Println(routeInfo[0].Method)
		paths := make([]string, len(routeInfo))

		for i, k := range routeInfo {
			paths[i] = k.Method + " : " + k.Path
		}

		c.JSON(http.StatusOK, gin.H{
			"routes": paths,
		})
	})

	r.Run("0.0.0.0")
}
