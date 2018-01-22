package model

const (
	StreamStateCreated     = "created"
	StreamStateActive      = "active"
	StreamStateInterrupted = "interrupted"
	StreamStateFinish      = "finish"
)

type Stream struct {
	ID    string
	State string
}

func (s *Stream) CanCahngeStateTo(newState string) bool {
	if newState == StreamStateActive {
		if s.State != StreamStateCreated &&
			s.State != StreamStateInterrupted {
			return false
		}
		return true
	}

	return false
}

func (s *Stream) Copy() *Stream {
	copy := *s
	return &copy
}
