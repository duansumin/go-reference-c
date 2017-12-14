## go reference c/c++ ##
### 目录结构 ###
```
|-project
|  |-lib
|  |  |-libsum.so
|  |-include
|  |  |-sum.h
|  |-src
|  |  |-main.go
|  |-pkg
|  |-bin
```

### 编译libsum.so ###

```
gcc -fPIC -shared sum.c -o libsum.so
```

### 简要步骤 ###

-  将c的方法提取到头文件.h中
-  编译c文件为动态链接库so文件  gcc -fPIC -shared sum.c -o libsum.so 
-  将头文件放入include目录 .so放入lib目录
-  go程序中指定 CFLAGS 和 LDFLAGS

```
#cgo  CFLAGS:  -I  ./include 
#cgo  LDFLAGS: -L  -lsum
```

-  运行go程序或者发布时候指定project中lib需要加入到LD_LIBRARY_PATH中
