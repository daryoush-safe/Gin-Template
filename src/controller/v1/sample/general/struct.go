package controller_v1_sample_general

type AddParams struct {
	Num1 int `uri:"num1" validate:"required,number,gt=10"`
	Num2 int `uri:"num2" validate:"required,number"`
}
