package testhelper

import (
	"encoder/domain"
	"time"

	uuid "github.com/satori/go.uuid"
)

func CreateValidVideo() *domain.Video {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "/path"
	video.CreatedAT = time.Now()

	return video
}
