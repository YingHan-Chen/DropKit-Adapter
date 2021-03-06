package response

import (
	"github.com/DropKit/DropKit-Adapter/constants"
)

func AuthResponseOk() interface{} {
	response := constants.AuthResponse{0, "Ok"}

	return response
}

func AuthVerifyResponse(result bool) interface{} {
	response := constants.AuthVerifyResponse{0, "Ok", result}

	return response
}
