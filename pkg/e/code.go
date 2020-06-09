package e

const (
	// #0 Generic Error
	SUCCESS                 = 200
	ERROR                   = 400
	ErrInvalidParams        = 403
	ErrUnKnownInternalError = 500

	// #1 Basic Auth -> to tell user that something went wrong when process basic auth
	ErrInvalidBasicAuthParam = 10001 + iota
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
