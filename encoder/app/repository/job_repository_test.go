package repository_test

import (
	"encoder/app/repository"
	"encoder/domain"
	testhelper "encoder/domain/test_helper"
	"encoder/infra/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertJobRepository(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()
	video := testhelper.CreateValidVideo()
	videoRepository := repository.VideoRepositoryDb{Db: db}
	videoRepository.Insert(video)

	job, err := domain.NewJob("path", "pending", video)
	require.Nil(t, err)

	repo := repository.JobRepositoryDb{Db: db}
	_, err = repo.Insert(job)

	require.Nil(t, err)

	jobFinded, err := repo.Find(job.ID)

	require.Nil(t, err)
	require.Equal(t, job.ID, jobFinded.ID)
	require.Equal(t, job.VideoID, video.ID)
}

func TestUpdateJobRepository(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()
	video := testhelper.CreateValidVideo()
	videoRepository := repository.VideoRepositoryDb{Db: db}
	videoRepository.Insert(video)

	job, err := domain.NewJob("path", "pending", video)
	require.Nil(t, err)

	repo := repository.JobRepositoryDb{Db: db}

	_, err = repo.Insert(job)
	require.Nil(t, err)

	job.Status = "Complete"

	_, err = repo.Update(job)
	require.Nil(t, err)

	jobFinded, err := repo.Find(job.ID)

	require.Nil(t, err)
	require.Equal(t, job.ID, jobFinded.ID)
	require.Equal(t, jobFinded.Status, "Complete")
	require.Equal(t, job.VideoID, video.ID)
}
