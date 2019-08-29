glog
====

在[golang/glog](https://github.com/golang/glog)的基础上做了一些修改。

## 修改的地方:
1. 增加每天切割日志文件的功能,程序运行时指定 --dailyRolling=true参数即可
2. 将日志等级由原来的INFO WARN ERROR FATAL改为DEBUG INFO ERROR FATAL
3. 增加日志输出等级设置,当日志信息等级低于输出等级时则不输出日志信息
4. 将默认的刷新缓冲区时间由20s改为5s

##使用示例 
```
func main() {
    //初始化命令行参数
    flag.Parse()
    //退出时调用，确保日志写入文件中
    defer glog.Flush()
    
    //一般在测试环境下设置输出等级为DEBUG，线上环境设置为INFO
    glog.SetLevelString("DEBUG") 
    
    glog.Info("hello, glog")
    glog.Warning("warning glog")
    glog.Error("error glog")
    
    glog.Infof("info %d", 1)
    glog.Warningf("warning %d", 2)
    glog.Errorf("error %d", 3)
 }
 
//假设编译后的可执行程序名为demo,运行时指定log_dir参数将日志文件保存到特定的目录
// ./demo --log_dir=./log --dailyRolling=true 
```
