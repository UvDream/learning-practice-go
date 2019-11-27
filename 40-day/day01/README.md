# 标识符与关键字

## 标识符与关键字

- 标识符
  在编程语言中标识符就是程序员定义的具有特殊意义的词，比如变量名、常量名、函数名等等。 Go 语言中标识符由字母数字和`_`(下划线）组成，并且只能以字母和`_`开头。 举几个例子：`abc, _, _123, a123`。
- 关键字

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

- 保留字

```
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

