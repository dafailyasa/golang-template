package tools

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func metaResponse(code int, status string) Meta {
	return Meta{
		Code:   code,
		Status: status,
	}
}

func ApiResponse(code int, status string, data interface{}) Response {
	return Response{
		Meta: metaResponse(code, status),
		Data: data,
	}
}

type ErrResponse struct {
	Meta   Meta        `json:"meta"`
	Errors interface{} `json:"errors"`
}

func ApiErrorResponse(code int, status string, errors interface{}) ErrResponse {
	return ErrResponse{
		Meta:   metaResponse(code, status),
		Errors: errors,
	}
}
