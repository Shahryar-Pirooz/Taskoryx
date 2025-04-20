package handlers

type Res struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
	Data   any    `json:"data"`
}
