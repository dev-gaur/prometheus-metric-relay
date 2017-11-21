package config

import (
	"flag"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type url struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type configuration struct {
	WorkerPoolSize      int      `yaml:"worker_pool_size"`
	TaskQueueSize       int      `yaml:"task_queue_size"`
	CronIntervalSeconds int      `yaml:"cron_interval_seconds"`
	Metrics             []string `yaml:"metrics"`
	Target              url      `yaml:"target"`
	Prometheus          url      `yaml:"prometheus"`
}

func init() {
	configpath := *flag.String("config", "./config.yaml", "Config file path")
	flag.Parse()

	filename, _ := filepath.Abs(configpath)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)

	if err != nil {
		panic(err)
	}
}

/*
	Config ... Singleton instance of Configuration
*/
var Config configuration
