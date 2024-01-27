package models

import (
	"time"

	"gorm.io/gorm"
)

// table for job
type Job1 struct {
	gorm.Model
	CronExpression string
	JobConfig      string
	JobName        string
	CreatedBy      int64
	UpdatedBy      int64
	Effective      bool
	Enabled        bool
	ConsumerID     int64
	EndpointID     int64
}

type JobData struct {
	gorm.Model
	JobID      int64
	ParamName  string
	ParamValue string
}

type JobTriggerEndpointData struct {
	gorm.Model
	EndpointID     string //unique index
	EndpointName   string
	EndpointConfig string
	EndpointType   string
	ConsumerID     string
	CreatedBy      int64
	UpdatedBy      int64
	Enabled        bool
}

type JobExecutions struct {
	gorm.Model
	JobID          int64
	JobConfig      string
	JobData        string
	TriggerTime    time.Time
	CompletionTime time.Time
	Status         string
}

type Job struct {
	gorm.Model
	CronExpression string `json:"cron"`
	JobData        []byte `json:"job_data"`
}
type ScheduleJob struct {
	gorm.Model
	ExecutionID int64  `json:"execution_id"`
	JobID       int64  `json:"job_id"`
	JobData     []byte `json:"job_data"`
}

type JobHistory struct {
	gorm.Model
	JobID       int64 `json:"job_id"`
	ExecutionID int64 `json:"execution_id"`
	// WorkerID    string `json:"worker_id"`
	Status string `json:"status"`
}
