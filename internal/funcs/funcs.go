package funcs

import (
	"encoding/json"
	"github.com/solverANDimprover/calc_go/internal/request"
	"io"
)

func SendErrJson(w io.Writer) {
	var BadRequest *request.RequestError = new(request.RequestError)
	BadRequest.Err = "Expression is not valid"
	answer, _ := json.Marshal(BadRequest.Err)
	w.Write(answer)
}
