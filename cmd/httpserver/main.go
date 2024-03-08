package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	appusecases "github.com/sousair/apollo-cloud/internal/application/usecases"
	"github.com/sousair/apollo-cloud/internal/infra/providers"
	gormrepositories "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm"
	gormmodels "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm/models"
	httphandlers "github.com/sousair/apollo-cloud/internal/presentation/http/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Error loading .env file")
	}

	postgresConnectionURL := os.Getenv("POSTGRES_CONNECTION_URL")

	db, err := gorm.Open(postgres.Open(postgresConnectionURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	err = db.AutoMigrate(gormmodels.OwnerModel{}, gormmodels.AlbumModel{}, gormmodels.MusicModel{})

	if err != nil {
		fmt.Println("Error migrating models")
		panic(err)
	}

	uuidProvider := uuidv4.NewUuidV4Provider()
	ownerRepository := gormrepositories.NewGormOwnerRepository(db)

	createOwnerUsecase := appusecases.NewCreateOwnerUseCase(uuidProvider, ownerRepository)

	validator := validator.New()
	createOwnerHandler := httphandlers.NewCreateOwnerHttpHandler(createOwnerUsecase, validator)

	e := echo.New()

	e.POST("/owners", createOwnerHandler.Handle)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))))
}
