package models

// Response представляет стандартный ответ API
type Response struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse представляет ответ с ошибкой
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Operation failed"`
	Error   string `json:"error,omitempty" example:"Detailed error message"`
}

// HealthResponse представляет ответ health check
type HealthResponse struct {
	Status  string `json:"status" example:"ok"`
	Message string `json:"message" example:"Server is running"`
}
