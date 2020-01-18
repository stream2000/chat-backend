package e

const (
	// #0 Generic Error
	SUCCESS = iota + 1
	ERROR
	ErrInvalidParams
	ErrUnKnownInternalError

	// #1 Basic Auth -> to tell user that something went wrong when process basic auth
	ErrInvalidBasicAuthParam
	ErrBasicAuthFailed

	// #2 Auth
	ErrInvalidBearerAuthParams
	ErrAuthCheckTokenFail
	ErrAuthCheckTokenTimeout
	ErrJwtAuth
	ErrUnAuthorized
	// #3 User
	ErrUserNotFound
	// #4 Group
	ErrDeleteGroup
	ErrGroupNotFound
)
