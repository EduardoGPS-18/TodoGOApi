package helpers

import (
	"fmt"
	"reflect"

	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/dtos"
)

func ConvertDomainErrorToDTO(err interface{}) *dtos.ErrorDto {
	fmt.Printf("ERROR: %v", reflect.TypeOf(err))

	switch err.(type) {
	case *errors.ValidationError, errors.ValidationError:
		return dtos.NewErrorDto(400, err.(*errors.ValidationError).Message())
	case *errors.SessionExpiredError, errors.SessionExpiredError:
		return dtos.NewErrorDto(401, err.(*errors.SessionExpiredError).Message())
	default:
		return dtos.NewErrorDto(500, "Internal server error!")
	}
}
