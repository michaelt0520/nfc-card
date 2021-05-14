package errors

import (
	"google.golang.org/grpc/status"
)

var (
	RecordNotFound   = status.Error(1001, "Record is not found")
	RecordInvalid    = status.Error(1002, "Record is invalid")
	ParameterInvalid = status.Error(1003, "Parameter is invalid")

	// Auth
	PasswordIncorrect             = status.Error(2001, "Password incorrect")
	ConfirmationPasswordIncorrect = status.Error(2002, "Confirmation password incorrect")
	InvalidToken                  = status.Error(2003, "Invalid Token")
	DontHavePermission            = status.Error(2004, "You don't have permission")

	// Deactivated
	DeactivatedCard = status.Error(3001, "Card is deactivated")

	DontExecuteYourSelf = status.Error(4001, "Dont execute yourself")
	UseableCard         = status.Error(4002, "Card has been used")
	CardNotCompatible   = status.Error(4003, "Card is not compatible")
)
