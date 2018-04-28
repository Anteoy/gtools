重要的说明
#cgo  CFLAGS:  -I  ./include 
#cgo  LDFLAGS:  -L ./lib  -lhello 

编译时候go文件需要以上两个路径来找到对应的include下的.h头文件和lib下是so文件
so文件一定是libhello.so 不然 -lhello 找不到对应的so文件

gcc -fPIC -shared -o libhello.so ./include/hello.c
go build main.go
运行时候一定要export指定so文件的环境变量
export LD_LIBRARY_PATH="./lib"
./main
