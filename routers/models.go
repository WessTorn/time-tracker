package routers

type Response struct { // Для нормального описания сваггеру :D
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type Passport struct { // и это
	PassportNumber string `json:"passportNumber"`
}

type TaskID struct { // и это
	TaskID int `json:"task_id"`
}
