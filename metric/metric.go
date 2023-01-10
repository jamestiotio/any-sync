package metric

import (
	"context"
	"github.com/anytypeio/any-sync/app"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

const CName = "common.metric"

func New() Metric {
	return new(metric)
}

type Metric interface {
	Registry() *prometheus.Registry
	app.ComponentRunnable
}

type metric struct {
	registry *prometheus.Registry
	config   Config
}

func (m *metric) Init(a *app.App) (err error) {
	m.registry = prometheus.NewRegistry()
	m.config = a.MustComponent("config").(configSource).GetMetric()
	return nil
}

func (m *metric) Name() string {
	return CName
}

func (m *metric) Run(ctx context.Context) (err error) {
	if err = m.registry.Register(collectors.NewBuildInfoCollector()); err != nil {
		return err
	}
	if err = m.registry.Register(collectors.NewGoCollector()); err != nil {
		return err
	}
	if m.config.Addr != "" {
		var errCh = make(chan error)
		http.Handle("/metrics", promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{}))
		go func() {
			errCh <- http.ListenAndServe(m.config.Addr, nil)
		}()
		select {
		case err = <-errCh:
		case <-time.After(time.Second / 5):
		}
	}
	return
}

func (m *metric) Registry() *prometheus.Registry {
	return m.registry
}

func (m *metric) Close(ctx context.Context) (err error) {
	return
}