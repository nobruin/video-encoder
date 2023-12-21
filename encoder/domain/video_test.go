package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValideVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestValideVideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "abc"
	err := video.Validate()

	require.Error(t, err)
}

func TestValideVideoIsValid(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceId = "ID"
	video.FilePath = "../path/"
	video.CreatedAT = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
