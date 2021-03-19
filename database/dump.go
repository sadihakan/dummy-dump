package database

type Dump interface {
	Check() error
	Export(user string, database string) error
	Import(user string, path string) error
}