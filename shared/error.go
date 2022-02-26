package shared

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const (
	KeyErrDB             = "DB_ERROR"
	KeyErrInvalidRequest = "ErrInvalidRequest"
	KeyErrInternal       = "ErrInternal"
	KeyErrRemoteApi      = "RemoteApiError"
)

var ErrRecordNotFound = NewCustomError(errors.New("record not found"),
	"record not found", "ErrRecordNotFound")

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), KeyErrDB)
}

func ErrRemoteAPI(err error) *AppError {
	return NewAppError(err,
		http.StatusBadRequest,
		"some thing went wrong when call remote API", err.Error(),
		KeyErrRemoteApi)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), KeyErrInvalidRequest)
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(err, "internal error", err.Error(), KeyErrInternal)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("User already exists %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrUserAlreadyExists%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotGetRequestEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Request %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotRequest%s", entity),
	)
}

func ErrCannotPostRequestEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Request %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotRequest%s", entity),
	)
}

func ErrCannotPutRequestEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Request %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotRequest%s", entity),
	)
}

func ErrCannotPatchRequestEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Request %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotRequest%s", entity),
	)
}

func ErrCannotDeleteRequestEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Request %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotRequest%s", entity),
	)
}
