package prometheusapi

/*
	ScalarType describes a metric value during a particular timestamp
	format :
		"value": [
				unixtimestamp(int64),
				string or number
			]
*/
type ScalarType [2]interface{}

/*
	metricLabelMap describes a map of labelname and labelvalue
*/
type metricLabelMap map[string]string

type vectorInstantType struct {
	Metric metricLabelMap `json:"metric,string"`
	Value  ScalarType     `json:"value, string"`
}

type vectorInstantList []vectorInstantType

/*
	Range Vectors are the output of range queries at GET /api/v1/query_range endpoint.

	Instant Vectors Result type format :
		[
			{
				"metric" :	{ "<label_name>": "<label_value>", ... },
				"value" :	{ <unix_time>, "<sample_value>" }
			}
		]
*/
type VectorInstantQuery struct {
	Status string `json:"status, string"`
	Data   struct {
		ResultType string            `json:"resulttype, string"`
		Result     vectorInstantList `json:"result, string"`
	} `json:"data, string"`
}

type vectorRangeType struct {
	Metric metricLabelMap `json:"metric,string"`
	Values []ScalarType   `json:"values, string"`
}

type vectorRangeList []vectorRangeType

/*
	Range Vectors are the output of range queries at GET /api/v1/query_range endpoint.
	Range Vectors Result type format:

		[
			{
				"metric" :	{ "<label_name>": "<label_value>", ... },
				"values" :	[ { <unix_time>, "<sample_value>" }, ... ]
			}
		]
*/
type VectorRangeQuery struct {
	Status string `json:"status, string"`
	Data   struct {
		ResultType string          `json:"resulttype, string"`
		Result     vectorRangeList `json:"result, string"`
	} `json:"data, string"`
}

/*
	Result format when querying for possible values for a metric label
	endpoint : GET /api/v1/label/<label_name>/values
*/
type LabelValuesQuery struct {
	Status string   `json:"status, string"`
	Data   []string `json:"data, string"`
}

type seriesMetadata metricLabelMap

/*
	Result Format while querying for a time series metadata.
	endpoint: DELETE /api/v1/series
*/
type SeriesMetadataQuery struct {
	Status string           `json:"status, string"`
	Data   []seriesMetadata `json:"data, string"`
}

type targetFormat struct {
	ActiveTargets []struct {
		DiscoveredLabels metricLabelMap `json:"discoveredLabels, string"`
		Labels           metricLabelMap `json:"labels, string"`
		ScrapeURL        string         `json:"scrapeUrl, string"`
		LastError        string         `json:"lastError, string"`
		LastScrape       string         `json:"lastScrape, string"`
		Health           string         `json:"health, string"`
	} `json:"activeTargets, string"`
}

/*
	Result format for querying for a prometheus node's scrape targets.
	endpoint: GET /api/v1/targets
*/
type TargetQuery struct {
	Status string       `json:"status, string"`
	Data   targetFormat `json:"data, string"`
}

type alertmanagerFormat struct {
	ActiveAlertmanagers []struct {
		URL string `json:"url, string"`
	} `json:"activeAlertmanagers, string"`
}

/*
	Result Format for querying for Alert Managers.
	endpoint: GET /api/v1/alertmanagers
*/
type AlertmanagerQuery struct {
	Status string             `json:"status, string"`
	Data   alertmanagerFormat `json:"data, string"`
}

/*
	SeriesDeleteQuery is the result format for successful series deletion request.
	endpoint: DELETE /api/v1/series
*/
type SeriesDeleteQuery struct {
	Status string `json:"status, string"`
	Data   struct {
		NumberDeleted int `json:"numDeleted, string"`
	} `json:"data, string"`
}

/*
	Result format for an erroraneous result.
*/
type QueryError struct {
	Status    string `json:"status, string"`
	ErrorType string `json:"errorType, string"`
	Error     string `json:"error, string"`
}
