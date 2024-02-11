package services

import (
	"encoder/app/repository"
	"encoder/domain"
)

type JobService struct {
	Job           *domain.Job
	JobRepository repository.JobRepository
	VideoService  VideoService
}

func (j *JobService) Start() error {
	return nil
}

func (j *JobService) changeJobStatus(status string) error {
	var err error

	j.Job.Status = status
	j.Job, err = j.JobRepository.Update(j.Job)

	if err != nil {
		return j.failJob(err)
	}

	return nil
}

func (j *JobService) failJob(errReceived error) error {
	j.Job.Status = "FAILED"
	j.Job.Error = errReceived.Error()

	_, err := j.JobRepository.Update(j.Job)
	if err != nil {
		return err
	}
	return errReceived
}
