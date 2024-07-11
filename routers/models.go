package routers

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type Passport struct {
	PassportNumber string `json:"passportNumber"`
}

type TaskID struct {
	TaskID int `json:"task_id"`
}
