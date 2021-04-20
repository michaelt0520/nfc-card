package errors

import (
	"google.golang.org/grpc/status"
)

var (
	PasswordIncorrect = status.Error(1001, "Password incorrect")
	RecordNotFound    = status.Error(1002, "Record is not found")
	RecordInvalid     = status.Error(1003, "Record is invalid")
	ParameterInvalid  = status.Error(1004, "Parameter is invalid")

  // Auth
	InvalidToken       = status.Error(2001, "Invalid Token")
	DontHavePermission = status.Error(2002, "You don't have permission")
)
