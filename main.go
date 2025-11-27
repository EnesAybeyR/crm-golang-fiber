package main

import (
	"fmt"

	"github.com/EnesAybeyR/crm-golang-fiber.git/database"
	"github.com/EnesAybeyR/crm-golang-fiber.git/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/lead", lead.GetLeads)
	app.Get("/api/lead/:id", lead.GetLead)
	app.Post("/api/lead", lead.NewLead)
	app.Delete("/api/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/crm_db?charset=utf8&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	fmt.Println("connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
