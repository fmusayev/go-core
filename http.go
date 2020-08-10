package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ParseRequestBody(req *http.Request, model interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &model)
	if err != nil {
		return err
	}

	return nil
}

func ParseResponseBody(req *http.Response, model interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &model)
	if err != nil {
		return err
	}

	return nil
}

func OK(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

func MapHeaders(r *http.Request, keys ...string) map[string]string {
	hmap := make(map[string]string)

	for _, v := range keys {
		hmap[v] = r.Header.Get(v)
	}

	return hmap
}
