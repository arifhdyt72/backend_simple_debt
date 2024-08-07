package main

import (
	"backend_test_debt/initializer"
	"backend_test_debt/models"
	"os"
)

func init() {
	// INIT ENV VARIABLE
	initializer.LoadEnv()

	// CONNECT TO DATABASE
	initializer.ConnectDB()

	// SET TIMEZONE
	os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	initializer.DB.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.MasterHutang{},
		&models.Hutang{},
		&models.Pembayaran{},
	)
}
