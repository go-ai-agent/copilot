package handler

import (
	"encoding/json"
	"github.com/go-sre/ai/percept"
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"net/http"
)

func AgentHandler_Ex3(w http.ResponseWriter, r *http.Request) {
	var status = runtime.NewStatusOK()

	switch r.Method {
	case http.MethodPost:
		// The exercise is to finish the code for the POST method
		var o percept.Observation

		buf, err := exchange.ReadAll(r.Body)
		//github:copilot
		if err != nil {
			status = runtime.NewStatus(http.StatusInternalServerError, location, err)
		} else {
			err = json.Unmarshal(buf, &o)
			if err != nil {
				status = runtime.NewStatus(http.StatusInternalServerError, location, err)
			} else {
				percept.Add(o)
			}
		}
	case http.MethodDelete:
	case http.MethodPut:

	default:
	}
	exchange.WriteResponse(w, nil, status)
}

//github:copilot
func unmarshal(r *http.Request, o *percept.Observation) error {
	buf, err := exchange.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, o)
}