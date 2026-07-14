package api

import (
	"context"
	"errors"
	"os"

	"github.com/google/uuid"
	ppool "github.com/jackc/pgx/v5/pgxpool"
)

// job store interface

type JobStore interface {
	SaveJob(context.Context, *Job) error
	GetJob(context.Context) (*Job, error)
	DeleteJob(ctx context.Context, id uuid.UUID) error
	GetJobsByTypeAndStatus(ctx context.Context, type_ string, status_ string) error
	DeleteJobsByStatus(context.Context, string) error
	DeleteJobsByStatusAndType(ctx context.Context, type_, status_ string) error
	// close all resources
	Close() error
}

type PgJobStore struct {
	pool *ppool.Pool
}

func (pgs *PgJobStore) Connect() (*ppool.Pool, error) {

	conn_url := os.Getenv("PG_CONN_URL")
	if conn_url == "" {
		return nil, errors.New("connection url is empty: can't connect to datastore")
	}

	pgxPool, err := ppool.New(context.Background(), conn_url)

	if err != nil {
		return nil, err
	}

	return pgxPool, nil

}

func (pgs *PgJobStore) SaveJob() {

}
func (pgs *PgJobStore) GetJob() {}

func (pgs *PgJobStore) DeleteJob() {

}

func (pgs *PgJobStore) GetJobsByTypeAndStatus()    {}
func (pgs *PgJobStore) DeleteJobsByStatusAndType() {}
func (pgs *PgJobStore) DeleteJobsByStatus()        {}
func (pgs *PgJobStore) Close() {

}
