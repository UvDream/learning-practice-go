# 变量和常量

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
