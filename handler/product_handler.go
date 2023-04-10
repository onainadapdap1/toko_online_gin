package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"toko_online_gin/models"
	"toko_online_gin/service"
	"toko_online_gin/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandlerInterface interface {
	CreateProduct(c *gin.Context)
}

type productHandler struct {
	service service.ProductServiceInterface
}

func NewProductHandler(service service.ProductServiceInterface) ProductHandlerInterface {
	return &productHandler{service: service}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	// get data from request
	name := c.PostForm("name")
	description := c.PostForm("description")
	price := c.PostForm("price")	
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		response := utils.APIResponse("failed to convert price", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	quantity := c.PostForm("quantity")
	quantityInt, err := strconv.Atoi(quantity) 
	if err != nil {
		response := utils.APIResponse("failed to convert price", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	category_id := c.PostForm("category_id")
	categoryIdInt, err := strconv.Atoi(category_id) 
	if err != nil {
		response := utils.APIResponse("failed to convert category id", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		response := utils.APIResponse("failed to load image file", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(models.User)
	userId := currentUser.ID

	fileName := fmt.Sprintf("%d-%s", userId, file.Filename)

	dirPath := filepath.Join(".", "static", "images", "products")
	filePath := filepath.Join(dirPath, fileName)
	// create directory if doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			response := utils.APIResponse("failed to upload product image", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	// create file that will hold the image
	outputFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// open the temporary file that contains the uploaded image
	inputFile, err := file.Open()
	if err != nil {
		response := utils.APIResponse("failed to open product input image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	defer inputFile.Close()

	// copy the temporary image to the permanent location outputFile
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		response := utils.APIResponse("failed to copy product input image to permanent location", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	category, err := h.service.GetCategoryByID(uint(categoryIdInt))
	if err != nil {
		response := utils.APIResponse("failed to get category by id", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	createProductInput := models.CreateProductInput{
		Name: name,
		Description: description,
		Price: float64(priceInt),
		Quantity: quantityInt,
		CategoryID: category.ID,
		ImageURL: filePath,
	}

	newProduct, err := h.service.CreateProduct(createProductInput)
	if err != nil {
		log.Printf("failed to create product: %v", err)
		response := utils.APIResponse("failed to create produt", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("success to create product", http.StatusOK, "success", utils.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)

}