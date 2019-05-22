package main

import (
	"fmt"
	"github.com/chenjiandongx/go-echarts/charts"
)

func main(){
	nameItems:=[]string{"衬衫","牛仔裤","运动裤","袜子","冲锋衣","羊毛衫"}
	bar:=charts.NewBar()
	fmt.Println(nameItems)
	bar.SetGlbalOptions(charts.TitleOpts{Title:"Bar-示例图"})
}
