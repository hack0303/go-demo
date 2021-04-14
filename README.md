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
# point 2 Exported names
[Exported names](https://tour.golang.org/basics/3)
```
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again.
```