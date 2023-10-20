package dto

type HTTPBadRequestDTO struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPPreconditionRequiredDTO struct {
	Code    int    `json:"code" example:"420"`
	Message string `json:"message" example:"precondition required"`
}

type HTTPInternalServerErrorDTO struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"internal server error"`
}
