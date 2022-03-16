package main

import (
	"fmt"
	"gin-demo/mysqls"
	"gin-demo/sqlx"
	"github.com/gin-gonic/gin"
)

func func1(c *gin.Context) {
	c.String(200, "Hello World")
	fmt.Println("func1")
}

func func2(c *gin.Context) {
	fmt.Println("func2")
	c.Next()
	fmt.Println("func2 结束")
}
func func3(c *gin.Context) {
	fmt.Println("func3")
}

func func4(c *gin.Context) {
	fmt.Println("func4")
	c.Set("name", "张三")
}
func func5(c *gin.Context) {
	fmt.Println("func5")
	v, ok := c.Get("name")
	if ok {
		fmt.Println(v)
		str := v.(string)
		fmt.Println(str)
	}
}
func main() {
	fmt.Println("测试")
	//connectToDatabase()
	sqlxConnect()
	c := gin.Default()
	c.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	shopGroup := c.Group("/shop", func1, func2)
	shopGroup.Use(func3)
	{
		shopGroup.GET("/index", func4, func5)
	}
	err := c.Run()
	if err != nil {
		return
	}
	defer mysqls.DB.Close()
}
func connectToDatabase() {
	mysqls.ConnectMysql()
	mysqls.QueryRow()
	mysqls.QueryMultiRow()
	mysqls.InsertRow()
	mysqls.UpdateRow()
	mysqls.DeleteRow()
	mysqls.Pretreatment()
	mysqls.PretreatmentInsert()
}
func sqlxConnect() {
	sqlx.InitDB()
	sqlx.QueryRow()
	sqlx.QueryMultiRow()
	sqlx.InsertRow()
	sqlx.UpdateRow()
	sqlx.DeleteRow()
}
