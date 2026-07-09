package api

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	Job_STATUS_PENDING   = "PENDING"
	JOB_STATUS_DONE      = "DONE"
	JOB_STATUS_EXECUTING = "EXECUTING"
)

// Job entity
type Job struct {
	Id         uuid.UUID       `json:"id"`
	JobName    string          `json:"jobName"`
	JobType    string          `json:"jobType"`
	Status     string          `json:"status"`
	CreatedAt  time.Time       `json:"createdAt"`
	FinishedAt time.Time       `json:"finishedAt"`
	Payload    json.RawMessage `json:"payload"`
}
