package api

import (
	"context"
	"errors"
	"os"

	"github.com/google/uuid"
)

// job store interface

type JobStore interface {
	SaveJob(context.Context, *Job) error
	GetJob(context.Context) (*Job, error)
	DeleteJob(ctx context.Context, id uuid.UUID) error
	GetJobsByTypeAndStatus(ctx context.Context, type_ string, status_ string) error
	DeleteJobsByStatus(context.Context, string) error
	DeleteJobsByStatusAndType(ctx context.Context, type_, status_ string) error
}

type PgJobStore struct{}

func (pgs *PgJobStore) Connect() error {

	conn_url := os.Getenv("PG_CONN_URL")
	if conn_url == "" {
		return errors.New("connection url is empty: can't connect to datastore")
	}

	return nil

}

func (pgs *PgJobStore) SaveJob() {

}
func (pgs *PgJobStore) GetJob() {}

func (pgs *PgJobStore) DeleteJob() {

}

func (pgs *PgJobStore) GetJobsByTypeAndStatus()    {}
func (pgs *PgJobStore) DeleteJobsByStatusAndType() {}
func (pgs *PgJobStore) DeleteJobsByStatus()        {}
