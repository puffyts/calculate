package request

type RequestError struct {
	Err string `json:"error"`
}

type request struct {
	expression string `json:"expression"`
}
