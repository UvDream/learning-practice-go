package main

import "fmt"

//字符串是不可变值类型,内部用指针指向utf-8字节数组
//1.默认值是空字符串""
//2.索引访问s[i]
//3.不能用序号获取字节元素指针,&s[i]非法
//4.不可变类型,无法修改字节数组
//5.字节数组尾部不包含NULL
func main(){
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
	//连接跨行字符串时,"+"必须加载上一行末尾,否则会导致编译错误
	s1 := "Hello" +
		"World!"
	println(s1)

	//支持索引返回子串
	a := "Hello,World!"
	a1 := a[:5]
	println(a1)
	a2 := a[7:]
	println(a2)
	a3 := a[1:5]
	println(a3)

	//要修改字符串,可先将其转换成[]rune或者[]byte,完成后再转换成string,无论哪种都会重新分配内存,并且复制字节数组
	b := "abcd"
	bs := []byte(b)
	bs[1] = 'B'
	println(string(bs))
	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	println(string(us))

	//用for循环遍历字符串,也有byte和rune
	c := "abc汉字"
	for i := 0; i < len(c); i++ {
		fmt.Printf("%c,byte:", c[i], "\n")
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%c,rune:", r, "\n")
	}
}
