// Code generated by GG version dev. DO NOT EDIT.

package server

import (
	"encoding/json"
	controller "github.com/555f/go-service-template/internal/usecase/controller"
	dto "github.com/555f/go-service-template/pkg/domain/dto"
	errors "github.com/555f/go-service-template/pkg/domain/errors"
	v5 "github.com/go-chi/chi/v5"
	"net/http"
)

func chiDefaultErrorEncoder(w http.ResponseWriter, err error) {
	var statusCode int = http.StatusInternalServerError
	if e, ok := err.(interface {
		StatusCode() int
	}); ok {
		statusCode = e.StatusCode()
	}
	if headerer, ok := err.(interface {
		Headers() http.Header
	}); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Set(k, v)
			}
		}
	}
	errorWrapper := errors.ErrorWrapper{}
	if e, ok := err.(interface {
		Data() interface{}
	}); ok {
		errorWrapper.Data = e.Data()
	}
	if e, ok := err.(interface {
		Error() string
	}); ok {
		errorWrapper.ErrorText = e.Error()
	}
	if e, ok := err.(interface {
		Code() string
	}); ok {
		errorWrapper.Code = e.Code()
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	bytes, err := json.Marshal(err)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(bytes)
}
func encodeBody(rw http.ResponseWriter, data any) {}

type UserControllerOption func(*UserControllerOptions)
type UserControllerOptions struct {
	errorEncoder     func(w http.ResponseWriter, err error)
	middleware       []func(http.Handler) http.Handler
	middlewareSearch []func(http.Handler) http.Handler
}

func UserControllerMiddleware(middleware ...func(http.Handler) http.Handler) UserControllerOption {
	return func(o *UserControllerOptions) {
		o.middleware = append(o.middleware, middleware...)
	}
}
func UserControllerWithErrorEncoder(errorEncoder func(w http.ResponseWriter, err error)) UserControllerOption {
	return func(o *UserControllerOptions) {
		o.errorEncoder = errorEncoder
	}
}
func UserControllerSearchMiddleware(middleware ...func(http.Handler) http.Handler) UserControllerOption {
	return func(o *UserControllerOptions) {
		o.middlewareSearch = append(o.middlewareSearch, middleware...)
	}
}
func SetupRoutesUserController(svc controller.UserController, r v5.Router, opts ...UserControllerOption) {
	o := &UserControllerOptions{errorEncoder: chiDefaultErrorEncoder}
	for _, opt := range opts {
		opt(o)
	}
	r.With(append(o.middleware, o.middlewareSearch...)...).Method("GET", "/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		q := r.URL.Query()
		nameQueryParam := q.Get("name")
		var paramName string
		if nameQueryParam != "" {
			paramName = nameQueryParam
		}
		users, err := svc.Search(r.Context(), paramName)
		if err != nil {
			o.errorEncoder(w, err)
			return
		}
		var resp struct {
			Users []*dto.User `json:"users"`
		}
		resp.Users = users
		var respData []byte
		respData, err = json.Marshal(resp)
		if err != nil {
			o.errorEncoder(w, err)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(200)
		w.Write(respData)
		return
	}))
}