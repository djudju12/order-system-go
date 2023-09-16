package controller

import (
	"strings"

	"github.com/djudju12/order-system/ms-products/service"
	"github.com/djudju12/order-system/ms-products/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	service *service.ProductService
	router  *gin.Engine
}

func NewServer(service *service.ProductService) *Server {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("price", utils.ValidPrice)
	}

	const productsPath = "/products"
	router.GET(path(productsPath, "/:id"), service.GetProduct)
	router.POST(path(productsPath), service.CreateProduct)

	return &Server{
		service: service,
		router:  router,
	}
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func path(path ...string) string {
	var sb strings.Builder
	for _, c := range path {
		sb.WriteString(c)
	}

	return sb.String()
}
