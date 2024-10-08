package ces

import (
	"ginDemo/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

//自定义一个go中间件 其它语言可能叫拦截器

func mayHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ces", "123")
		c.Next() // 放行之后会把后面的东西执行完 再回来执行剩余代码
		//c.Abort() //阻止 中断

	}
}

type User struct {
	gorm.Model
	//'gorm:"type:varchar(20);not null"'
	Name string
	Age  string
}

// DemoRouter 外部使用方法首字母必须大写公有方法，  内部使用方法首字母小写
type DemoRouter struct{}

// 在DemoRouter扩展方法InitDemoRouter 用s *DemoRouter接受 则不会复制DemoRouter结构体 而如果直接传入DemoRouter则会复制一份虽然都可以
// 尽管 InitDemoRouter 方法没有修改 DemoRouter 结构体的状态，
//使用指针接收者 (*DemoRouter) 依然是推荐的做法。这样可以确保代码的一致性、优化性能以及方便将来的扩展和维护。
// 一个指针变量指向了一个值的内存地址。 反过来 一个变量的的内存地址 是 指针类型
//interface{} 是一种万能类型，可以接受任何类型的值，包括指针类型、结构体类型、基本类型等

func (s *DemoRouter) InitDemoRouter(ginServer *gin.Engine) {
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
		global.DB.Create(&newUser)
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
		global.DB.Where("Age = ?", age).First(&product)
		var users []User
		global.DB.Find(&users)

		//result := global.DB.Find(&users)
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
}
