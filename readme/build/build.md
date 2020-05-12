###build 一些错误
* windows : go build && cstest (不推荐，很多命令行操作没有很难受)  ps: cstest为编译后文件名
* linux/mac : go build && ./cstest (推荐)
* windows+linu子系统 : go build && ./cstest (不推荐很烦，有的用起来很难受)
* windows+git的bash : go build && ./cstest (推荐)
* build错误1：
    * 不要用github.com/gopm/modules  这有点儿恶心
     