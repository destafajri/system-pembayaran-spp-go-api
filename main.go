package main

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/controller"
	_ "github.com/golang/mock/mockgen/model"
)

func main() {
	controller.Controller()
}
