package slis

import (
	"cmp"
	"context"
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"time"

	"github.com/jenciso/slis-pyrra/pkg/apps"
	"github.com/jenciso/slis-pyrra/pkg/slogen"
	"github.com/mitchellh/mapstructure"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func (s *SliSchemaYaml) setValuesEndpointsRatio(conf SliConf, opts slogen.CmdOpts) {

	queryGetEndpoints := fmt.Sprintf(
		`group by (endpoint)(label_replace(%s_count{app="%s"}, "endpoint", "$1", "%s", "(.*)"))`,
		conf.Conf.MetricName, conf.App, conf.Conf.EndpointLabel)

	r := queryPrometheus(queryGetEndpoints, opts.PrometheusURL)

	switch res := r.(type) {
	case model.Vector:
		for _, v := range res {
			var e Endpoints
			eName := fmt.Sprint(v.Metric["endpoint"])
			if eName != "" {
				e.Name = eName
				e.Target = conf.Conf.TargetDefault
				s.Spec.Slis.Ratio.Endpoints = append(s.Spec.Slis.Ratio.Endpoints, e)
			}
		}
	}
	slices.SortFunc(s.Spec.Slis.Ratio.Endpoints, CmpName)
}

func (s *SliSchemaYaml) setValuesEndpointsLatency(conf SliConf, opts slogen.CmdOpts) {

	queryGetEndpointsWithAvgRt := fmt.Sprintf(
		`histogram_quantile(0.99, (sum by (le, endpoint)(label_replace(%s_bucket{app="%s"}, 
		"endpoint", "$1", "%s", "(.*)"))))`,
		conf.Conf.MetricName, conf.App, conf.Conf.EndpointLabel)

	r := queryPrometheus(queryGetEndpointsWithAvgRt, opts.PrometheusURL)
	buckets := getBuckets(fmt.Sprintf(`%s_bucket{app="%s"}`, conf.Conf.MetricName, conf.App), opts)
	switch res := r.(type) {
	case model.Vector:
		for _, v := range res {
			var e Endpoints
			eName := fmt.Sprint(v.Metric["endpoint"])
			if eName != "" {
				e.Name = eName
				e.Target = conf.Conf.TargetDefault

				avgRt, _ := strconv.ParseFloat(v.Value.String(), 64)
				defaultBucketNumber, _ := strconv.ParseFloat(conf.Conf.BucketDefaultThreshold, 64)

				for _, b := range buckets {
					if avgRt <= b.BucketNumber {
						if b.BucketNumber >= defaultBucketNumber {
							e.Bucket = b.BucketString
						} else {
							e.Bucket = conf.Conf.BucketDefaultThreshold
						}
						break
					}
				}
				if e.Bucket != "" {
					s.Spec.Slis.Latency.Endpoints = append(s.Spec.Slis.Latency.Endpoints, e)
				}
			}
		}
	}
	slices.SortFunc(s.Spec.Slis.Latency.Endpoints, CmpName)
}

func CmpName(a, b Endpoints) int {
	return cmp.Compare(a.Name, b.Name)
}

func AppLoadSpec(filename string, appName string) (sli apps.AppYamlSpec) {
	appList := apps.ReadAppsList(filename)
	for _, data := range appList {
		if data.Name == appName {
			err := mapstructure.Decode(data.Spec, &sli)
			if err != nil {
				panic(err)
			}
		}
	}
	return sli
}

func GetAppTypeFromAppList(filename string, appName string) (appType string) {
	appList := apps.ReadAppsList(filename)
	for _, data := range appList {
		if data.Name == appName {
			appType = data.Spec.Type
		}
	}
	return appType
}

func getBuckets(metric string, opts slogen.CmdOpts) []BucketHistogram {
	var buckets []BucketHistogram
	var b BucketHistogram
	match := []string{metric}
	res := labelValuesPrometheus("le", match, opts.PrometheusURL)
	for _, v := range res {
		b.BucketString = string(v)
		var validBucket = regexp.MustCompile(`^\d+(\.\d+)*$`)
		if validBucket.MatchString(b.BucketString) {
			b.BucketNumber, _ = strconv.ParseFloat(b.BucketString, 64)
			buckets = append(buckets, b)
		}
	}
	sort.SliceStable(buckets, func(i, j int) bool {
		return buckets[i].BucketNumber < buckets[j].BucketNumber
	})
	return buckets
}

func queryPrometheus(query string, prometheusURL string) model.Value {
	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, warnings, err := v1api.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	return result
}

func labelValuesPrometheus(label string, match []string, prometheusURL string) model.LabelValues {
	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, warnings, err := v1api.LabelValues(ctx, label, match, time.Now().Add(time.Minute*-1), time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	return result
}
