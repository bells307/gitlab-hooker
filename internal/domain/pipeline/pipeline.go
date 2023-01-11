package pipeline

type Pipeline struct {
	Project string
	Branch  string
	Status  PipelineStatus
}

type PipelineStatus string

const (
	Created            PipelineStatus = "created"
	WaitingForResource                = "waiting_for_resource"
	Preparing                         = "preparing"
	Pending                           = "pending"
	Running                           = "running"
	Success                           = "success"
	Failed                            = "failed"
	Canceled                          = "canceled"
	Skipped                           = "skipped"
	Manual                            = "manual"
	Schedules                         = "scheduled"
)
