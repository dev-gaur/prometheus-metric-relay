package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"hdd/scout/config"
	"hdd/scout/prometheusapi"
	"hdd/scout/util"

	"github.com/dev-gaur/workers"
	"github.com/robfig/cron"
)

type payload []byte

func pipeMetric(metricname string, client http.Client) error {
	fmt.Printf("\tGOROUTINE FOR METRIC : %s\n\n", metricname)

	promQueryEndpoint := "http://" + config.Config.Prometheus.Host + ":" + config.Config.Prometheus.Port + "/api/v1/query"
	supervisorEndpoint := "http://" + config.Config.Target.Host + ":" + config.Config.Target.Port + "/post"

	var err error
	var jsonResponse, postResponse payload
	var queryResult prometheusapi.VectorInstantQuery

	jsonResponse, err = util.SendGetRequest(&client, promQueryEndpoint, "query", metricname)

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonResponse, &queryResult)

	if err != nil {
		return err
	}

	postResponse, err = util.PostJSON(&client, supervisorEndpoint, queryResult.Data)

	if err != nil {
		return err
	}

	fmt.Printf("Response from PostJson :\n %#v \n\n", string(postResponse))

	return nil
}

func main() {
	fmt.Println("Launching....")
	errChan := make(chan error, 10)
	wrap, pooldone := make(chan struct{}), make(chan struct{})

	done := make(chan struct{})
	go func() {
		for {
			var c int
			fmt.Scanf("%d", c)

			if c == 0 {
				close(done)
				break
			}
		}
	}()

	cronString := "*/" + strconv.FormatInt(int64(config.Config.CronIntervalSeconds), 10) + " * * * * *"

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	var newtask = workers.NewTask
	n := int(config.Config.WorkerPoolSize)
	q := int(config.Config.TaskQueueSize)

	var pool = workers.GetPool(n, errChan, wrap, pooldone, q)

	c := cron.New()

	c.AddFunc(cronString, func() {

		var metricnames = config.Config.Metrics

		fmt.Println("\tCRON INTERVAL : ", time.Now())

		for i := range metricnames {
			task := newtask(func() error {
				return pipeMetric(metricnames[i], client)
			})

			pool.AssignTask(task)
		}

	})

	c.Start()
	<-done
	fmt.Println("About to quit...")
	c.Stop()
	fmt.Println("SHUTTING DOWN THE worker pool...")
	close(wrap)
	fmt.Println("WAITING for pooldone signal")
	<-pooldone
}
