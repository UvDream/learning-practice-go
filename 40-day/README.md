

# 标识符与关键字

## 标识符与关键字

- 标识符
  在编程语言中标识符就是程序员定义的具有特殊意义的词，比如变量名、常量名、函数名等等。 Go 语言中标识符由字母数字和`_`(下划线）组成，并且只能以字母和`_`开头。 举几个例子：`abc, _, _123, a123`。
- 关键字

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

- 保留字

```go
Constants:    true  false  iota  nil

Types:  int  int8  int16  int32  int64
        uint  uint8  uint16  uint32  uint64  uintptr
        float32  float64  complex128  complex64
        bool  byte  rune  string  error

Functions:  make  len  cap  new  append  copy  close  delete
            complex  real  imag
            panic  recover
```

# 变量

## 变量的来历

程序运行过程中的数据都是保存在内存中，我们想要在代码中操作某个数据时就需要去内存上找到这个变量，但是如果我们直接在代码中通过内存地址去操作变量的话，代码的可读性会非常差而且还容易出错，所以我们就利用变量将这个数据的内存地址保存起来，以后直接通过这个变量就能找到内存上对应的数据了。

## 变量声明

Go 语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。 并且 Go 语言的变量声明后必须使用。

### 标准声明

```go
var 变量名 变量类型
```

### 批量声明

```go
var (
    name string  //*
    age int      //0
    isOk bool     //false
)
```

### 变量的初始化

```go
var 变量名 类型 = 表达式
```

### 类型推导

```go
var name = "Q1mi"
var age = 18
```

### 短变量声明

在函数内部，可以使用更简略的 := 方式声明并初始化变量。

```go
func main(){
    n:=11
    fmt.Println(n)
}
```

### 匿名变量

在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线\_表示，

```go
func foo() (int, string) {
    return 10, "Q1mi"
}
func main() {
    x, _ := foo()
    _, y := foo()
    fmt.Println("x=", x)
    fmt.Println("y=", y)
}
```

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 (在 Lua 等编程语言里，匿名变量也被叫做哑元变量。)

- 注意事项

1. 函数外的每个语句都必须以关键字开始（var、const、func 等）
2. :=不能使用在函数外。
3. \_多用于占位，表示忽略值

# 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把`var`换成了`const`，常量在定义的时候必须赋值。

## 声明

```go
const pi = 3.1415
const e = 2.7182
```

声明了`pi`和`e`这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。

## 批量申明

const 同时声明多个常量时，如果省略了值则表示和上面一行的值相同

```go
const (
    n1 = 100
    n2
    n3
)
```

```go
100,100,100
```

## iota

`iota`是 go 语言的常量计数器，只能在常量的表达式中使用。

`iota`在 const 关键字出现时将被重置为 0。const 中每新增一行常量声明将使`iota`计数一次(iota 可理解为 const 语句块中的行索引)。 使用 iota 能简化定义，在定义枚举时很有用。

```go
const (
        n1 = iota //0
        n2        //1
        n3        //2
        n4        //3
    )
```

### 使用`_`跳过某个值

```go
const (
        n1 = iota //0
        n2        //1
        _
        n4        //3
    )
```

### `iota`声明中间插队

```go
const (
        n1 = iota //0
        n2 = 100  //100
        n3 = iota //2
        n4        //3
    )
    const n5 = iota //0
```

### 定义数量级

（这里的`<<`表示左移操作，`1<<10`表示将 1 的二进制表示向左移 10 位，也就是由`1`变成了`10000000000`，也就是十进制的 1024。同理`2<<2`表示将 2 的二进制表示向左移 2 位，也就是由`10`变成了`1000`，也就是十进制的 8。）

```go
const (
        _  = iota
        KB = 1 << (10 * iota)
        MB = 1 << (10 * iota)
        GB = 1 << (10 * iota)
        TB = 1 << (10 * iota)
        PB = 1 << (10 * iota)
    )
```

### 多个`iota`定义在一行

```go
const (
        a, b = iota + 1, iota + 2 //1,2
        c, d                      //2,3
        e, f                      //3,4
    )
```

# 基本数据类型

## 整型

整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64

其中，`uint8`就是我们熟知的`byte`型，`int16`对应 C 语言中的`short`型，`int64`对应 C 语言中的`long`型

|  类型  | 描述                                                           |
| :----: | -------------------------------------------------------------- |
| uint8  | 无符号 8 位整型 (0 到 255)                                     |
| uint16 | 无符号 16 位整型 (0 到 65535)                                  |
| uint32 | 无符号 32 位整型 (0 到 4294967295)                             |
| uint64 | 无符号 64 位整型 (0 到 18446744073709551615)                   |
|  int8  | 有符号 8 位整型 (-128 到 127)                                  |
| int16  | 有符号 16 位整型 (-32768 到 32767)                             |
| int32  | 有符号 32 位整型 (-2147483648 到 2147483647)                   |
| int64  | 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

## 特殊整型

|  类型   | 描述                                                     |
| :-----: | :------------------------------------------------------- |
|  uint   | 32 位操作系统上就是`uint32`，64 位操作系统上就是`uint64` |
|   int   | 32 位操作系统上就是`int32`，64 位操作系统上就是`int64`   |
| uintptr | 无符号整型，用于存放一个指针                             |

**注意：** 在使用`int`和 `uint`类型时，不能假定它是 32 位或 64 位的整型，而是考虑`int`和`uint`可能在不同平台上的差异。

**注意事项** 获取对象的长度的内建`len()`函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用`int`来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用`int`和 `uint`。

## 八进制和十六进制

```go
package main

import "fmt"

func main(){
    // 十进制
    var a int = 10
    fmt.Printf("%d \n", a)  // 10
    fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制

    // 八进制  以0开头
    var b int = 077
    fmt.Printf("%o \n", b)  // 77

    // 十六进制  以0x开头
    var c int = 0xff
    fmt.Printf("%x \n", c)  // ff
    fmt.Printf("%X \n", c)  // FF
}
```

## 浮点数

Go 语言支持两种浮点型数：`float32`和`float64`。这两种浮点型数据格式遵循`IEEE 754`标准： `float32` 的浮点数的最大范围约为 `3.4e38`，可以使用常量定义：`math.MaxFloat32`。 `float64` 的浮点数的最大范围约为 `1.8e308`，可以使用一个常量定义：`math.MaxFloat64`

打印浮点数时，可以使用`fmt`包配合动词`%f`

```go
package main
import (
        "fmt"
        "math"
)
func main() {
        fmt.Printf("%f\n", math.Pi)
        fmt.Printf("%.2f\n", math.Pi)
}
```

## 复数

complex64 和 complex128

```go
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i
fmt.Println(c1)
fmt.Println(c2)
```

## 布尔值

Go 语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true（真）`和`false（假）`两个值。

**注意：**

1. 布尔类型变量的默认值为`false`。
2. Go 语言中不允许将整型强制转换为布尔型.
3. 布尔型无法参与数值运算，也无法与其他类型进行转换。

## 字符串

Go 语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用`UTF-8`编码。 字符串的值为`双引号(")`中的内容，可以在 Go 语言的源码中直接添加非 ASCII 码字符

### 字符串转义符

Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义符 | 含义                               |
| :----: | :--------------------------------- |
|  `\r`  | 回车符（返回行首）                 |
|  `\n`  | 换行符（直接跳到下一行的同列位置） |
|  `\t`  | 制表符                             |
|  `\'`  | 单引号                             |
|  `\"`  | 双引号                             |
|  `\\`  | 反斜杠                             |

## 多行字符串

```go
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```

## 字符串的常用操作

| 方法                                | 介绍           |
| :---------------------------------- | :------------- |
| len(str)                            | 求长度         |
| +或 fmt.Sprintf                     | 拼接字符串     |
| strings.Split                       | 分割           |
| strings.contains                    | 判断是否包含   |
| strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(),strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string, sep string) | join 操作      |

## byte 和 rune 类型

当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`。

```go
// 遍历字符串
func traversalString() {
    s := "hello沙河"
    for i := 0; i < len(s); i++ { //byte
        fmt.Printf("%v(%c) ", s[i], s[i])
    }
    fmt.Println()
    for _, r := range s { //rune
        fmt.Printf("%v(%c) ", r, r)
    }
    fmt.Println()
}
```

输出

```bash
104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³)
104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)
```

## 修改字符串

要修改字符串，需要先将其转换成`[]rune`或`[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。

```go
func changeString() {
    s1 := "big"
    // 强制类型转换
    byteS1 := []byte(s1)
    byteS1[0] = 'p'
    fmt.Println(string(byteS1))

    s2 := "白萝卜"
    runeS2 := []rune(s2)
    runeS2[0] = '红'
    fmt.Println(string(runeS2))
}
```

## 类型转换

Go 语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

### 基本语法

```bash
T(表达式)
```

# 运算符

Go 语言内置的运算符有：

1. 算术运算符
2. 关系运算符
3. 逻辑运算符
4. 位运算符
5. 赋值运算符

## 算数运算符

| 运算符 | 描述 |
| :----: | :--: |
|   +    | 相加 |
|   -    | 相减 |
|   \*   | 相乘 |
|   /    | 相除 |
|   %    | 求余 |

**注意：** `++`（自增）和`--`（自减）在 Go 语言中是单独的语句，并不是运算符。

## 关系运算符

| 运算符 |                              描述                              |
| :----: | :------------------------------------------------------------: |
|   ==   |     检查两个值是否相等，如果相等返回 True 否则返回 False。     |
|   !=   |   检查两个值是否不相等，如果不相等返回 True 否则返回 False。   |
|   >    |   检查左边值是否大于右边值，如果是返回 True 否则返回 False。   |
|   >=   | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 |
|   <    |   检查左边值是否小于右边值，如果是返回 True 否则返回 False。   |
|   <=   | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 |

## 逻辑运算符

| 运算符 |                                  描述                                   |
| :----: | :---------------------------------------------------------------------: |
|   &&   | 逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。  |
|  \|\|  | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。 |
|   !    |      逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。       |

## 位运算符

位运算符对整数在内存中的二进制位进行操作。

| 运算符 |                                            描述                                             |
| :----: | :-----------------------------------------------------------------------------------------: |
|   &    |                  参与运算的两数各对应的二进位相与。 （两位均为 1 才为 1）                   |
|   \|   |                参与运算的两数各对应的二进位相或。 （两位有一个为 1 就为 1）                 |
|   ^    | 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为 1。 （两位不一样则为 1） |
|   <<   | 左移 n 位就是乘以 2 的 n 次方。 “a<<b”是把 a 的各二进位全部左移 b 位，高位丢弃，低位补 0。  |
|   >>   |           右移 n 位就是除以 2 的 n 次方。 “a>>b”是把 a 的各二进位全部右移 b 位。            |

## 赋值运算符

| 运算符 |                      描述                      |
| :----: | :--------------------------------------------: |
|   =    | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
|   +=   |                  相加后再赋值                  |
|   -=   |                  相减后再赋值                  |
|  \*=   |                  相乘后再赋值                  |
|   /=   |                  相除后再赋值                  |
|   %=   |                  求余后再赋值                  |
|  <<=   |                   左移后赋值                   |
|  >>=   |                   右移后赋值                   |
|   &=   |                  按位与后赋值                  |
|  \|=   |                  按位或后赋值                  |
|   ^=   |                 按位异或后赋值                 |

# 流程控制

## if else(分支结构)

### if 条件判断基本写法

```go
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

### if 条件判断特殊写法(局部变量)

```go
if score := 65; score >= 90 {
        fmt.Println("A")
    } else if score > 75 {
        fmt.Println("B")
    } else {
        fmt.Println("C")
    }
```

## for 循环

```bash
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

- for 循环的初始语句可以被忽略，但是初始语句后的分号必须要写

```go
func forDemo2() {
    i := 0
    for ; i < 10; i++ {
        fmt.Println(i)
    }
}
```

- for 循环的初始语句和结束语句都可以省略

```go
func forDemo3() {
    i := 0
    for i < 10 {
        fmt.Println(i)
        i++
    }
}
```

## 无限循环

```go
for {
    循环体语句
}
```

for 循环可以通过`break`、`goto`、`return`、`panic`语句强制退出循环

## for range(键值循环)

Go 语言中可以使用`for range`遍历数组、切片、字符串、map 及通道（channel）。 通过`for range`遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值。
2. map 返回键和值。
3. 通道（channel）只返回通道内的值。

```go
s:="hello南京"
for i,v:=range s{
    fmt.Printf("%d %c\n",i,v)
}
```

## switch case

```go
func switchDemo1() {
    finger := 3
    switch finger {
    case 1:
        fmt.Println("大拇指")
    case 2:
        fmt.Println("食指")
    case 3:
        fmt.Println("中指")
    case 4:
        fmt.Println("无名指")
    case 5:
        fmt.Println("小拇指")
    default:
        fmt.Println("无效的输入！")
    }
}
```

## break(跳出循环)

```go
    for i:=0;i<10;i++{
        if i==5{
            break
        }
        fmt.Println(i)
    }
    fmt.Println("结束")
```

## continue(继续下次循环)

```go
for i:=0;i<10;i++{
        if i==5{
            continue
        }
        fmt.Println(i)
    }
    fmt.Println("结束")
```

## goto

`goto`语句通过标签进行代码间的无条件跳转。`goto`语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go 语言中使用`goto`语句能简化一些代码的实现过程。 例如双层嵌套的 for 循环要退出时：

```go
    var flag=false
    for i:=0;i<10;i++{
        for j:='A';j<'Z';j++{
            if j=='C'{
                flag=true
                break //跳出内层循环
            }
            fmt.Printf("%v-%c\n",i,j)
        }
        if flag{
            break
        }
    }
```

使用 goto

```go
for i:=0;i<10;i++{
        for j:='A';j<'Z';j++{
            if j=='C'{
                goto xx
            }
            fmt.Printf("%v-%c\n",i,j)
        }

    }
xx:
    fmt.Println("结束")
```

# 数组

## 数组定义

```bash
var 数组变量名 [元素数量]T
```

## 数组的初始化

### 方式 1

```go
a1=[3]bool{true,true,true}
fmt.Println(a1)
```

### 方式 2

```go
a100:=[10]int{0,1,2,3,4,5,6,7,8,9}
fmt.Println(a100)
```

### 方式 3(根据索引初始化)

```go
a3:=[5]int{0:1,2:6}
fmt.Println(a3)
```

## 数组遍历

### 1.根据索引

```go
cityS := [...]string{"北京", "上海", "深圳"}
for i:=0;i<len(cityS);i++{
        fmt.Println(cityS[i])
    }
```

### 2.for range

```go
for i,v:=range cityS{
        fmt.Println(i,v)
    }
```

## 多维数组

```go
a := [3][2]string{
        {"北京", "上海"},
        {"广州", "深圳"},
        {"成都", "重庆"},
    }
    fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
    fmt.Println(a[2][1]) //支持索引取值:重庆
```

### 多维数组遍历

```go
    a := [3][2]string{
        {"北京", "上海"},
        {"广州", "深圳"},
        {"成都", "重庆"},
    }
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s\t", v2)
        }
        fmt.Println()
    }
```

**注意：** 多维数组**只有第一层**可以使用`...`来让编译器推导数组长度。例如：

```go
//支持的写法
a := [...][2]string{
    {"北京", "上海"},
    {"广州", "深圳"},
    {"成都", "重庆"},
}
//不支持多维数组的内层使用...
b := [3][...]string{
    {"北京", "上海"},
    {"广州", "深圳"},
    {"成都", "重庆"},
}
```

### 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```go
// 数组值类型
	b1:=[3]int{1,2,3}   //b1 [1,2,3]
	b2:=b1				//b2 [1,2,3]
	b2[0]=100			//b2 [100,2,3]
	fmt.Println(b1,b2)
```

**注意：**

1. 数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
2. `[n]*T`表示指针数组，`*[n]T`表示数组指针 。

# 切片

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

## 切片的定义

```go
var name []T
```

```go
	// 切片定义
	var s1 []int
	var s2 []string
	fmt.Println(s1,s2)
	fmt.Println(s1==nil)  //true
	fmt.Println(s2==nil)  //true
	// 初始化
	s1=[]int{1,2,3}
	s2=[]string{"上海","北京","深圳"}
	fmt.Println(s1,s2)
	fmt.Println(s1==nil)  //false
	fmt.Println(s2==nil)  //false
```

## 切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的 len()函数求长度，使用内置的 cap()函数求切片的容量。

```go
	// 切片定义
	var s1 []int
	var s2 []string
	// 初始化
	s1=[]int{1,2,3}
	s2=[]string{"上海","北京","深圳"}
	fmt.Printf("len(s1):%d cap(s1):%d\n",len(s1),cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n",len(s2),cap(s2))
```

## 基于数组定义切片

由于切片的底层就是一个数组，所以我们可以基于数组定义切片。

```go
	// 由数组得到切片
	a1:=[ ... ]int{1,3,4,5,6,7,9}
	s3:=a1[0:4]   //基于数组切割左包含,右不含[)
	fmt.Println(s3)  //[1 3 4 5]
	s4:=a1[1:6]
	fmt.Println(s4) //[3 4 5 6 7]
	s5:=a1[:2]
	fmt.Println(s5) //[1 3]
	s6:=a1[:]
	fmt.Println(s6)  //[1 3 4 5 6 7 9]
```

## 切片再切片(指向底层数组)

```go
a1:=[ ... ]int{1,3,4,5,6,7,9}
s5:=a1[:4]
// 切片的容量是指底层数组的容量
fmt.Printf("len(s5):%d cap(s5):%d\n",len(s5),cap(s5))  //len(s5):4 cap(s5):7
// 底层数组从切片的第一个元素指向到最后一个元素的数量
fmt.Printf("len(s5):%d cap(s5):%d\n",len(s4),cap(s4))  //len(s5):5 cap(s5):6
//切片再切片
s8:=s4[3:]
fmt.Println("这是s4",s4)  	// [3 4 5 6 7]
fmt.Println("这是s8",s8)	//[6 7]
fmt.Printf("len(s8):%d cap(s8):%d\n",len(s8),cap(s8)) //len(s8):2 cap(s8):4
```

## 使用 make()函数构造切片

```bash
make([]T, size, cap)
```

- T:切片的元素类型
- size:切片中元素的数量
- cap:切片的容量

```go
// make()函数创造切片
s1 := make([]int, 5, 10)
fmt.Printf("s=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
```

## 切片的本质

切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

```go
a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}`，切片`s1 := a[:5]
```

![slice_01](https://www.liwenzhou.com/images/Go/slice/slice_01.png)

切片`s2 := a[3:6]`，相应示意图如下：

![slice_02](https://www.liwenzhou.com/images/Go/slice/slice_02.png)

## 切片不能直接比较

切片之间是不能比较的，我们不能使用`==`操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和`nil`比较。 一个`nil`值的切片并没有底层数组，一个`nil`值的切片的长度和容量都是 0。但是我们不能说一个长度和容量都是 0 的切片一定是`nil`

```go
var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
```

所以要判断一个切片是否是空的，要是用`len(s) == 0`来判断，不应该使用`s == nil`来判断。

## 切片的赋值拷贝

```go
	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3 // s3和s4都指向同一个底层数组
	s4[0] = 100
	fmt.Println(s3, s4) //[100 3 5] [100 3 5]
	s3[0] = 200
	fmt.Println(s3, s4) //[200 3 5] [200 3 5]
```

## 切片的遍历

```go
	// 索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	// for range
	for _, v := range s3 {
		fmt.Println(v)
	}
```

## append()方法为切片添加元素

Go 语言的内建函数`append()`可以为切片动态添加元素。 每个切片会指向一个底层数组，这个数组能容纳一定数量的元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在`append()`函数调用时

```go
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):3 cap(s1):3
	// s1[3]="南京" //错误写法,索引越界
	// append,调用必须使用原来的切片接受返回值,底层数组放不下的时候,原来的底层数组换一个位置3->6->12
	s1 = append(s1, "南京")
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):4 cap(s1):6
	fmt.Println(s1)                                         //[北京 上海 深圳 南京]
	// 多个元素
	s1 = append(s1, "合肥", "杭州", "无锡")
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):7 cap(s1):12
	// 追加切片
	a := []string{"成都", "重庆"}
	s1 = append(s1, a...) //a...拆开切片
	fmt.Println(s1)       //[北京 上海 深圳 南京 合肥 杭州 无锡 成都 重庆]
```

## 切片的扩容策略

- 首先判断，如果新申请容量（cap）大于 2 倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
- 否则判断，如果旧切片的长度小于 1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
- 否则判断，如果旧切片长度大于等于 1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的 1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
- 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

## 使用 copy()函数复制切片

```go
	// copy
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)            // 使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a1, a2, a3) //[1 3 5] [1 3 5] [1 3 5]
	a1[0] = 100
	fmt.Println(a1, a2, a3) //[100 3 5] [100 3 5] [1 3 5]
```

## 从切片中删除元素

```go
	// 删除元素
	// 删除索引为1的元素(其实就是就是拼接原理),切完容量不变
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)                                         //[100 5]
	fmt.Printf("len(a1):%d cap(a1):%d\n", len(a1), cap(a1)) //len(a1):2 cap(a1):3

	// 切片容量
	fmt.Println("切片容量")
	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Printf("%p\n", s1)
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):3 cap(s1):3
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	s1 = append(s1[:1], s1[2:]...) //修改了底层数组
	fmt.Printf("%p\n", s1)
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) // len(s1):2 cap(s1):3
	fmt.Println(x1)                                         // [1 5 5]
	// 通过切片可以修改底层数组
	s1[0]=100  //x1[100,5,5]
```

# 指针

区别于 C/C++中的指针，Go 语言中的指针不能进行偏移和运算，是安全指针。

要搞明白 Go 语言中的指针需要先知道 3 个概念：指针地址、指针类型和指针取值。

```go
	// 1.& 取地址
	n := 18
	fmt.Println(&n) //0xc000092008
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) //*int
	// 2.* 根据地址取值
	m := *p
	fmt.Println(m)        //18
	fmt.Printf("%T\n", p) //*int
```

取地址操作符`&`和取值操作符`*`是一对互补操作符，`&`取出地址，`*`根据地址取出地址指向的值。

变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：

- 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
- 指针变量的值是指针地址。
- 对指针变量进行取值（\*）操作，可以获得指针变量指向的原变量的值。

## new 和 make

```go
func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
```

### new

```go
func new(Type) *Type
```

- Type 表示类型，new 函数只接受一个参数，这个参数是一个类型
- \*Type 表示类型指针，new 函数返回一个指向该类型内存地址的指针。

new 函数不太常用，使用 new 函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。

```go
	var a = new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
```

### make

make 也是用于内存分配的，区别于 new，它只用于 slice、map 以及 chan 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make 函数的函数签名如下：

```go
func make(t Type, size ...IntegerType) Type
```

make 函数是无可替代的，我们在使用 slice、map 以及 channel 的时候，都需要使用 make 进行初始化，然后才可以对它们进行操作。这个我们在上一章中都有说明，关于 channel 我们会在后续的章节详细说明。

本节开始的示例中`var b map[string]int`只是声明变量 b 是一个 map 类型的变量，需要像下面的示例代码一样使用 make 函数进行初始化操作之后，才能对其进行键值对赋值：

```go
func main() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
```

### new 与 make 的区别

1. 二者都是用来做内存分配的。
2. make 只用于 slice、map 以及 channel 的初始化，返回的还是这三个引用类型本身；
3. 而 new 用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

# map

Go 语言中提供的映射关系容器为`map`，其内部使用`散列表（hash）`实现。

map 是一种无序的基于`key-value`的数据结构，Go 语言中的 map 是引用类型，必须初始化才能使用。

## map 定义

```go
map[KeyType]ValueType
```

- KeyType:表示键的类型。
- ValueType:表示键对应的值的类型。

```go
	// map,需要初始化后才能使用
	var m1 map[string]int
	// 初始化map
	m1 = make(map[string]int, 10)
	m1["北京"] = 1
	m1["上海"] = 2
	fmt.Println(m1)
	fmt.Println(m1["北京"])
```

## 查询是否存在某个键

```go
v, ok := m1["呵呵"]
	if !ok {
		fmt.Println("查无此地")
	} else {
		fmt.Println(v)
	}
```

## map 的遍历

```go
for k, v := range m1 {
		fmt.Println(k, v)
	}
```

## 使用 delete()函数删除键值对

使用`delete()`内建函数从 map 中删除一组键值对，`delete()`函数的格式如下：

```go
delete(map, key)
```

其中，

- map:表示要删除键值对的 map
- key:表示要删除的键值对的键

```go
	// 删除
	delete(m1, "上海")
	fmt.Println(m1)
	delete(m1, "呵呵") //删除不存在的key
```

## 按照指定顺序遍历 map

```go
func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	fmt.Println("map", scoreMap)
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```

## 元素为 map 类型的切片

```go
// 元素类型为map的切片
var s1 = make([]map[int]string, 10, 10)
// s1[0][100] = "1" //内部map没初始化
s1[0] = make(map[int]string, 1)
s1[0][10] = "北京"
fmt.Println(s1) //[map[1:北京] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
```

## 值为切片类型的 map

```go
// 值为切片类型的map
var m1 = make(map[string][]int, 10)
m1["北京"] = []int{10, 20, 30}
fmt.Println(m1) //map[北京:[10 20 30]]
```

# 函数

## 函数定义

Go 语言中定义函数使用`func`关键字，具体格式如下：

```go
func 函数名(参数)(返回值){
    函数体
}
```

其中：

- 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
- 参数：参数由参数变量和参数变量的类型组成，多个参数之间使用`,`分隔。
- 返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用`()`包裹，并用`,`分隔。
- 函数体：实现指定功能的代码块。

## 参数

### 类型简写

```go
func intSum(x, y int) int {
	return x + y
}
```

### 可变参数

```go
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
```

## 返回值

### 多返回值

```go
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
```

### 返回值命名

```go
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}
```

# 函数进阶

# defer

Go 语言中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行。

```go
	fmt.Println("start")
	defer fmt.Println("呵呵")
	fmt.Println("end")
```

```go
start
end
呵呵
```

由于`defer`语句延迟调用的特性，所以`defer`语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

## defer 执行时机

在 Go 语言的函数中`return`语句在底层并不是原子操作，它分为给返回值赋值和 RET 指令两步。而`defer`语句执行的时机就在返回值赋值操作后，RET 指令执行前。具体如下图所示：

![defer执行时机](https://www.liwenzhou.com/images/Go/func/defer.png)

## defer 经典

```go
// defer 延迟到函数即将返回的时候再执行
// 1.返回值赋值
// defer
// 2.真正的RET返回
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("呵呵")
	defer fmt.Println("嘿嘿")
	fmt.Println("end")
}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值=y=x=5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值=x=5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
```

```go
5
6
5
5
```

# 进阶函数

## 变量作用域

### 全局变量

全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。

```go
//定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}
func main() {
	testGlobalVar() //num=10
}
```

### 局部变量

局部变量又分为两种： 函数内定义的变量无法在该函数外使用，例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x：

```go
func testLocalVar() {
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

func main() {
	testLocalVar()
	fmt.Println(x) // 此时无法使用变量x
}
```

如果局部变量和全局变量重名，优先访问局部变量。if语句块局部作用域

## 函数类型和变量

### 定义函数类型

我们可以使用`type`关键字来定义一个函数类型，具体格式如下：

```go
type calculation func(int, int) int
```

eg:

```go
func add(x, y int) int {
	return x + y
}
//add能赋值给calculation类型的变量。
var c calculation
c = add
```

# 高阶函数

## 函数作为参数

## 函数作为返回值

```go

// 函数类型
func f1() {
	fmt.Println("函数")
}

func f2() int {
	return 10
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println("函数内部", ret)
}
func f4(x, y int) int {
	return x + y
}

// 函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}
func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(f2)
	f3(b)
	f7 := f5(f2)
	fmt.Printf("%T\n", f7)
}
```

## 匿名函数

函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数

```go
func(参数)(返回值){
    函数体
}
```

eg:

```go
// 函数内部不能声明带名字的函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(1, 2)
	// 立即执行函数
	func(x, y int) {
		fmt.Println(x, y)
	}(100, 200)
```

## 闭包

难点eg:

```go
package main

import "fmt"

// 闭包
func f1(f func()) {
	fmt.Println("这是f1")
	f()
}
func f2(x, y int) {
	fmt.Println("这是f2")
	fmt.Println(x + y)
}

// 定义一个函数对函数包装
func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		fmt.Println("1111")
		f(m, n)
	}
	return tmp
}

func main() {
	f1(f3(f2, 2, 1))
}
```



```go
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	ret := adder()
	ret2 := ret(200)
	fmt.Println(ret2)
}
```

# 内置函数介绍

|    内置函数    |                             介绍                             |
| :------------: | :----------------------------------------------------------: |
|     close      |                     主要用来关闭channel                      |
|      len       |      用来求长度，比如string、array、slice、map、channel      |
|      new       | 用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针 |
|      make      |   用来分配内存，主要用来分配引用类型，比如chan、map、slice   |
|     append     |                 用来追加元素到数组、slice中                  |
| panic和recover |                        用来做错误处理                        |

## panic/recover

Go语言中目前（Go1.12）是没有异常机制，但是使用`panic/recover`模式来处理错误。 `panic`可以在任何地方引发，但`recover`只有在`defer`调用的函数中有效。

```go

func funcA() {
	fmt.Println("A")
}
func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover B")
		}
	}()
	panic("出现严重错误!") //程序崩溃退出

}
func funcC() {
	fmt.Println("C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
```

**注意：**

1. `recover()`必须搭配`defer`使用。
2. `defer`一定要在可能引发`panic`的语句之前定义。



# 递归

```go
// 递归
func f(n uint64) uint64 {
	// 出口
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

func main() {
	ret := f(5)
	fmt.Println(ret)
}
```

# 结构体

Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

## 类型别名和自定义类型

### 自定义类型

在Go语言中有一些基本的数据类型，如`string`、`整型`、`浮点型`、`布尔`等数据类型， Go语言中可以使用`type`关键字来定义自定义类型。

```go
type myInt int     //自定义类型
```

### 类型别名

类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。

```go
type yourInt = int //类型别名
```

### 类型定义和类型别名的区别

```go
func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune
	c = '中'
	fmt.Printf("%T\n", c) //int32
}
```

## 结构体

Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称`struct`。 也就是我们可以通过`struct`来定义自己的类型了。

### 结构体定义/实例化

```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```

其中：

- 类型名：标识自定义结构体的名称，在同一个包内不能重复。
- 字段名：表示结构体字段名。结构体中的字段名必须唯一。
- 字段类型：表示结构体字段的具体类型。

```go
// 结构体
// 定义
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var w person
	// 实例化
	w.name = "wzj"
	w.age = 18
	w.gender = "男"
	w.hobby = []string{"篮球", "足球"}
	fmt.Println(w)
	fmt.Printf("%T\n", w)
	// 访问名字
	fmt.Println(w.name)
}
```

## 匿名结构体

在定义一些临时数据结构等场景下还可以使用匿名结构体。

```go
// 匿名结构体
var s struct {
	name string
	age  int
}
s.name = "张三"
s.age = 19
fmt.Println(s)
```

## 创建指针类型结构体

```go
// 结构体是值类型
type person struct {
	name string
	age  int
}

func f(x person) {
	x.age = 20
}
func f2(x *person) {
	//(*x).age = 30 //根据内存地址找到那个变量,修改的谁原来的变量
	x.age = 30 //语法糖,自动根据指针找对应变量
	fmt.Println("f2", *x)
}
func main() {
	var p person
	p.name = "呵呵"
	p.age = 10
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println("f2地址", p)
}
```

## 结构体初始化

没有初始化的结构体，其成员变量都是对应其类型的零值。

```go
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
}
```

### 使用键值对初始化

```go
p5 := person{
	name: "小王子",
	city: "北京",
	age:  18,
}
fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
```

也可以对结构体指针进行键值对初始化

```go
p6 := &person{
	name: "小王子",
	city: "北京",
	age:  18,
}
fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
```

### 使用值的列表初始化

```go
p8 := &person{
	"沙河娜扎",
	"北京",
	28,
}
fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
```

使用这种格式初始化时，需要注意：

1. 必须初始化结构体的所有字段。
2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
3. 该方式不能和键值初始化方式混用。

### 结构体内存布局

结构体占用一块连续的内存。

```go
// 结构体占用一块连续的内存。
type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a) //n.a 0xc00005a080
	fmt.Printf("n.b %p\n", &n.b) //n.b 0xc00005a081
	fmt.Printf("n.c %p\n", &n.c) //n.c 0xc00005a082
	fmt.Printf("n.d %p\n", &n.d) //n.d 0xc00005a083
}
```

## 构造函数

Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个`person`的构造函数。 因为`struct`是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。

```go
type person struct{
  name string
  age int64
}
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPerson("张三", 11)
	fmt.Println(p1)
}
```

当结构体比较大的时候尽量使用结构体指针,减少程序的内存开销,

* 构造函数new开头

## 方法和接收者

Go语言中的`方法（Method）`是一种作用于特定类型变量的函数。这种特定类型变量叫做`接收者（Receiver）`。接收者的概念就类似于其他语言中的`this`或者 `self`。

方法的定义格式如下：

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

其中，

- 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是`self`、`this`之类的命名。例如，`Person`类型的接收者变量应该命名为 `p`，`Connector`类型的接收者变量应该命名为`c`等。
- 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
- 方法名、参数列表、返回参数：具体格式与函数定义相同。

```go
// go语言如果标识符首字母大写就表示对外部可见

//Dog 这是一个狗的结构体//注释
// type Dog struct {
// 	name string
// }
type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法作用于特定的函数
func (d dog) wang() {
	fmt.Println("wangwangwang", d.name)
}

func main() {
	d1 := newDog("aaa")
	d1.wang()
}
```

## 指针类型的接受者/值接受者

```go
type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}
// 值类型(p person)为拷贝
func (p person) guonian() {
	p.age++
}

// (p *person)指针接受者
func (p *person) zGuoNian() {
	p.age++
}
func main() {
	p1 := newPerson("张三", 18)
	fmt.Println(p1)
	p1.guonian()
	fmt.Println(p1.age)
	p1.zGuoNian()
	fmt.Println(p1.age)
}
```

### 什么时候应该使用指针类型接收者

1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

## 给任意类型添加方法

```go
// 任意类型添加方法
// 不能给别的包的类型添加方法,只能给自己包的类型添加方法

//myInt 将int定义为自定义myInt类型
type myInt int

func (i myInt) hello() {
	fmt.Println("hello")

}
func main() {
	m := myInt(100)
	m.hello()
}
```

## 结构体匿名字段

结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

```go
//匿名字段
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"张三",
		9000,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
}
```

匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

## 嵌套结构体

```go
// 结构体嵌套
type address struct {
	province string
	city     string
}
type person struct {
	name string
	age  int
	addr address
}
type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "张三",
		age:  12,
		addr: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.addr.city)

}
```

### 嵌套匿名结构体

```go
type address struct {
	province string
	city     string
}
type person struct {
	name    string
	age     int
	address //匿名嵌套
}
type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "张三",
		age:  12,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.city)//匿名结构体,可以直接访问,先自己结构体找,然后匿名嵌套里面查找
}
```

### 嵌套结构体的字段名冲突

```go
type address struct {
	province string
	city     string
}
type workPlace struct {
	province string
	city     string
}
type person struct {
	name    string
	age     int
	address //匿名嵌套
	workPlace
}
...
	fmt.Println(p1.address.city)//只能这样取了
...
```

## 结构体的“继承”

```go
// 结构体实现继承
type animal struct {
	name string
}

// 给动物实现动的方法

func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// 狗
type dog struct {
	foot   int8
	animal //继承了animal
}

// 给狗实现叫
func (d dog) wang() {
	fmt.Printf("%s在叫\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "小三",
		},
		foot: 4,
	}
	fmt.Println(d1)
	d1.wang() //小三在叫
	d1.move() //小三会动
}
```

## 结构体字段的可见性

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

## 结构体与JSON序列化

```go
type person struct {
	Name string `json:"name" db:"name" ini:"name"` //实现json里面小写
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "张三",
		Age:  300,
	}
	p, err := json.Marshal(p1) //序列化:go语言结构体变量==>json格式的字符串
	//反序列化:json格式的字符串==>go语言中能识别的结构体变量
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(p))
	fmt.Println(string(p))
	// 反序列化
	str := `{"name":"张三","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}
```

# 接口interface

接口是一种类型

## 接口类型

在Go语言中接口（interface）是一种类型，一种抽象的类型。

`interface`是一组`method`的集合，是`duck-type programming`的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

```go
// 引出接口的实例
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}
type person struct {
}
type cat struct {
}
type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵猫")
}
func (d dog) speak() {
	fmt.Println("汪汪汪")
}
func (p person) speak() {
	fmt.Println("救命")
}

// 接受不定类型调用
func play(x speaker) {
	// 接受一个参数,传入进来就打印啥
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	play(c1)
	play(d1)
	play(p1)
}
```

## 接口的定义

Go语言提倡面向接口编程。

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

其中：

- 接口名：使用`type`将接口定义为自定义的类型名。Go语言的接口在命名时，<u>一般会在单词后面添加er</u>，如有写操作的接口叫`Writer`，有字符串功能的接口叫`Stringer`等。接口名最好要能突出该接口的类型含义。
- 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

```go
type writer interface{
    Write([]byte) error
}
```

当你看到这个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的Write方法来做一些事情

## 实现接口的条件

一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个**需要实现的方法列表**。

我们来定义一个`Sayer`接口：

```go
// Sayer 接口
type Sayer interface {
	say()
}
```

定义`dog`和`cat`两个结构体：

```go
type dog struct {}

type cat struct {}
```

因为`Sayer`接口里只有一个`say`方法，所以我们只需要给`dog`和`cat `分别实现`say`方法就可以实现`Sayer`接口了。

```go
// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}
```

接口的实现就是这么简单，只要实现了接口中的所有方法，就实现了这个接口。

eg:

```go
type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫")
}
func (c cat) eat(e string) {
	fmt.Printf("猫吃----%s\n", e)

}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡")
}
func (c chicken) eat(e string) {
	fmt.Printf("鸡吃----%s\n", e)
}

func main() {
	var a1 animal
	bc := cat{
		name: "小猫",
		feet: 4,
	}
	a1 = bc
	fmt.Println(a1)
	fmt.Printf("%T\n", a1) //main.cat
	bc.eat("猫粮")
	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)  //main.chicken
}
```

这个例子中,如果eat方法在实现的时候不传参,那么a1=bc的时候报错

![image-20191212112552486](/Users/wangzhongjie/Library/Application Support/typora-user-images/image-20191212112552486.png)

接口分为动态类型和动态值,这样就实现了接口变量能够存储不同的值

## 接口类型变量

那实现了接口有什么用呢？

接口类型变量能够存储所有实现了该接口的实例。 例如上面的示例中，`Sayer`类型的变量能够存储`dog`和`cat`类型的变量。

```go
func main() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}
```

## 值接收者和指针接收者实现接口的区别

```go
type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}
```

### 值接受者实现接口

```go
func (c cat) move() {
 	fmt.Println("猫")
 }
```

此时实现接口的cat类型

```go
	var a animal
	c1 := cat{"tom", 4}
	a = c1
	fmt.Println(a)        //{tom 4}
	fmt.Printf("%T\n", a) //main.cat
	a = &c1
	fmt.Println(a)        //&{tom 4}
	fmt.Printf("%T\n", a) //*main.cat
```

从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是cat结构体还是结构体指针cat类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，cat指针a内部会自动求值`*a`。

### 指针接受者实现接口

```go
func (c *cat) move() {
	fmt.Println("猫")
}
```

此时实现接口cat类型

```go
	var a animal
	c1 := cat{"tom", 4}
	//a = c1 //不可以接受c1
	a = &c1
	fmt.Println(a)        //&{tom 4}
	fmt.Printf("%T\n", a) //*main.cat
```

此时`move`接受的是`*cat`类型,所以`c1`不能传入`a`,此时`a`只能存储`*cat`类型的值

## 接口和类型的关系

### 一个类型实现多个接口

类似接口嵌套

```go
// 同一个结构体可以实现多个接口
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string)
}
type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("猫")
}
func (c *cat) eat(s string) {
	fmt.Printf("猫吃---:%s", s)
}

func main() {
	var a animal
	c1 := cat{name: "猫", feet: 1}
	a = &c1
	fmt.Println(a)
	a.eat("猫粮")

}
```

## 空接口

### 空接口的定义

空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

空接口类型的变量可以存储任意类型的变量。

```go
interface{} //空接口
```

### 空接口的应用

### 空接口作为函数的参数

```go
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

#### 空接口作为map的值

```go
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	fmt.Println(m1)
	m1["name"] = "张三"
	m1["age"] = 200
	m1["married"] = false
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)
}
```

## 类型断言

### 接口值

一个接口的值（简称接口值）是由`一个具体类型`和`具体类型的值`两部分组成的。这两部分分别称为接口的`动态类型`和`动态值`。

![接口值图解](https://www.liwenzhou.com/images/Go/interface/interface.png)

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

```go
x.(T)
```

其中：

- x：表示类型为`interface{}`的变量
- T：表示断言`x`可能是的类型。

```go
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int", t)

	case bool:
		fmt.Println("bool", t)
	case int64:
		fmt.Println("int64", t)
	}
}
func main() {
	assign(199)
}
```

# 包

## 包介绍

`包（package）`是多个Go源码的集合，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如`fmt`、`os`、`io`等

## 定义包

```go
package 包名
```

注意事项：

- 一个文件夹下面直接包含的文件只能归属一个`package`，同样一个`package`的文件不能在多个文件夹下。
- 包名可以不和文件夹的名字一样，包名不能包含 `-` 符号。
- 包名为`main`的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含`main`包的源代码则不会得到可执行文件。

## 可见性

如果想在一个包中引用另外一个包里的标识符（如变量、常量、类型、函数等）时，该标识符必须是对外可见的（public）。在Go语言中只需要将标识符的首字母<u>大写</u>就可以让标识符对外可见了。

```go
package pkg2

import "fmt"

// 包变量可见性

var a = 100 // 首字母小写，外部包不可见，只能在当前包内使用

// 首字母大写外部包可见，可在其他包中使用
const Mode = 1

type person struct { // 首字母小写，外部包不可见，只能在当前包内使用
	name string
}

// 首字母大写，外部包可见，可在其他包中使用
func Add(x, y int) int {
	return x + y
}

func age() { // 首字母小写，外部包不可见，只能在当前包内使用
	var Age = 18 // 函数局部变量，外部包不可见，只能在当前函数内使用
	fmt.Println(Age)
}
```

```go
type Student struct {
	Name  string //可在包外访问的方法
	class string //仅限包内访问的字段
}

type Payer interface {
	init() //仅限包内访问的方法
	Pay()  //可在包外访问的方法
}
```

## 包的导入

要在代码中引用其他包的内容，需要使用`import`关键字导入使用的包。具体语法如下:

```go
import "包的路径"
```

注意事项：

- import导入语句通常放在文件开头包声明语句的下面。
- 导入的包名需要使用双引号包裹起来。
- 包名是从`$GOPATH/src/`后开始计算的，使用`/`进行路径分隔。
- Go语言中禁止循环导入包。

## 自定义包名

在导入包名的时候，我们还可以为导入的包设置别名。通常用于导入的包名太长或者导入的包名冲突的情况。具体语法格式如下：

```go
import 别名 "包的路径"
```

单行导入方式定义别名：

```go
import "fmt"
import m "github.com/Q1mi/studygo/pkg_test"

func main() {
	fmt.Println(m.Add(100, 200))
	fmt.Println(m.Mode)
}
```

多行导入方式定义别名：

```go
import (
    "fmt"
    m "github.com/Q1mi/studygo/pkg_test"
 )

func main() {
	fmt.Println(m.Add(100, 200))
	fmt.Println(m.Mode)
}
```

## 匿名导入包

如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。具体的格式如下：

```go
import _ "包的路径"
```

匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中。

## init()初始化函数

### init()函数介绍

在Go语言程序执行时导入包语句会自动触发包内部`init()`函数的调用。需要注意的是： `init()`函数没有参数也没有返回值。 `init()`函数在程序运行时自动被调用执行，不能在代码中主动调用它。

包初始化执行的顺序如下图所示：![包中的init()执行时机](https://www.liwenzhou.com/images/Go/package/init01.png)

### init()函数执行顺序

Go语言包会从`main`包开始检查其导入的所有包，每个包中又可能导入了其他的包。Go编译器由此构建出一个树状的包引用关系，再根据引用顺序决定编译顺序，依次编译这些包的代码。

在运行时，被最后导入的包会最先初始化并调用其`init()`函数， 如下图示：![包之间的init()执行顺序](https://www.liwenzhou.com/images/Go/package/init02.png)

