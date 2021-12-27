# NOTICE
```
# 配置 GOPROXY 环境变量
export GOPROXY=https://goproxy.io,direct
# 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
export GOPRIVATE=git.mycompany.com,github.com/my/private
```
# REF
[GO 代理](https://goproxy.io/zh/docs/getting-started.html)
[.gitignore](https://www.jianshu.com/p/1c74f84e56b4)
[GO MODLUE](https://www.jianshu.com/p/760c97ff644c)
[VSCODE](https://www.cnblogs.com/majiang/p/14177790.html)

# ackownledge
## point 1 understand package
[see me](https://tour.golang.org/basics/1)
```
Packages
Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.

(To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed.)
```
```
go程序启动，从package main开始
```



# point 2 Exported names

[Exported names](https://tour.golang.org/basics/3)
```
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again.
```
```
go引入包 
import "xxx"
import "yyy"
或
import (
import "xxx"
import "yyy"
)

```



# point 3 defer

[defer](https://tour.golang.org/flowcontrol/12)
```
Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
```
# point 4 final tour
[final tour](https://tour.golang.org/methods/1)



## QUICK PATTERN

### functions

```
func methodName(param1 type,param2 type) type{
}

func methodName(param1 type,param2 type) (type,type){
return param2,param1
}

func methodName(param1 type,param2 type) (x type,y type){
x=0
y=1
return
}
```

###  Variables

The `var` statement declares a list of variables; as in function argument lists, the type is last.

A `var` statement can be at **package or function level**. We see both in this example.



```
package main

var x,y,z int
func afunc(){
var w,c int
}
```



### Variables with initializers

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```
package main

var x,y,z int =1,2,3
func afunc(){
var w,c=1,"xxx"
}
```



### Short variable declarations

**Inside a function**, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

```
package main

func main(){
x,y,z:=1,2,3
}
```

###  Basic types

Go's basic types are

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.

###  Zero values

Variables declared without an explicit initial value are given their *zero value*.

The zero value is:

- `0` for numeric types,
- `false` for the boolean type, and
- `""` (the empty string) for strings.

## Type conversions

**The expression `T(v)` converts the value `v` to the type `T`.**

Some numeric conversions:

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, put more simply:

```
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the `float64` or `uint` conversions in the example and see what happens.

###  Type inference

When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

```
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant:

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

Try changing the initial value of `v` in the example code and observe how its type is affected.

###  Constants

Constants are declared like variables, but with the `const` keyword.

Constants can be **character**, **string**, **boolean**, or **numeric** values.

Constants **cannot** be declared using the `:=` syntax.

```
package main

const world="世界"

func main(){
const hello="你好"
}
```

[**Numeric Constants**](https://go.dev/tour/basics/16) 待完善

# go web
[go web ref](https://www.zhihu.com/question/27370112)
[http go](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484112&idx=1&sn=79d0d3167d0d962fe41ec00cdafffbb0&chksm=fa80d347cdf75a51183182f14622af766538ca0c5335012e5e1cc50b100e78f2954fa3943770&scene=21#wechat_redirect)
[Golang 开源框架](https://www.cnblogs.com/linguoguo/p/12439211.html)