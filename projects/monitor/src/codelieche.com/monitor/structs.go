package monitor

import (
	"time"

	"sync"

	"codelieche.com/event"
)

/*
监控的结构体
1. 监控：Monitor
2. 步骤：Step

- 监控包含有多个步骤：
*/

type Monitor struct {
	Id          int      // 监控ID
	Project     string   // 项目名称
	Name        string   // 名称
	Category    string   `json:"category"` // 分类
	Scene       string   // 场景：开发、测试、Bata、正式
	Level       int      // 级别
	Retries     int      // 重试次数
	Notify      []string // 通知人是字符串的列表，比如：["user01", "user02"]
	Creator     string   // 添加者
	Status      string   // 状态
	IsActive    bool     `json:"is_active"` // 这条监控是否启用
	Description string   // 监控描述
	Steps       []Step   // 监控需要执行的步骤列表
}

// 监控执行的步骤
type Step struct {
	Id          int    // 步骤ID
	Monitor     string // 关联的监控名
	Name        string // 步骤名称：比如：访问首页、登录页
	Sleep       int    // 睡眠(毫秒)，这个步骤执行前需要sleep多久
	Order       int    // 排序：当一个监控有多个步骤的时候，通过order来排序, 越小，越执行在前面
	Url         string // 要访问的URL
	Method      string // 请求方法：GET、POST、PUT、DELETE等
	Data        string // 请求的数据
	Redirect    bool   //是否自动跳转
	Timeout     int    // 超时时间(毫秒)：发起http请求的时候设置个超时时间
	IsDeleted   bool   `json:"is_deleted"` // 是否删除了，有时候加入了几个步骤，如果标记了删除，这几个步骤就无需执行
	Description string // 步骤的描述

	// 响应相关字段
	CodeMin     int    `json:"code_min"`      // 响应码：最小值
	CodeMinExpr string `json:"code_min_expr"` // 响应码最小值是：>=、>、 =
	CodeMax     int    `json:"code_max"`      // 响应码：最大值
	CodeMaxExpr string `json:"code_max_expr"` // 最大值是：<=、=、<
	Title       string // 对相应的title进行校验，为空就不用校验
	Body        string // 响应的Body，如果为空就不校验

}

/*
监控打点日志
- 监控日志会写到数据库中：比如：influxdb
- 再利用grafana来做数据展示
- 后续再加入更多的字段
*/
type Log struct {
	Id       int       // 监控ID
	Time     time.Time // 日志时间
	Elapesed float64   // 监控操作消耗了多长时间
	Success  bool      // 是否成功
}

type Process struct {
	Source         Lister          // 监控列表的源【web监控列表/ 服务器监控列表 / 其它监控列表】
	TaskExecute    Executer        // 任务执行器【需要用到Execute的方法】请与source一一对应
	ExecuteInfoMap *ExecuteInfoMap // 任务执行信息
	TaskChan       chan Task       // 监控执行任务的channel
	LogChan        chan Log        // 监控执行日志的channel

}

// 监控的任务
type Task struct {
	Monitor      *Monitor  // 监控对象
	Status       string    // 任务的状态
	ExecutedTime time.Time // 执行时间
	ExpiredTime  time.Time // 过期时间
	//Execute      Executer  // 执行任务的接口
}

// 监控结果
type Result struct {
	Success  bool         // 成功
	Event    *event.Event // 事件，当失败的时候需要创建事件
	Executed bool         // 是否执行了，有些任务是过期了就不用执行了，这里会是false
	Elapsed  float64      // 执行时间（请求时间合计）参考了py中的requests的模块命名
}

// 监控执行信息
type ExecuteInfo struct {
	IsOk            bool      // 当前监控是否正常(指监控的web服务是否正常)
	ExecuteTime     time.Time // 上次执行的时间
	LastExecuteTime time.Time // 最后一次任务时间，生成任务以它为参考
	ExecuteCount    int       // 执行次数【当前这个监控执行测次数】
	ErrorCount      int       // 出错次数，当监控的服务正常了后，重置为0
}

// 监控执行信息映射：
type ExecuteInfoMap struct {
	Data *map[int]ExecuteInfo // 是个数组：[{id: info}]  id是监控的id，info是ExecuteInfo
	Lock *sync.RWMutex        // 读写锁，因为需要对监控的执行信息读取与写入，所以需要加入这个
}

// 结构体的方法: 获取监控的执行信息
func (info *ExecuteInfoMap) Get(k int) ExecuteInfo {
	// 先设置个读锁
	info.Lock.RLock()
	d := (*info.Data)[k]
	defer info.Lock.RUnlock() // 记得释放锁
	return d
}

// 设置监控的执行信息
func (info *ExecuteInfoMap) Set(k int, v ExecuteInfo) {
	info.Lock.Lock()
	(*info.Data)[k] = v
	defer info.Lock.Unlock()
}
