package router

import (
	"net/http"
	"toko_online_gin/driver"
	"toko_online_gin/handler"
	"toko_online_gin/middlewares"
	"toko_online_gin/repository"
	"toko_online_gin/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Router() *gin.Engine {
	router := gin.Default()

	db := driver.ConnectDB()
	// models.InitTable(db)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	api := router.Group("api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	userRouter := api.Group("/users")
	{
		userRouter.Use(middlewares.Authentication())
		userRouter.POST("avatar", userAuthorization(userService), userHandler.UploadAvatar)
		userRouter.GET("fetch", userAuthorization(userService), userHandler.FetchUser)
	}
	categoryRouter := api.Group("/categories")
	{
		categoryRouter.GET("", categoryHandler.FindAllCategory)
		categoryRouter.GET("category/:slug", categoryHandler.FindBySlug)
		categoryRouter.Use(middlewares.Authentication())
		categoryRouter.POST("category", userAdminAuthorization(userService), categoryHandler.CreateCategory)
		categoryRouter.PUT("category/:slug", userAdminAuthorization(userService), categoryHandler.UpdateCategory)
	}
	productRouter := api.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("product", userAdminAuthorization(userService), productHandler.CreateProduct)
	}

	return router
}

func userAuthorization(userService service.UserServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// db := driver.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["user_id"].(float64))
		// user := models.User{}

		// select user_id from product where id = productID
		user, err := userService.GetUserByID(userId)
		// err := db.Where("id = ?", userId).Find(&user).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
		}

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are now allowed to access this data",
			})
			return
		}
		c.Set("currentuser", user)
		c.Next()
	}
}

func userAdminAuthorization(userService service.UserServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// db := driver.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["user_id"].(float64))
		// user := models.User{}

		// select user_id from product where id = productID
		user, err := userService.GetUserByID(userId)
		// err := db.Where("id = ?", userId).Find(&user).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
		}

		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are now allowed to access this data",
			})
			return
		}
		c.Set("currentUser", user)
		c.Next()
	}
}
