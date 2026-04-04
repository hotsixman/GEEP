package types

// PM
type PMInterface interface {
	NewProcess(name string, udsServer UDSServerInterface, args ...string) error
	Input(name string, command string)
}

// Logger
type LoggerInterface interface {
}

// UDS
type UDSServerInterface interface {
	Broadcast(name string, JSON []byte)
}

// Error
type UndefinedProcessNameError struct{}

func (_ UndefinedProcessNameError) Error() string {
	return "'name' field is not defined."
}
