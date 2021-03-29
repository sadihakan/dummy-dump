package errors

type Err string

const (
	ConfigUserNil            Err = "user can not be nil"
	ConfigSourceNil          Err = "select source"
	ConfigPathNotExist       Err = "path is not exist"
	ConfigDbNotExist         Err = "DB can not be nil"
	ConfigBinaryPathNotExist Err = "binary path can not be nil"
	ConfigMethodError        Err = "select method"
)
