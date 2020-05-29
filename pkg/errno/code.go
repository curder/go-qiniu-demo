package errno

//nolint: golint
var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &Errno{Code: 10003, Message: "Param error"}
	ErrNotFound         = &Errno{Code: 10004, Message: "The incorrect API route."}

	ErrValidation         = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase           = &Errno{Code: 20002, Message: "Database error."}
	ErrToken              = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrInvalidTransaction = &Errno{Code: 20004, Message: "invalid transaction."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrUserExists        = &Errno{Code: 20105, Message: "The user was already exists."}

	// account errors
	ErrAccountNotFound = &Errno{Code: 20200, Message: "The account resource was not found."}

	// bucket errors
	ErrBucketNotFound = &Errno{Code: 20300, Message: "The bucket resource was not found."}

	//
	ErrDomainNotFound = &Errno{Code: 20400, Message: "The bucket resource was not found"}
)
