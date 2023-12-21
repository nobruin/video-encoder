package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValideJobIDIsNotUUID(t *testing.T) {
	video := createValidVideo()
	Job, err := domain.NewJob("path", "converted", video)
	Job.ID = "abc"

	require.NotEmpty(t, Job)
	require.Error(t, err)
}

func TestValideJobIsValid(t *testing.T) {
	video := createValidVideo()
	Job, err := domain.NewJob("path", "converted", video)

	require.NotEmpty(t, Job)
	require.Nil(t, err)
}

func createValidVideo() *domain.Video {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "/path"
	video.CreatedAT = time.Now()

	return video
}
