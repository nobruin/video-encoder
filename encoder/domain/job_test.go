package domain_test

import (
	"encoder/domain"
	testhelper "encoder/domain/test_helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValideJobIDIsNotUUID(t *testing.T) {
	video := testhelper.CreateValidVideo()
	job, _ := domain.NewJob("path", "converted", video)
	job.ID = "abc"

	err := job.Validate()

	require.NotEmpty(t, job)
	require.Error(t, err)
}

func TestValideJobIsValid(t *testing.T) {
	video := testhelper.CreateValidVideo()
	job, err := domain.NewJob("path", "converted", video)

	require.NotEmpty(t, job)
	require.Nil(t, err)
}
