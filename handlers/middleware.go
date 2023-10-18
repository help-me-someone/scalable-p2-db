package handlers

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type ConnectionMiddleware struct {
	Connection *gorm.DB
}

func (d *ConnectionMiddleware) Attach(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := context.WithValue(r.Context(), "database", d.Connection)
		next(w, r.WithContext(ctx), p)
	}
}
