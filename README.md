# prometheus-metric-relay
Queries Prometheus server for metrics, transforms the payload, relays it to the specified target endpoint.

The list of metricnames, cron interval, source (prometheus) and target, worker_pool_size, task_queue_size can be specified in the config.yaml
file. 

The target server in mind is something like https://github.com/dev-gaur/mock_server
