*必须加参数-fPIC -shared 编译成动态链接库 供go调用*
gcc -fPIC -shared -o libtest.so test.c test.h
*否则不加-fPIC -shared编译成可执行文件,cgo不能正确引用,否则这里会报错*
go build main.go
运行时候一定要export指定so文件的环境变量
export LD_LIBRARY_PATH="./"
./main

*这个是编译成可执行文件*
gcc -g -D__TEST__ -o test test.c test.h
-g gdb调试信息  -D__TEST__ define定义一个宏定义


$ go build main.go           
# command-line-arguments
.//libtest.so：在函数‘_fini’中：
(.fini+0x0): `_fini'被多次定义
/usr/lib/gcc/x86_64-linux-gnu/5/../../../x86_64-linux-gnu/crti.o:(.fini+0x0)：第一次在此定义
.//libtest.so：在函数‘data_start’中：
(.data+0x0): `__data_start'被多次定义
/usr/lib/gcc/x86_64-linux-gnu/5/../../../x86_64-linux-gnu/crt1.o:(.data+0x0)：第一次在此定义
.//libtest.so：在函数‘data_start’中：
(.data+0x8): `__dso_handle'被多次定义
/usr/lib/gcc/x86_64-linux-gnu/5/crtbegin.o:(.data+0x0)：第一次在此定义
.//libtest.so:(.rodata+0x0): `_IO_stdin_used'被多次定义
/usr/lib/gcc/x86_64-linux-gnu/5/../../../x86_64-linux-gnu/crt1.o:(.rodata.cst4+0x0)：第一次在此定义
.//libtest.so：在函数‘_start’中：
(.text+0x0): `_start'被多次定义
/usr/lib/gcc/x86_64-linux-gnu/5/../../../x86_64-linux-gnu/crt1.o:(.text+0x0)：第一次在此定义
.//libtest.so：在函数‘main’中：
./test.c:10: `main'被多次定义
/tmp/go-build862989160/command-line-arguments/_obj/_cgo_main.o:/tmp/go-build/command-line-arguments/_obj/_cgo_main.c:1：第一次在此定义
.//libtest.so：在函数‘_init’中：
(.init+0x0): `_init'被多次定义
/usr/lib/gcc/x86_64-linux-gnu/5/../../../x86_64-linux-gnu/crti.o:(.init+0x0)：第一次在此定义
/usr/lib/gcc/x86_64-linux-gnu/5/crtend.o:(.tm_clone_table+0x0): `__TMC_END__'被多次定义
.//libtest.so:(.data+0x10)：第一次在此定义
/usr/bin/ld ： 在.//libtest.so (.eh_frame)中发生错误; .eh_frame_hdr表格不会被创建。
collect2: error: ld returned 1 exit status
