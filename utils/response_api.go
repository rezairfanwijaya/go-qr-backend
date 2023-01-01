package utils

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ResponseAPI(message, status string, code int, data interface{}) responseAPI {
	meta := meta{
		Status:  status,
		Message: message,
		Code:    code,
	}

	return responseAPI{
		Meta: meta,
		Data: data,
	}
}
