package common

import "fmt"

// Error indicating a problem connecting to the backend.
type ErrorDatastoreError struct {
	Err error
}

func (e ErrorDatastoreError) Error() string {
	return e.Err.Error()
}

// Error indicating a resource does not exist.  Used when attempting to delete or
// udpate a non-existent resource.
type ErrorResourceDoesNotExist struct {
	Err error
	Name string
}

func (e ErrorResourceDoesNotExist) Error() string {
	return fmt.Sprintf("resource does not exists with name '%s'", e.Name)
}

// Error indicating a resource already exists.  Used when attempting to create a
// resource that already exists.
type ErrorResourceAlreadyExists struct {
	Err  error
	Name string
}

func (e ErrorResourceAlreadyExists) Error() string {
	return fmt.Sprintf("resource already exists with name '%s'", e.Name)
}

// Error indicating a problem connecting to the backend.
type ErrorConnectionUnauthorized struct {
	Err error
}

func (e ErrorConnectionUnauthorized) Error() string {
	return "connection is unauthorized"
}

// Validation error containing the fields that are failed validation.
type ErrorValidation struct {
	ErrFields []ErroredField
}

type ErroredField struct {
	Name         string
	DisplayValue string
	Reason       string
}

func (e ErrorValidation) Error() string {
	if len(e.ErrFields) == 0 {
		return "unknown validation error"
	} else if len(e.ErrFields) == 1 {
		return fmt.Sprintf("error with field %s=%s (%s)",
			e.ErrFields[0].Name,
			e.ErrFields[0].DisplayValue,
			e.ErrFields[0].Reason)
	} else {
		s := "error with the following fields:\n"
		for _, f := range e.ErrFields {
			s = s + fmt.Sprintf("-  %s=%s (%s)\n",
				f.Name,
				f.DisplayValue,
				f.Reason)
		}
		return s
	}
}

type ErrorInsufficientIdentifiers struct {
}

func (e ErrorInsufficientIdentifiers) Error() string {
	return "insufficient identifiers"
}