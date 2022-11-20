package views

import "time"

type CreateTodoPayload struct {
	ID          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Title TODO"`
	Description string    `json:"description" example:"Desc TODO"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetTodoPayload struct {
	ID          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Title TODO"`
	Description string    `json:"description" example:"Desc TODO"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTodoSuccessSwag struct {
	Status  int               `json:"status" example:"201"`
	Message string            `json:"message" example:"Create Todo Success"`
	Payload CreateTodoPayload `json:"payload"`
}

type CreateTodoFailureSwag struct {
	Status         int               `json:"status" example:"400"`
	Message        string            `json:"message" example:"Create Todo Fail"`
	Error          string            `json:"error" example:"Bad Request"`
	AdditionalInfo map[string]string `json:"additional_info" example:"error:Title harus diisi"`
}

type GetTodosSuccessSwag struct {
	Status  int              `json:"status" example:"200"`
	Message string           `json:"message" example:"Get All Todos Success"`
	Payload []GetTodoPayload `json:"payload"`
}

type GetTodoByIdSuccessSwag struct {
	Status  int            `json:"status" example:"200"`
	Message string         `json:"message" example:"Get Todo By ID Success"`
	Payload GetTodoPayload `json:"payload"`
}

type GetTodosFailureSwag struct {
	Status  int    `json:"status" example:"404"`
	Message string `json:"message" example:"Data Is Empty"`
	Error   string `json:"error" example:"Data Not Found"`
}

type UpdateTodoByIdSuccessSwag struct {
	Status  int            `json:"status" example:"200"`
	Message string         `json:"message" example:"Get Todo By ID Success"`
	Payload GetTodoPayload `json:"payload"`
}

type UpdateTodoFailureSwag struct {
	Status         int               `json:"status" example:"400"`
	Message        string            `json:"message" example:"Update Todo Fail"`
	Error          string            `json:"error" example:"Bad Request"`
	AdditionalInfo map[string]string `json:"additional_info" example:"error:Title harus diisi"`
}

type DeleteTodoByIdSuccessSwag struct {
	Status  int            `json:"status" example:"200"`
	Message string         `json:"message" example:"Get Todo By ID Success"`
	Payload GetTodoPayload `json:"payload"`
}

type DeleteTodoFailureSwag struct {
	Status  int    `json:"status" example:"404"`
	Message string `json:"message" example:"Delete Todo Fail"`
	Error   string `json:"error" example:"Data Not Found"`
}
