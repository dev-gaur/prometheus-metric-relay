package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"hdd/scout/config"
	"hdd/scout/prometheusapi"
	"hdd/scout/util"

	"github.com/robfig/cron"
)

type payload []byte

func pipeMetric(metricname string, client http.Client, done chan struct{}) {
	fmt.Printf("\tGOROUTINE FOR METRIC : %s\n\n", metricname)

	promQueryEndpoint := "http://" + config.Config.Prometheus.Host + ":" + config.Config.Prometheus.Port + "/api/v1/query"
	supervisorEndpoint := "http://" + config.Config.Target.Host + ":" + config.Config.Target.Port + "/post"

	var err error
	var jsonResponse, postResponse payload
	var queryResult prometheusapi.VectorInstantQuery

	jsonResponse, err = util.SendGetRequest(&client, promQueryEndpoint, "query", metricname)

	if err != nil {
		log.Println(err)
		close(done)
		return
	}

	err = json.Unmarshal(jsonResponse, &queryResult)

	if err != nil {
		log.Println(err)
		close(done)
		return
	}

	postResponse, err = util.PostJSON(&client, supervisorEndpoint, queryResult.Data)

	if err != nil {
		log.Println("Error occured while Posting metrics to Supervisor.")
		log.Println(err)
		close(done)
		return
	}

	fmt.Printf("Response from PostJson :\n %#v \n\n", string(postResponse))
}

func main() {
	fmt.Println("Launching....")

	cronString := "*/" + strconv.FormatInt(int64(config.Config.CronIntervalSeconds), 10) + " * * * * *"

	//	promQueryEndpoint := "http://" + config.Config.Prometheus.Host + ":" + config.Config.Prometheus.Port + "/api/v1/query"
	//	supervisorEndpoint := "http://" + config.Config.Target.Host + ":" + config.Config.Target.Port + "/post"

	//	fmt.Println(cronString)
	//	fmt.Println(promQueryEndpoint)
	//	fmt.Println(supervisorEndpoint)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	c := cron.New()
	//var p *workers.Pool = workers.GetPool()
	//p.Run()

	done := make(chan struct{})

	c.AddFunc(cronString, func() {

		var metricnames = config.Config.Metrics

		fmt.Println("\tCRON INTERVAL : ", time.Now())

		for i := range metricnames {
			pipeMetric(metricnames[i], client, done)
		}

	})

	c.Start()
	<-done
	fmt.Println("About to quit...")
	c.Stop()
}
