//go:generate go run ../generator/fields_generator.go -o fields_gen.go
package collector

import (
	"os/exec"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	cfg    *Config
	logger log.Logger
}

type Config struct {
}

func New(cfg *Config, logger log.Logger) (*Collector, error) {
	if cfg == nil {
		cfg = &Config{}
	}

	return &Collector{
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	err := execToMetrics("lvs", ch)
	if err != nil {
		level.Error(c.logger).Log("err", err)
	}

	err = execToMetrics("vgs", ch)
	if err != nil {
		level.Error(c.logger).Log("err", err)
	}

	err = execToMetrics("pvs", ch)
	if err != nil {
		level.Error(c.logger).Log("err", err)
	}
}

func execToMetrics(exe string, ch chan<- prometheus.Metric) error {
	cmd := exec.Command(exe,
		"--reportformat",
		"json",
		"-o",
		"all",
		"--units",
		"B",
		"--binary",
	)

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	metrics, err := ReportsToMetrics(out)
	if err != nil {
		return err
	}

	for _, ms := range metrics {
		var labelsNames, labelsValues []string
		for _, l := range ms.Labels {
			labelsNames = append(labelsNames, l.Field)
			labelsValues = append(labelsValues, l.Value)
		}

		for _, m := range ms.Metrics {
			gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace:   "lvm",
				Subsystem:   ms.Namespace,
				Name:        strings.TrimPrefix(m.Field, ms.Namespace+"_"),
				Help:        m.Help,
				ConstLabels: nil,
			}, labelsNames).WithLabelValues(labelsValues...)
			gauge.Set(m.Value)
			ch <- gauge
		}
	}

	return nil
}
