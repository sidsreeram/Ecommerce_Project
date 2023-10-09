package http

import (
    "net/http"

    "github.com/ecommerce/internal/app/user/entities"
    "github.com/ecommerce/internal/app/user/usecases"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, userUsecase usercases.UserUsecase) {
    userRouter := r.Group("/users")
    {
        // Pass the userUsecase dependency to the SignUpPostHandler
        userHandler := UserHandler{userUsecase: userUsecase}
        userRouter.POST("/signup", userHandler.SignUpPostHandler)
    }
}

type UserHandler struct {
    userUsecase usercases.UserUsecase
}

func (u *UserHandler) SignUpPostHandler(c *gin.Context) {
    // Parse user input from the request (email, password, confirmPassword)
    var userInput entities.SignUpRequest
    if err := c.ShouldBindJSON(&userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Call the Register method from the user use case
    if err := u.userUsecase.Register(userInput.Email, userInput.Password, userInput.ConfirmPassword, &userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // User registered successfully
    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginHandler(userUsecase usercases.UserUsecase) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Implement user login logic using the userUsecase
    }
}
