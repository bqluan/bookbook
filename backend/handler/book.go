package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func (h *apiHandler) Borrow(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var b model.Borrow
	if err := dec.Decode(&b); err != nil {
		http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusBadRequest)
		return
	}
	b.CreatedAt = time.Now()
	b.UpdatedAt = b.CreatedAt
	stmt, err := h.db.Prepare("insert into borrows (created_at, updated_at, book_id, wechat_id) values (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusInternalServerError)
		return
	}
	res, err := stmt.Exec(b.CreatedAt, b.UpdatedAt, b.BookID, b.WechatID)
	if err != nil {
		http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusInternalServerError)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusInternalServerError)
		return
	}
	b.ID = uint(id)
	httpx.JSON(w, b, http.StatusCreated)
}
