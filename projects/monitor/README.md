## Monitor项目说明
> 监控项目：主要是web监控。

### 第三方依赖

**1. grequests:**
- 安装`grequests`: `go get -u github.com/levigross/grequests`

**2. influxdb client**
- 位置：`github.com/influxdata/influxdb/client/v2`


### 运行
- `cd src/codelieche.com/entry`
- 执行：`go run monitor.go --username admin --password AdminPwd --db devops --dbhost 127.0.0.1 --dbport 8086`

---

## 结构
通用说明：  
- 定义的结构体，一般写在`structs.go`文件中
- 定义的接口，一般写在`interfaces.go`文件中

### struct
> 核心的包是：`monitor`, 基本的结构体都在:`monitor/structs.go`文件中。

#### 基本的Struct
- `Monitor`：监控对象
- `Step`: 监控的步骤，每个Monitor可以有多个步骤
- `Task`: 监控的任务
- `Result`：执行完监控任务，要返回条结果，是否成功，失败的话有条`*Event`
- `Log`：每次执行完监控任务，会有条日志
- `Event`: 当出现了异常时候，需要创建事件，或者由异常变成了正常也需要创建事件
- `Process`：监控处理

#### 事件相关
- `FixResponse`: 自动修复的响应，当监控由异常变成正常，需要操作自动修复

#### Process
- `Source`: 需要实现Lister的接口，用来获取监控的列表
- `TaskExecute`: 需要实现Executer接口，用来执行监控任务
- `TaskChan`: 监控执行任务的channel
- `LogChan`: 监控执行日志的channel
- `EventHandle`: 处理异常事件的操作
- `LogHandle`: 处理日志相关的操作
- 实例化Process的时候需要传递上面的值

### package
- `monitor`: 监控的核心包：主要定义了`struct`和`interface`
- `source`: 监控数据源的包，主要是实现`monitor.Lister`接口的`struct`
- `execute`: 任务执行相关的包，主要是实现`monitor.Executer`接口的`struct`
- `event`: 任务异常事件处理相关的包，主要实现`monitor.Handler`接口的`struct`
- `logs`: 任务执行日志处理的包，主要实现`monitor.LogHandler`接口
- `settings`: 配置相关的包
- `server`: 监控执行服务端的包，主要是实例化`monitor.Process`然后执行`Run`

### Interface
- `Lister`: 获取监控列表的接口
    1. 方法：`List()([]Monitor, error)`
- `Executer`: 执行监控任务的接口
    1. 方法：`Execute(task *Task)(*Result, error)`
- `Handler`：异常事件处理的接口
    1. 方法：`Report(event *Event)(bool, string)`
    2. 方法：`AutoFix(*Monitor)(FixResponse, error)`
- `LogHandler`: 处理日志相关的接口
    1. 方法：`RecordLog(c chan Log)`