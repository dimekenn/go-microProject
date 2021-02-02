package handler

import (
	"fmt"
	"github.com/unistack-org/micro/v3/codec"
	"net/http"
)

type Handler struct {
	codec codec.Codec
}

func NewHandler(codec codec.Codec) *Handler {
	return &Handler{
		codec: codec,
	}
}

func (h *Handler) PanById(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Everything is OK")
}
