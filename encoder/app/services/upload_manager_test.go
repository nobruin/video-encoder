package services_test

import (
	"encoder/app/services"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
}

func TestVideoServiceUpload(t *testing.T) {
	video, repo := prepare("video_exemplo.mp4")

	videoService := services.NewVideoService()
	videoService.Video = &video
	videoService.VideoRepository = &repo

	err := videoService.Downlod("test")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	videoUpload := services.NewVideoUpload()
	videoUpload.OutputBucket = "test"
	videoUpload.VideoPath = os.Getenv("localStoragePath") + "/" + video.ID

	doneUpload := make(chan string)
	videoUpload.ProcessUpload(50, doneUpload)

	result := <-doneUpload
	require.Equal(t, result, "upload completed")

}
