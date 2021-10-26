package kv

import (
	"encoding/json"
	"errors"
	"kv/storage"
	"log"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(err.Error()))
}

func HandleResponse(response string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

// type Key struct {
// 	Key string `json:"key"`
// }

// func (k *Key) GetKey() string {
// 	return k.Key
// }

// type Value struct {
// 	Value []byte `json:"value"`
// }

// func (v *Value) GetValue() []byte {
// 	return v.Value
// }

type GetParams struct {
	Key string `json:"key"`
}

type AddParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UpdateParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DeleteParams struct {
	Key string `json:"key"`
}

func readJsonParams(r *http.Request, params interface{}) error {
	err := json.NewDecoder(r.Body).Decode(params)

	if err != nil {
		log.Println(err)
		return errors.New("could not read params")
	}

	return nil
}

func GetHandler(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		key := r.URL.Query().Get("key")
		if key == "" {
			HandleError(errors.New("key can not be empty"), w)
			return
		}

		value, err := s.Get(key)
		if err != nil {
			HandleError(err, w)
			return
		}

		HandleResponse(value, w)
	}
}

func AddHandler(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := AddParams{}

		err := readJsonParams(r, &params)
		if err != nil {
			HandleError(err, w)
			return
		}

		err = s.Add(params.Key, params.Value)
		if err != nil {
			HandleError(err, w)
			return
		}

		HandleResponse("succesfully added", w)
	}
}

func UpdateHandler(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := UpdateParams{}

		err := readJsonParams(r, &params)
		if err != nil {
			HandleError(err, w)
			return
		}

		err = s.Update(params.Key, params.Value)
		if err != nil {
			HandleError(err, w)
			return
		}

		HandleResponse("succesfully updated", w)
	}
}

func DeleteHandler(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := DeleteParams{}

		err := readJsonParams(r, &params)
		if err != nil {
			HandleError(err, w)
			return
		}

		_, err = s.Delete(params.Key)
		if err != nil {
			HandleError(err, w)
			return
		}

		HandleResponse("succesfully deleted", w)
	}
}
