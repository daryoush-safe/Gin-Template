package controller_v1_sample_general

type AddParams struct {
	Num1 string `uri:"num1" validate:"required,numeric,gt=10"`
	Num2 string `uri:"num2" validate:"required,numeric"`
}
