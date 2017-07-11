package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type apiHandler struct {
	db *sql.DB
}

func CreateAPIHandler(db *sql.DB) (http.Handler, error) {
	r := httprouter.New()
	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	h := apiHandler{db}
	r.GET("/book/:id", wrap(h.GetBookDetail))
	r.GET("/borrow", wrap(h.GetBorrowList))
	r.POST("/borrow", wrap(h.Borrow))
	r.GET("/code2openid", wrap(h.CodeToOpenID))

	return r, nil
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

const paramsKey key = 0

// ParamFromRequest returns the param value associated with this request for name, if any.
func ParamFromRequest(r *http.Request, name string) (string, bool) {
	params, ok := r.Context().Value(paramsKey).(httprouter.Params)
	if !ok {
		return "", false
	}
	val := params.ByName(name)
	if val == "" {
		return "", false
	}
	return val, true
}

func wrap(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := context.WithValue(r.Context(), paramsKey, p)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
