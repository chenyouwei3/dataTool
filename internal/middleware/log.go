package middleware

import (
	"dataTool/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		path := c.Request.URL.Path      //请求路径
		query := c.Request.URL.RawQuery //query参数
		endTime := time.Now()
		costTime := endTime.Sub(startTime).Milliseconds()
		var username string
		ctxUser, exists := c.Get("user")
		if !exists {
			username = "未登录"
		}
		user, ok := ctxUser.(model.User)
		if !ok {
			username = "未登录"
		}
		username = user.Name //获取用户名
		operationLog := OperationLog{
			Model:     gorm.Model{},
			Username:  username,
			Ip:        c.ClientIP(),
			Method:    c.Request.Method,
			Query:     query,
			Path:      path,
			Desc:      "",
			Status:    c.Writer.Status(),
			StartTime: time.Time{},
			TimeCost:  costTime,
			UserAgent: c.Request.UserAgent(),
			Errors:    c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}
		fmt.Println(operationLog)
	}
}

type OperationLog struct {
	gorm.Model
	Username  string    `gorm:"type:varchar(20);comment:'用户登录名'" json:"username"`
	Ip        string    `gorm:"type:varchar(20);comment:'Ip地址'" json:"ip"`
	Method    string    `gorm:"type:varchar(20);comment:'请求方式'" json:"method"`
	Query     string    `gorm:"type:varchar(50)" json:"query"`
	Path      string    `gorm:"type:varchar(100);comment:'访问路径'" json:"path"`
	Desc      string    `gorm:"type:varchar(100);comment:'说明'" json:"desc"`
	Status    int       `gorm:"type:int(4);comment:'响应状态码'" json:"status"`
	StartTime time.Time `gorm:"type:datetime(3);comment:'发起时间'" json:"startTime"`
	TimeCost  int64     `gorm:"type:int(6);comment:'请求耗时(ms)'" json:"timeCost"`
	UserAgent string    `gorm:"type:varchar(20);comment:'浏览器标识'" json:"userAgent"`
	Errors    string    `gorm:"type:varchar(100)"json:"errors"`
}
