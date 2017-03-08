package demo

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/mb/parse"
)

// init registers the MetricSet with the central registry.
func init() {
	if err := mb.Registry.AddMetricSet("elasticon", "demo", New, hostParser); err != nil {
		panic(err)
	}
}

var (
	hostParser = parse.URLHostParserBuilder{
		DefaultScheme: "http",
	}.Build()
)

type MetricSet struct {
	mb.BaseMetricSet
	http *helper.HTTP
}

func New(base mb.BaseMetricSet) (mb.MetricSet, error) {

	return &MetricSet{
		BaseMetricSet: base,
		http:          helper.NewHTTP(base),
	}, nil
}

func (m *MetricSet) Fetch() (common.MapStr, error) {

	data, err := m.http.FetchJSON()
	if err != nil {
		return nil, err
	}

	event := common.MapStr{
		"counter": data["counter"],
	}

	return event, nil
}
