package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LAzyDev-101/stupid-server/app"
)

func PostChallenge(rw http.ResponseWriter, r *http.Request) {
	var params app.RequestParams

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	if err = json.Unmarshal(body, &params); err != nil {
		log.Printf("error: %+v", err)
		return
	}

	resp, err := app.ProcessRequest(params)
	if err != nil {
		log.Printf("error: %+v", err)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(rw).Encode(resp); err != nil {
		log.Printf("error: %+v", err)
		return
	}

}