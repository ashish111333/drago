package api

// job store represents the DB/store where jobs will be stored

type JobStore interface {
	saveJob(*Job) error
	getJob() (*Job, error)
	deleteJob() error
	getJobsByTypeAndStatus(string, string) error
}
