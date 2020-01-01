package response

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/conf"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/serializers"
)

// ErrorResponse return error message
func ErrorResponse(err error) serializers.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))

			return serializers.ParameterError(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializers.ParameterError("Json Type is not matched", err)
	}

	return serializers.ParameterError("Paramter error", err)
}
