package logs

import (
	"codelieche.com/monitor"
	"codelieche.com/settings"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"strconv"
	"time"
)

func (recordLog *RecordLogToStdOut) RecordLog(c chan monitor.Log) {
	for {
		executeLog := <-c
		log.Println("处理日志：", executeLog)
	}
}

func (recordLog *RecordLogToInfluxdb) RecordLog(c chan monitor.Log) {
	// 第1步：先获取数据库的连接、BatchPoint
	influxdbClient, bp, err := ConnectInfluxdb()
	if err != nil {
		log.Fatal(err)
	}
	defer (*influxdbClient).Close()

	// 写入数据库的点，每次写多少条数据
	var pointCountCache = settings.Config.DataBase.WriteLogLength

	var count = 0
	var timeWrite = time.Now()

	for {
		result, _ := <-c
		count += 1

		// Create a point and add to batch
		// 标签相当于索引，值是字符型
		tags := map[string]string{"monitor": strconv.Itoa(result.Id)}
		fields := map[string]interface{}{
			"success": result.Success,
			"elapsed": result.Elapesed,
		}

		pt, err := client.NewPoint("monitor", tags, fields, result.Time)

		if err != nil {
			log.Println(err)
			// 出现错误就重新连接数据库
			influxdbClient, bp, err = ConnectInfluxdb()
		} else {
			(*bp).AddPoint(pt)
		}

		if count%pointCountCache == 0 || time.Since(timeWrite).Seconds() > 60.0 {
			// Write the batch
			if err := (*influxdbClient).Write(*bp); err != nil {
				log.Println(err)
				// 出现错误就重新连接数据库
				influxdbClient, bp, err = ConnectInfluxdb()
			} else {
				// 更新写入时间，每分钟至少写一次，否则：缓存长度太长，数据少的时候，总是log不更新
				timeWrite = time.Now()
				count = 0
			}
		}
	}
}
