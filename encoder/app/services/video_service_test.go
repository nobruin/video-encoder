package services_test

import (
	"encoder/app/repository"
	"encoder/app/services"
	"encoder/domain"
	"encoder/infra/database"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
}

func prepare(videoPath string) (domain.Video, repository.VideoRepository) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = videoPath
	video.CreatedAT = time.Now()

	repo := repository.VideoRepositoryDb{Db: db}
	return *video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare("video_exemplo.mp4")

	videoService := services.NewVideoService()
	videoService.Video = &video
	videoService.VideoRepository = &repo

	err := videoService.Downlod("test")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)
}
