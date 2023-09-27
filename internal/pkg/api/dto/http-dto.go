package dto

type HTTPBadRequestDTO struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPInternalServerErrorDTO struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"internal server error"`
}
