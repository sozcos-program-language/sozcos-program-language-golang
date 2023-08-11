#### Go 语言的主要特征

```text
1.垃圾自动立即回收。
2.丰富的内置类型。
3.函数多返回值。
4.错误处理。
5.匿名函数和闭包。
6.类型和接口。
7.并发编程。
8.反射。
9.语言交互性。
```

#### Go 语言命名：

1.Go的函数、变量、常量、自定义类型、包(package)的命名方式遵循以下规则：

```text
1）首字符可以是任意的Unicode字符或者下划线
2）剩余字符可以是Unicode字符、下划线、数字
3）字符长度不限
```    

2.Go只有25个关键字

```text
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

3.Go还有37个保留字

```text
Constants:    true  false  iota  nil

Types:    int  int8  int16  int32  int64  
          uint  uint8  uint16  uint32  uint64  uintptr
          float32  float64  complex128  complex64
          bool  byte  rune  string  error

Functions:   make  len  cap  new  append  copy  close  delete
             complex  real  imag
             panic  recover
```

4.可见性：

```text
1）声明在函数内部，是函数的本地值，类似private
2）声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect
3）声明在函数外部且首字母大写是所有包可见的全局值,类似public
```

#### Go 语言的声明：

```text
var（声明变量）, const（声明常量）, type（声明类型） ,func（声明函数）。
```

#### Go 编译

```text
golang的编译使用命令 go build , go install;
注意点：
同一个目录中所有的go文件的package声明必须相同，所以main方法要单独放一个文件，否则在eclipse和liteide中都会报错；
```

#### Go mod模式创建项目一般目录结构

```text
project_name
 ├── bin 用于保存生成的可执行文件
 ├── cmd 保存可编译的主程序，既main方法
 └── example package 为example的go源文件
```