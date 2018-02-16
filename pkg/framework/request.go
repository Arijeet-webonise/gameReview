package framework

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// this struct basically adds a context to the http.Request so that
// authenticator or any other middleward could push out the data
// to main request handler
type Request struct {
	*http.Request
	context map[string]interface{}
}

func (r *Request) Push(key string, value interface{}) {
	if r.context == nil {
		r.context = map[string]interface{}{}
	}
	r.context[key] = value
}

func (r *Request) Value(key string) interface{} {
	return r.context[key]
}

func (r *Request) QueryParam(key string) string {
	return r.URL.Query().Get(key)
}

func (r *Request) ReadBody() (map[string]interface{}, error) {
	return ReadBody(r.Request)
}

func ReadBody(r *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	bodyMap := make(map[string]interface{})
	err := decoder.Decode(&bodyMap)
	if err != nil {
		return bodyMap, err
	}
	return bodyMap, nil
}

func (r *Request) Bind(v interface{}) error {
	return Bind(r.Request.Body, v)
}

func Bind(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	err := json.NewDecoder(body).Decode(v)
	return err
}

func GetPublicIPFromRequest(r *http.Request) (string, error) {
	pubIp := strings.Split(r.RemoteAddr, ":")[0]
	fIps := r.Header["X-Forwarded-For"]
	if len(fIps) < 1 {
		return "", errors.New("no ip in X-Forwarded-For header")
	}
	pubIp = strings.TrimSpace(strings.Split(fIps[0], ",")[0])
	return pubIp, nil
}
