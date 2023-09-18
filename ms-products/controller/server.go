package controller

import (
	"strings"

	"github.com/djudju12/ms-products/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	controller ProductController
	router     *gin.Engine
}

func NewServer(controller ProductController) *Server {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("price", model.ValidPrice)
		v.RegisterValidation("status", model.ValidStatus)
	}

	const productsPath = "/products"
	router.GET(joinPath(productsPath, "/:id"), controller.getProduct)
	router.GET(productsPath, controller.listProducts)
	router.POST(productsPath, controller.createProduct)
	router.DELETE(joinPath(productsPath, "/:id"), controller.inactiveProduct)
	router.PATCH(productsPath, controller.updateProductStatus)

	return &Server{
		controller: controller,
		router:     router,
	}
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func joinPath(path ...string) string {
	var sb strings.Builder
	for _, c := range path {
		sb.WriteString(c)
	}

	return sb.String()
}
