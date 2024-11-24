package controllers

import "github.com/gin-gonic/gin"

func AddUserController(ginEngine *gin.Engine) {

	ginEngine.POST("/signup", userSignUpHandler)

}

func userSignUpHandler(context *gin.Context) {
	context.Request.Body
}
