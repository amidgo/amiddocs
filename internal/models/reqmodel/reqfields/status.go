package reqfields

type Status string

const (
	SEND        Status = "SEND"
	IN_PROGRESS Status = "IN_PROGRESS"
	DONE        Status = "DONE"
)
