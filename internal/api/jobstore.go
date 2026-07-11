package api

// job store represents the DB/store where jobs will be stored

type JobStore interface {
	saveJob() error
	getJob() (*Job, error)
	deleteJob() error
}
