// Package lrerr represents API error interface accessors for the SDK.
package lrerror

// An Error wraps lower level errors with code, message and an original error.
// The underlying concrete error type may also satisfy other interfaces which
// can be to used to obtain more specific information about the error.
//
// Calling Error() or String() will always include the full information about
// an error based on its underlying type.
//
// Example:
//
//     output, err := loginradius.GetAuthVerifyEmail()
//     if err != nil {
//         if lrError, ok := err.(lrError.Error); ok {
//             // Get error details
//             log.Println("Error:", lrError.Code(), lrError.Message())
//
//             // Prints out full error message, including original error if there was one.
//             log.Println("Error:", lrError.Error())
//
//             // Get original error
//             if origErr := lrError.OrigErr(); origErr != nil {
//                 // operate on original error.
//             }
//         } else {
//             ...
//         }
//     }
//
type Error interface {
	// Satisfy the generic error interface.
	error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErr() error
}

// BatchError is a batch of errors which also wraps lower level errors with
// code, message, and original errors. Calling Error() will include all errors
// that occurred in the batch.

type BatchError interface {
	// Satisfy the generic error interface.
	error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

// BatchedErrors is a batch of errors which also wraps lower level errors with
// code, message, and original errors. Calling Error() will include all errors
// that occurred in the batch.
//
// Replaces BatchError
type BatchedErrors interface {
	// Satisfy the base Error interface.
	Error

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

// New returns an Error object described by the code, message, and origErr.
//
// If origErr satisfies the Error interface it will not be wrapped within a new
// Error object and will instead be returned.
func New(code, message string, origErr error) Error {
	var errs []error
	if origErr != nil {
		errs = append(errs, origErr)
	}
	return newBaseError(code, message, errs)
}

// NewBatchError returns an BatchedErrors with a collection of errors as an
// array of errors.
func NewBatchError(code, message string, errs []error) BatchedErrors {
	return newBaseError(code, message, errs)
}
