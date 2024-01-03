package repository_test

import (
	"encoder/app/repository"
	"encoder/domain"
	"encoder/infra/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValideVideoRepository(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAT = time.Now()

	repo := repository.VideoRepositoryDb{Db: db}
	repo.Insert(video)
	videoFinded, err := repo.Find(video.ID)

	require.Nil(t, err)
	require.Equal(t, video.ID, videoFinded.ID)
}
