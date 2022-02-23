package main
import (
        "fmt"
        "math"
)
func main() {
        fmt.Printf("%f\n", math.Pi)
		fmt.Printf("%.2f\n", math.Pi)
		f1:=1.23456
		fmt.Printf("%T\n",f1)//默认go语言中小数都是float64类型
		f2:=float32(1.23456)
		fmt.Printf("%T\n",f2)
		//f1=f2  float32类型的书不能直接赋值给float63类型的变量
}