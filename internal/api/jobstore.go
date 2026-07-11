package api

import "github.com/google/uuid"

// job store represents the DB/store where jobs will be stored

type JobStore interface {
	saveJob(*Job) error
	getJob() (*Job, error)
	deleteJob(id uuid.UUID) error
	getJobsByTypeAndStatus(type_ string, status_ string) error
}
