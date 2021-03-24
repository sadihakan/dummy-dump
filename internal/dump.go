package internal

// Dump ...
type Dump interface {
	Check() error
	Export(binaryPath string, user string, database string) error
	Import(binaryPath string, user string, path string) error
}