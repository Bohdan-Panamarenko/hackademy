package kv

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"kv/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type parsedResponse struct {
	status int
	body   []byte
}

func createRequester(t *testing.T) func(req *http.Request, err error) parsedResponse {
	return func(req *http.Request, err error) parsedResponse {
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return parsedResponse{}
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return parsedResponse{}
		}

		resp, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return parsedResponse{}
		}

		return parsedResponse{res.StatusCode, resp}
	}
}

func prepareParams(t *testing.T, params interface{}) io.Reader {
	body, err := json.Marshal(params)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	return bytes.NewBuffer(body)
}

func assertStatus(t *testing.T, expected int, r parsedResponse) {
	if r.status != expected {
		t.Errorf("Unexpected response status. Expected: %d, actual: %d", expected, r.status)
	}
}

func assertBody(t *testing.T, expected string, r parsedResponse) {
	actual := string(r.body)
	if actual != expected {
		t.Errorf("Unexpected response body. Expected: %s, actual: %s", expected, actual)
	}
}

func assertResponse(t *testing.T, status int, body string, r parsedResponse) {
	assertStatus(t, status, r)
	assertBody(t, body, r)
}

func TestKeyValue(t *testing.T) {
	test := assert.New(t)
	doRequest := createRequester(t)

	t.Run("handle response test", func(t *testing.T) {

		s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			HandleResponse("hello world", rw)
		}))
		defer s.Close()

		resp := doRequest(http.NewRequest(http.MethodPost, s.URL, nil))
		assertResponse(t, http.StatusOK, "hello world", resp)
	})

	t.Run("handle error test", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			HandleError(errors.New("hello world"), rw)
		}))
		defer s.Close()

		resp := doRequest(http.NewRequest(http.MethodPost, s.URL, nil))
		assertResponse(t, http.StatusUnprocessableEntity, "hello world", resp)
	})

	t.Run("add test", func(t *testing.T) {
		stor := storage.NewStorage()

		serv := httptest.NewServer(AddHandler(stor))
		defer serv.Close()

		params := AddParams{
			Key:   "apple",
			Value: "juice",
		}

		resp := doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, params)))
		assertResponse(t, http.StatusOK, "succesfully added", resp)

		resp = doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, AddParams{
			Key:   "",
			Value: "juice",
		})))
		assertResponse(t, http.StatusUnprocessableEntity, "key can not be empty", resp)

		resp = doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, AddParams{
			Key:   "apple",
			Value: "",
		})))
		assertResponse(t, http.StatusUnprocessableEntity, "value can not be empty", resp)

		resp = doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, AddParams{
			Key:   "apple",
			Value: "juice",
		})))
		assertResponse(t, http.StatusUnprocessableEntity, "key 'apple' already exists", resp)
	})

	t.Run("update test", func(t *testing.T) {
		stor := storage.NewStorage()
		stor.Add("apple", "juice")

		serv := httptest.NewServer(UpdateHandler(stor))
		defer serv.Close()

		params := UpdateParams{
			Key:   "apple",
			Value: "jack",
		}

		resp := doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, params)))
		assertResponse(t, http.StatusOK, "succesfully updated", resp)

		value, err := stor.Get(params.Key)
		test.NoError(err)
		test.Equal(params.Value, value)

		resp = doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, UpdateParams{
			Key:   "",
			Value: "jack",
		})))
		assertResponse(t, http.StatusUnprocessableEntity, "key can not be empty", resp)

		resp = doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, UpdateParams{
			Key:   "apple",
			Value: "",
		})))
		assertResponse(t, http.StatusUnprocessableEntity, "value can not be empty", resp)
	})

	t.Run("get test", func(t *testing.T) {
		stor := storage.NewStorage()
		stor.Add("apple", "juice")

		serv := httptest.NewServer(GetHandler(stor))
		defer serv.Close()

		params := GetParams{
			Key: "apple",
		}

		resp := doRequest(http.NewRequest(http.MethodGet, serv.URL+"?key="+params.Key, nil))
		assertResponse(t, http.StatusOK, "juice", resp)

		resp = doRequest(http.NewRequest(http.MethodGet, serv.URL, nil))
		assertResponse(t, http.StatusUnprocessableEntity, "key can not be empty", resp)

		resp = doRequest(http.NewRequest(http.MethodGet, serv.URL+"?key=orange", nil))
		assertResponse(t, http.StatusUnprocessableEntity, "key 'orange' does not exist", resp)
	})

	t.Run("delete test", func(t *testing.T) {
		stor := storage.NewStorage()
		stor.Add("apple", "juice")

		serv := httptest.NewServer(DeleteHandler(stor))
		defer serv.Close()

		resp := doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, DeleteParams{
			Key: "",
		})))
		assertResponse(t, http.StatusUnprocessableEntity, "key can not be empty", resp)

		resp = doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, DeleteParams{
			Key: "orange",
		})))
		assertResponse(t, http.StatusUnprocessableEntity, "key 'orange' does not exist", resp)

		resp = doRequest(http.NewRequest(http.MethodPost, serv.URL, prepareParams(t, DeleteParams{
			Key: "apple",
		})))
		assertResponse(t, http.StatusOK, "succesfully deleted", resp)

		value, err := stor.Get("apple")
		test.Empty(value)
		test.EqualError(err, "key 'apple' does not exist")

	})
}
