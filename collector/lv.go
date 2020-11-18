package collector

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	lvmLabels = map[string]bool{
		"lv_layout": true,
		"lv_host":   true,
		"lv_uuid":   true,
		"vg_uuid":   true,
		"pv_uuid":   true,
		"lv_name":   true,
		"vg_name":   true,
		"pv_name":   true,
	}
	lvmEnumsCustom = map[string][]string{
		"raid_sync_action":     {"idle", "frozen", "resync", "recover", "check", "repair"},
		"lv_health_status":     {"", "partial", "refresh needed", "mismatches exist"},
		"vg_allocation_policy": {"normal", "contiguous", "cling", "anywhere", "inherited"},
	}
)

type Metric struct {
	Field string
	Value float64
	Label *Label
	Help  string
}

type Label struct {
	Field string
	Value string
}

type MetricsLabels struct {
	Namespace string
	Metrics   []Metric
	Labels    []Label
}

func isIgnore(field string) bool {
	// nolint
	if strings.HasSuffix(field, "_attr") {
		return true
	}

	return false
}

func isLabel(field string) bool {
	return lvmLabels[field]
}

func toMetric(field string, value string) (*Metric, error) {
	enum, ok := lvmEnumsCustom[field]
	if !ok {
		enum, ok = lvmEnums[field]
	}
	if ok {
		mv := -1
		for i, v := range enum {
			if v == value {
				mv = i
			}
		}

		m := Metric{
			Field: field,
			Value: float64(mv),
		}

		f, ok := lvmFields[field]
		if ok {
			m.Help = f.Name + " - " + f.Description
		}

		m.Help += ` -1: undefined, `
		for i, s := range enum {
			m.Help += fmt.Sprintf(`%d: %s`, i, s)
			if i != len(enum)-1 {
				m.Help += `, `
			}
		}

		return &m, nil
	}

	f := lvmFields[field]

	m := Metric{
		Field: field,
		Help:  f.Name + " - " + f.Description,
	}

	switch f.Type {
	case FieldTypeSTR, FieldTypeSTR_LIST:
		return nil, nil
	case FieldTypeTIM:
		if value == "" {
			return &m, nil
		}

		// 2020-10-05 18:06:13 +0000
		t, err := time.Parse("2006-01-02 15:04:05 -0700", value)
		if err != nil {
			return nil, fmt.Errorf("field(%s): %w", field, err)
		}
		m.Value = float64(t.Unix())
		return &m, nil

	case FieldTypeBIN, FieldTypeNUM, FieldTypeSNUM, FieldTypeSIZ, FieldTypePCT:
		value = strings.TrimSuffix(value, "%")
		value = strings.TrimSuffix(value, "B")
	}

	switch value {
	case "unmanaged", "auto":
		value = ""
	}

	if value != "" {
		var err error
		m.Value, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("field(%s): %w", field, err)
		}
	} else {
		m.Value = -1
	}

	return &m, nil
}

func ReportsToMetrics(b []byte) ([]MetricsLabels, error) {
	reports, err := parseReports(b)
	if err != nil {
		return nil, err
	}

	var mls []MetricsLabels
	for _, report := range reports {
		for namespace, fieldsValues := range report {
			for _, fields := range fieldsValues {
				ml := MetricsLabels{Namespace: namespace}
				for field, value := range fields {
					if isIgnore(field) {
						continue
					}
					if isLabel(field) {
						ml.Labels = append(ml.Labels, Label{
							Field: field,
							Value: value,
						})
						continue
					}

					m, err := toMetric(field, value)
					if err != nil {
						return nil, err
					}
					if m != nil {
						ml.Metrics = append(ml.Metrics, *m)
					}
				}
				if len(ml.Metrics) > 0 {
					sort.Slice(ml.Metrics, func(i, j int) bool {
						return ml.Metrics[i].Field < ml.Metrics[j].Field
					})
					sort.Slice(ml.Labels, func(i, j int) bool {
						return ml.Labels[i].Field < ml.Labels[j].Field
					})

					mls = append(mls, ml)
				}
			}
		}
	}

	return mls, nil
}

type Reports []map[string][]map[string]string // []map[TYPE(lv/vg/pv)][]map[FIELD]VALUE

//
func parseReports(b []byte) (Reports, error) {
	var v struct {
		Report Reports
	}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}

	return v.Report, nil
}
