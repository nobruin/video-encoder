package repository

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{Db: db}
}

func (vr VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	err := vr.Db.Create(video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (vr VideoRepositoryDb) Find(id string) (*domain.Video, error) {
	var video domain.Video

	vr.Db.Preload("Job").First(&video, "id = ?", id)
	if video.ID == "" {
		return nil, fmt.Errorf("video not find")
	}

	return &video, nil
}
