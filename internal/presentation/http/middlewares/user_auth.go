package http_middlewares

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sousair/americastech-exchange/pkg/user_client"
)

type UserTokenPayload struct {
	UserId string `json:"id"`
	Email  string `json:"email"`
}

func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Authorization header missing",
			})
		}

		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Authorization header malformed",
			})
		}

		token := authHeaderParts[1]

		userApiBaseUrl := os.Getenv("USER_GRPC_API_URL")
		userClient := user_client.NewUserServiceGRPCClient(userApiBaseUrl)

		res, err := userClient.ValidateUserToken(c.Request().Context(), token)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "error validating user token",
			})
		}

		if !res {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid user token",
			})
		}

		tokenPayload, err := decodeJwtPayload(token)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
		}

		c.Set("user_id", tokenPayload.UserId)

		return next(c)
	}
}

func decodeJwtPayload(token string) (*UserTokenPayload, error) {
	encodedPayload := strings.Split(token, ".")[1]
	decoded, err := base64.RawStdEncoding.DecodeString(encodedPayload)

	if err != nil {
		return nil, err
	}

	var payload UserTokenPayload

	json.Unmarshal(decoded, &payload)
	return &payload, nil

}
