package jalmot

type Jalmot interface {
	error
	From() Location
}

func New(message string) Jalmot {
	return &impl{
		message:  message,
		location: newLocation(2),
	}
}

type impl struct {
	message  string
	location Location
}

func (i *impl) Error() string {
	if i == nil {
		return "nil"
	}
	return i.message
}

func (i *impl) From() Location {
	if i == nil {
		return nil
	}
	return i.location
}
