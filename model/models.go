package model

type Counter struct {
	ReferenceId string
	Sequence    int32
}

func (e ServiceError) Error() string {
	return e.Msg
}
