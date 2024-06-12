package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

//自定义一个go中间件 其它语言可能叫拦截器

func mayHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ces", "123")
		c.Next() // 放行
		//c.Abort() //阻止 中断

	}
}

type User struct {
	gorm.Model
	//'gorm:"type:varchar(20);not null"'
	Name string
	Age  string
}

func InitDB() *gorm.DB {
	//前提是你要先在本机用Navicat创建一个名为go_db的数据库
	host := "192.168.0.105"
	port := "3306"
	database := "go"
	username := "adminAll"
	password := "123456"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	//这里 gorm.Open()函数与之前版本的不一样，大家注意查看官方最新gorm版本的用法
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to Db connection, err: " + err.Error())
	}
	//这个是gorm自动创建数据表的函数。它会自动在数据库中创建一个名为users的数据表
	_ = db.AutoMigrate(&User{})
	return db
}

func main() {
	db := InitDB()
	// 切换到 release 模式
	gin.SetMode(gin.ReleaseMode)
	ginServer := gin.Default()
	// 设置受信任的代理（根据需要）
	ginServer.SetTrustedProxies([]string{"0.0.0.0"})

	//ginServer.Use(mayHandler()) // 全局注册中间件

	// 定义路由  进入这个方法前就会进入这个mayHandler中间件  这属于单个中间件
	ginServer.GET("/ping", mayHandler(), func(c *gin.Context) {

		ces := c.MustGet("ces").(string)
		c.JSON(200, gin.H{
			"message": "start",
			"ces":     ces,
		})
	})
	//接受参数
	//http://127.0.0.1:8088/user?name=he&age=25
	ginServer.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		//创建新用户
		newUser := User{
			Name: name,
			Age:  age,
		}
		db.Create(&newUser)
		c.JSON(200, gin.H{
			"message": "/user",
			"name":    name,
			"age":     age,
		})
	})
	//接受参数
	//http://127.0.0.1:8088/users/he/25
	ginServer.GET("/users/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		// 删除用户
		// Read
		var product User
		db.Where("Age = ?", age).First(&product)
		var users []User
		db.Find(&users)

		//result := db.Find(&users)
		c.JSON(200, gin.H{
			"message": "/users/:name/:age",
			"name":    name,
			"age":     age,
			"data":    users,
			"product": product,
			//"result":  resultID,
			//"err":     err,
		})
	})
	// 404 NoRoute
	ginServer.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"message": "404请求失败"})
	})

	// 路由组 /order/add
	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("add")
		orderGroup.GET("delete")
	}

	// 正确的监听端口配置
	if err := ginServer.Run(":8088"); err != nil {
		panic(err)
	}
}
