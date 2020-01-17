package e

const (
	SUCCESS = 200
	ERROR   = 500

	ErrGetRecordFailed = 30001

	ErrUserAlreadyExists = 40001
	ErrCreateUser        = 40002
	ErrGetAllGroup       = 40003
)

// About Error,
const (
	ErrorV2                 = iota + 1
	ErrInvalidParams        = 400
	ErrUnKnownInternalError = 400
	// #1 Basic Auth -> to tell user that something went wrong when process basic auth
	ErrInvalidBasicAuthParam
	ErrBasicAuthFailed

	// #2 Jwt Auth
	ErrInvalidBearerAuthParams = 400
	ErrAuthCheckTokenFail      = 20001
	ErrAuthCheckTokenTimeout   = 20002
	ErrJwtAuth
	//
)
