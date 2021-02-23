package role

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const JWT_SECRET = "UCNLCU10DUC87%=58N"

type BaseContext struct {
	*gin.Context
}

func (c *BaseContext) PostFormGetInt(key string) (int, error) {
	return strconv.Atoi(c.PostForm(key))
}

func (c *BaseContext) PostFormGetFloat64(key string) (float64, error) {
	return strconv.ParseFloat(c.PostForm(key), 64)
}

func (c *BaseContext) QueryGetInt(key string) (int, error) {
	v := c.Query(key)
	if v == "" {
		return 0, nil
	} else {
		return strconv.Atoi(c.Query(key))
	}
}

type Resp struct {
	Code int         `json:"status_code"` //2 操作状态ID
	Text string      `json:"status_text"` //3 操作状态描述
	Data interface{} `json:"data"`
}

func (c *BaseContext) Json(code int, data interface{}, text string) {
	c.Context.JSON(200, Resp{
		Code: code,
		Data: data,
		Text: text,
	})
}

func verification(ctx *BaseContext) {

	if strings.Contains(ctx.Request.URL.String(), "/login") ||
		strings.Contains(ctx.Request.URL.String(), "/register") {
		ctx.Next()
		//不需要登录访问的接口
		return
	}

	token := ctx.GetHeader("token")
	token_str, err := ParseToken(token, JWT_SECRET)
	if err != nil {
		log.Println(err)
		ctx.Abort()
	}

	foo := struct {
		Uid string `json:"uid"`
		Exp int64  `json:"exp"`
	}{}

	err = json.Unmarshal([]byte(token_str), &foo)
	if err != nil {
		log.Println(err)
		ctx.Abort()
	}

	if foo.Exp > time.Now().Unix() {
		log.Println("token timeout")
		ctx.Abort()
	}

	ctx.Set("id", foo.Uid)

	ctx.Next()
}
