// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/waydxd/paopao-ce/internal/conf"
	"github.com/waydxd/paopao-ce/internal/core"
	"github.com/waydxd/paopao-ce/internal/events"
)

func scheduleJobs(metrics *metrics) {
	spec := conf.JobManagerSetting.UpdateMetricsInterval
	schedule, err := cron.ParseStandard(spec)
	if err != nil {
		panic(err)
	}
	events.OnTask(schedule, metrics.onUpdate)
	logrus.Debug("shedule prometheus metrics update jobs complete")
}

func NewHandler(ds core.DataService, wc core.WebCache) http.Handler {
	// Create non-global registry.
	registry := prometheus.NewRegistry()
	// Add go runtime metrics and process collectors.
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
	metrics := newMetrics(registry, ds, wc)
	scheduleJobs(metrics)
	return promhttp.HandlerFor(registry, promhttp.HandlerOpts{EnableOpenMetrics: true})
}
