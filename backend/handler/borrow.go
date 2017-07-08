package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bqluan/httpx"

	"github.com/bqluan/bookbook/model"
)

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

func (h *apiHandler) GetBorrowList(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("book_id")
	wechatID := r.FormValue("wechat_id")

	rows, err := h.db.Query("select * from borrows where book_id = ? and wechat_id = ?", bookID, wechatID)
	if err != nil {
		http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var borrows []model.Borrow
	for rows.Next() {
		var b model.Borrow
		err := rows.Scan(
			&b.ID,
			&b.CreatedAt,
			&b.UpdatedAt,
			&b.BookID,
			&b.WechatID,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusInternalServerError)
			return
		}
		borrows = append(borrows, b)
	}

	httpx.JSON(w, borrows, http.StatusOK)
}
