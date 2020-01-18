/*
@Time : 2020/1/18 10:19
@Author : Minus4
*/
package service

type WarpedError struct {
	Err  error
	Code int
}

func (e WarpedError) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

//NewServiceError
func NewServiceError(code int, err error) *WarpedError {
	return &WarpedError{
		Err:  err,
		Code: code,
	}
}

const (
	StatusOK = iota + 1
	StatusDeleted
)

const (
	GroupAdminUser   = 1
	GroupManagerUser = 2
)
