package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/bqluan/httpx"

	"github.com/bqluan/bookbook/model"
)

func (h *apiHandler) GetBookDetail(w http.ResponseWriter, r *http.Request) {
	id, ok := ParamFromRequest(r, "id")
	if !ok {
		http.Error(w, "handler: missing book id", http.StatusBadRequest)
		return
	}

	var b model.Book
	err := h.db.QueryRow("select * from books where id = ?", id).Scan(
		&b.ID,
		&b.CreatedAt,
		&b.UpdatedAt,
		&b.Title,
		&b.Author,
		&b.Pub,
		&b.Desc,
		&b.Qty,
	)
	switch {
	case err == sql.ErrNoRows:
		http.Error(w, fmt.Sprintf("handler: book #%s not found", id), http.StatusNotFound)
	case err != nil:
		http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusInternalServerError)
	default:
		httpx.JSON(w, b, http.StatusOK)
	}
}
