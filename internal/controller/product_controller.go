package controller

import (
	"net/http"
	"strconv"

	"github.com/emejotaw/product-api/internal/service"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(db *gorm.DB) *ProductController {

	productService := service.NewProductService(db)
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) CreateProduct(ctx *fiber.Ctx) {

}

func (pc *ProductController) FindAll(ctx *fiber.Ctx) {

	page, errPage := strconv.Atoi(ctx.Query("page"))

	if errPage != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	size, errSize := strconv.Atoi(ctx.Query("size"))

	if errSize != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	sort := ctx.Query("sort")

	products, err := pc.productService.FindAll(page, size, sort)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(products)
}

func (pc *ProductController) FindByID(ctx *fiber.Ctx) {

	productId := ctx.Query("productId")

	product, err := pc.productService.FindByID(productId)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(product)
}

func (pc *ProductController) Update(ctx *fiber.Ctx) {

}

func (pc *ProductController) Delete(ctx *fiber.Ctx) {

}