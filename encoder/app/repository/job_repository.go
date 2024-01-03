package repository

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepositoryDb {
	return &JobRepositoryDb{Db: db}
}

func (vr JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {
	if job.ID == "" {
		job.ID = uuid.NewV4().String()
	}

	err := vr.Db.Create(job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (vr JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job

	vr.Db.Preload("Video").First(&job, "id = ?", id)
	if job.ID == "" {
		return nil, fmt.Errorf("job not find")
	}

	return &job, nil
}

func (vr JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := vr.Db.Save(&job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}
