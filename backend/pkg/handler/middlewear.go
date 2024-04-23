package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		fmt.Println("empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		fmt.Println("invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func (h *Handler) adminIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		fmt.Println("empty auth header")
		return
	}
	if strings.Split(header," ")[1] == "administrator"{
		c.Set("admin","admin")
	}else{
		fmt.Println("wrong token")
		return 
	}


}

