package e

const (
	SUCCESS          = 200
	ERROR            = 500
	ErrInvalidParams = 400

	ErrAuthCheckTokenFail    = 20001
	ErrAuthCheckTokenTimeout = 20002
	ErrAuthToken             = 20003
	ErrAuth                  = 20004

	ErrGetRecordFailed = 30001

	ErrUserAlreadyExists = 40001
	ErrCreateUser        = 40002
	ErrGetAllGroup       = 40003
)
