package main

import (
	"fmt"
	"os"

	"github.com/rudysacosta/go-and-mysql-dockerizing/config"
	"github.com/rudysacosta/go-and-mysql-dockerizing/db"
)

func main() {

	// Create database instance
	database, err := db.NewDatabase(config.NewConfig().Database)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}

	defer database.Close()

	// Run migrations
	if err := db.RunMigrations(database, "/app/migrations"); err != nil {
		fmt.Println("Failed to run migrations:", err)
		os.Exit(1)
	}

	fmt.Println("Migrations completed successfully!")

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()
}
