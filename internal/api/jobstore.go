package api

import "github.com/google/uuid"

// job store interface

type JobStore interface {
	saveJob(*Job) error
	getJob() (*Job, error)
	deleteJob(id uuid.UUID) error
	getJobsByTypeAndStatus(type_ string, status_ string) error
	deleteJobsByStatus(string) error
	deleteJobsByStatusAndType(type_, status_ string) error
}
