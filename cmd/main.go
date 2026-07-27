package cmd

import (
	"time"
	"user-service/config"
	"user-service/database/seeders"
	"user-service/domain/models"
	"user-service/repositories"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use: "serve",
	Short: "Start the server",
	Run: func(c *cobra.Command, args []string) {
		_ = godotenv.Load()
		config.Init()
		db, err := config.InitDatabase()
		if err != nil {
			panic(err)
		}

		// set lokasi
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			panic(err)
		}
		time.Local = loc

		err = db.AutoMigrate(
			&models.Role{},
			&models.User{},
		)
		if err != nil {
			panic(err)
		}

		seeders.NewSeederRegistry(db).Run()
		repository := repositories.NewRepositoryRegistry(db)
		service := services.NewServiceRegistry(repository)
		controller := controllers.NewControllerRegistry(service)
	}
}
