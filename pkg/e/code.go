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

	// #2 Jwt Auth
	ErrInvalidBearerAuthParams
	ErrAuthCheckTokenFail
	ErrAuthCheckTokenTimeout
	ErrJwtAuth
	//
)
