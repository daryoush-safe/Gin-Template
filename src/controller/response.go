package controller

type Response struct {
	Status  int16  `json:"statusCode"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ValidationResponse struct {
	Status   int16             `json:"statusCode"`
	Messages map[string]string `json:"messages"`
}
