package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	appusecases "github.com/sousair/apollo-cloud/internal/application/usecases"
	"github.com/sousair/apollo-cloud/internal/infra/providers"
	gormrepositories "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm"
	gormmodels "github.com/sousair/apollo-cloud/internal/infra/repositories/gorm/models"
	"github.com/sousair/apollo-cloud/internal/infra/repositories/s3"
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

	var (
		awsAccessKeyID     = os.Getenv("AWS_ACCESS_KEY_ID")
		awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
		awsRegion          = os.Getenv("AWS_REGION")
		awsEndpoint        = os.Getenv("AWS_ENDPOINT")
	)

	awsSession, err := session.NewSession(&aws.Config{
		Region:           aws.String(awsRegion),
		Credentials:      credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(awsEndpoint),
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var (
		privateBucketName = os.Getenv("AWS_S3_PRIVATE_BUCKET_NAME")
		publicBucketName  = os.Getenv("AWS_S3_PUBLIC_BUCKET_NAME")
	)

	uuidProvider := uuidv4.NewUuidV4Provider()
	ownerRepository := gormrepositories.NewGormOwnerRepository(db)
	fileRepository := s3.NewS3FileRepository(awsSession, privateBucketName, publicBucketName)
	musicRepository := gormrepositories.NewGormMusicRepository(db)
	albumRepository := gormrepositories.NewGormAlbumRepository(db)

	createOwnerUsecase := appusecases.NewCreateOwnerUsecase(uuidProvider, ownerRepository)
	createMusicUsecase := appusecases.NewCreateMusicUsecase(uuidProvider, fileRepository, musicRepository)
	releaseAlbumUsecase := appusecases.NewReleaseAlbumUsecase(fileRepository, uuidProvider, albumRepository, createMusicUsecase)
	getAlbumUsecase := appusecases.NewGetAlbumUsecase(albumRepository)
	getMusicUsecase := appusecases.NewGetMusicUsecase(musicRepository)

	validator := validator.New()

	createOwnerHandler := httphandlers.NewCreateOwnerHttpHandler(validator, createOwnerUsecase)
	createMusicHandler := httphandlers.NewCreateMusicHttpHandler(validator, createMusicUsecase)
	releaseAlbumHandler := httphandlers.NewReleaseAlbumHttpHandler(validator, releaseAlbumUsecase)
	getAlbumHandler := httphandlers.NewGetAlbumHttpHandler(validator, getAlbumUsecase)
	getMusicHandler := httphandlers.NewGetMusicHttpHandler(validator, getMusicUsecase)

	e := echo.New()

	e.POST("/owners", createOwnerHandler.Handle)

	e.POST("/musics", createMusicHandler.Handle)
	e.GET("/musics/:id", getMusicHandler.Handle)

	e.POST("/albums", releaseAlbumHandler.Handle)
	e.GET("/albums/:id", getAlbumHandler.Handle)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))))
}
