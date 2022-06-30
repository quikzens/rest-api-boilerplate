package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quikzens/rest-api-boilerplate/helper"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func VerifyAuth(c *gin.Context) {
	authorizationHeader := c.GetHeader(authorizationHeaderKey)

	if len(authorizationHeader) == 0 {
		helper.SendUnauthorized(c, errors.New("authorization header is not provided"))
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		helper.SendUnauthorized(c, errors.New("invalid authorization header format"))
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		helper.SendUnauthorized(c, fmt.Errorf("unsupported authorization type %s", authorizationType))
		return
	}

	accessToken := fields[1]
	payload, err := helper.VerifyToken(accessToken)
	if err != nil {
		helper.SendUnauthorized(c, err)
		return
	}

	c.Set(authorizationPayloadKey, payload)
	c.Next()
}
