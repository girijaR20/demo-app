package errorenum

import "demo-app/shared/model/apperror"

const (
	SomethingError       apperror.ErrorType = "ER0000 something error"
	UsernameMustNotEmpty apperror.ErrorType = "ER0001 username must not empty"
)
