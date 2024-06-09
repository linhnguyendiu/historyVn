package middleware

import (
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	//"strings"
)

func UserJwtAuthMiddleware(jwtAuth auth.JwtAuth, userService service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getHeader := ctx.GetHeader("Authorization")
		// if !strings.Contains(getHeader, "Bearer") {
		// 	panic(helper.NewUnauthorizedError("Who you are,"))
		// }
		//valueHeader := strings.Split(getHeader, " ")
		token := getHeader

		validateJwtToken, err := jwtAuth.ValidateJwtToken(token)
		if !validateJwtToken.Valid || err != nil {
			panic(helper.NewUnauthorizedError("Who you are, Hah?"))
		}

		// claims := validateJwtToken.Claims.(jwt.MapClaims)
		// userID := int(claims["user_id"].(float64))

		claims, ok := validateJwtToken.Claims.(jwt.MapClaims)
		if !ok {
			panic(helper.NewUnauthorizedError("Invalid token claims"))
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			panic(helper.NewUnauthorizedError("Invalid user ID in token"))
		}

		userId := int(userIDFloat)

		findByID := userService.FindById(userId)
		ctx.Set("current_user", findByID)
	}
}

func AuthorJwtAuthMiddleware(jwtAuth auth.JwtAuth, authorService service.AuthorService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getHeader := ctx.GetHeader("Authorization")
		// if !strings.Contains(getHeader, "Bearer") {
		// 	panic(helper.NewUnauthorizedError("You're not an author"))
		// }
		// valueHeader := strings.Split(getHeader, " ")
		token := getHeader

		validateJwtToken, err := jwtAuth.ValidateJwtToken(token)
		if !validateJwtToken.Valid || err != nil {
			panic(helper.NewUnauthorizedError("You're not an author"))
		}

		claims := validateJwtToken.Claims.(jwt.MapClaims)
		if claims["author_id"] == nil {
			panic(helper.NewUnauthorizedError("You're not an author"))
		}

		authorID := int(claims["author_id"].(float64))

		findByID := authorService.FindById(authorID)
		ctx.Set("current_author", findByID)
	}
}
