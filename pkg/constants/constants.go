package constants

const (
	SubmissionQueue = "submission_queue"
)

const (
	SubmissionPending   = "PENDING"
	SubmissionRunning   = "RUNNING"
	SubmissionCompleted = "COMPLETED"
	SubmissionFailed    = "FAILED"
)

const (
	VerdictAccepted         = "ACCEPTED"
	VerdictWrongAnswer      = "WRONG_ANSWER"
	VerdictTLE              = "TIME_LIMIT_EXCEEDED"
	VerdictMLE              = "MEMORY_LIMIT_EXCEEDED"
	VerdictRuntimeError     = "RUNTIME_ERROR"
	VerdictCompilationError = "COMPILATION_ERROR"
)
