package logs

import (
	"codelieche.com/settings"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

// 连接Influxdb
func ConnectInfluxdb() (*client.Client, *client.BatchPoints, error) {
	// 连接到数据库，然后不断的从channel中取出记录数据，写入到数据库中
	var (
		database = settings.Config.DataBase.Database
		username = settings.Config.DataBase.User
		password = settings.Config.DataBase.Password
		host     = settings.Config.DataBase.Host
		port     = settings.Config.DataBase.Port
	)
	address := fmt.Sprintf("http://%s:%d", host, port)

	httpConfig := client.HTTPConfig{
		Addr:     address,
		Username: username,
		Password: password,
	}

	influxdbClient, err := client.NewHTTPClient(httpConfig)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s", // 精度 秒
	})

	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	return &influxdbClient, &bp, nil
}
