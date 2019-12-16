# go语言标准库

## 1.fmt

### Print

```go
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

### Fprint

`Fprint`系列函数会将内容输出到一个`io.Writer`接口类型的变量`w`中，我们通常用这个函数往文件中写入内容。

### Sprint

`Sprint`系列函数会把传入的数据生成并返回一个字符串。

### Errorf

`Errorf`函数根据format参数生成格式化字符串并返回一个包含该字符串的错误

```go
err := fmt.Errorf("这是一个错误")
```

### 格式化占位符

#### 通用占位符

| 占位符 |                说明                |
| :----: | :--------------------------------: |
|   %v   |          值的默认格式表示          |
|  %+v   | 类似%v，但输出结构体时会添加字段名 |
|  %#v   |           值的Go语法表示           |
|   %T   |            打印值的类型            |
|   %%   |               百分号               |

#### 布尔型

| 占位符 |    说明     |
| :----: | :---------: |
|   %t   | true或false |

#### 整型

| 占位符 |                             说明                             |
| :----: | :----------------------------------------------------------: |
|   %b   |                         表示为二进制                         |
|   %c   |                    该值对应的unicode码值                     |
|   %d   |                         表示为十进制                         |
|   %o   |                         表示为八进制                         |
|   %x   |                   表示为十六进制，使用a-f                    |
|   %X   |                   表示为十六进制，使用A-F                    |
|   %U   |          表示为Unicode格式：U+1234，等价于”U+%04X”           |
|   %q   | 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 |

#### 浮点数与复数

| 占位符 |                          说明                          |
| :----: | :----------------------------------------------------: |
|   %b   |   无小数部分、二进制指数的科学计数法，如-123456p-78    |
|   %e   |              科学计数法，如-1234.456e+78               |
|   %E   |              科学计数法，如-1234.456E+78               |
|   %f   |           有小数部分但无指数部分，如123.456            |
|   %F   |                        等价于%f                        |
|   %g   | 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出） |
|   %G   | 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出） |

#### 字符串和[]byte

| 占位符 |                             说明                             |
| :----: | :----------------------------------------------------------: |
|   %s   |                   直接输出字符串或者[]byte                   |
|   %q   | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
|   %x   |           每个字节用两字符十六进制数表示（使用a-f            |
|   %X   |          每个字节用两字符十六进制数表示（使用A-F）           |

#### 指针

| 占位符 |              说明              |
| :----: | :----------------------------: |
|   %p   | 表示为十六进制，并加上前导的0x |

#### 宽度标识符

宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下：

| 占位符 |        说明        |
| :----: | :----------------: |
|   %f   | 默认宽度，默认精度 |
|  %9f   |  宽度9，默认精度   |
|  %.2f  |  默认宽度，精度2   |
| %9.2f  |    宽度9，精度2    |
|  %9.f  |    宽度9，精度0    |

#### 其他falg

| 占位符 |                             说明                             |
| :----: | :----------------------------------------------------------: |
|  ’+’   | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）； |
|  ’ ‘   | 对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格 |
|  ’-’   | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）； |
|  ’#’   | 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值； |
|  ‘0’   | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面； |

### 获取输入

Go语言`fmt`包下有`fmt.Scan`、`fmt.Scanf`、`fmt.Scanln`三个函数，可以在程序运行过程中从标准输入获取用户的输入

#### fmt.Scan

- Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
- 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

#### fmt.Scanf

- Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

#### fmt.Scanln

- Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

```go
func main() {
	var (
		name string
		age  int
		show bool
	)
	// 1
	// fmt.Scan(&name, &age, &show)
	// fmt.Println("扫描结果", name, age, show)
	// 2
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &show)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, show)
	// 3
	// fmt.Scanln(&name, &age, &show)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, show)

}
```

## 2.文件操作

### 打开和关闭文件

`os.Open()`函数能够打开一个文件，返回一个`*File`和一个`err`。对得到的文件实例调用`close()`方法能够关闭文件。

```go
// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
```

为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句。

### 读取文件

#### file.Read()

##### Read 方法定义

```go
func (f *File) Read(b []byte) (n int, err error)
```

eg

```go
func fileRead() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))

}
```

#### bufio读取文件

bufio是在file的基础上封装了一层API，支持更多的功能。

```go
func bufioRead() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("读取失败,%v\n", err)
			return
		}
		fmt.Print(line)
	}

}
```

#### ioutil读取整个文件

`io/ioutil`包的`ReadFile`方法能够读取完整的文件，只需要将文件名作为参数传入。

```go
func ioutilRead() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("读取失败,%v\n", err)
		return
	}
	fmt.Println(string(ret))
}
```

## 文件写入操作

`os.OpenFile()`函数能够以指定模式打开文件，从而实现文件写入相关功能。

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	...
}
```

其中：

`name`：要打开的文件名 `flag`：打开文件的模式。 模式有以下几种：

|     模式      |   含义   |
| :-----------: | :------: |
| `os.O_WRONLY` |   只写   |
| `os.O_CREATE` | 创建文件 |
| `os.O_RDONLY` |   只读   |
|  `os.O_RDWR`  |   读写   |
| `os.O_TRUNC`  |   清空   |
| `os.O_APPEND` |   追加   |

`perm`：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。

### Write和WriteString

```go
// 打开文件写内容
func main() {
	fileObj, err := os.OpenFile("./40-day/day05/12file_write/log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	// 关闭文件
	defer fileObj.Close()
	var str = "字符串\n"
	// 写入文件
	// Write
	fileObj.Write([]byte(str)) //写入字节切片数据
	// WriteString
	fileObj.WriteString("嘻嘻") //直接写入字符串数据
}
```

### bufio.NewWriter

```go
func write2() {
	fileObj, err := os.OpenFile("./40-day/day05/12file_write/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("又来写了") //写入缓存
	wr.Flush()             //将缓存中的写入文件
}
```

### ioutil.WriteFile

```go
func write3() {
	str := "再次来写"
	err := ioutil.WriteFile("./40-day/day05/12file_write/log.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
```

