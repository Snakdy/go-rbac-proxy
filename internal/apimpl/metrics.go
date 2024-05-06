package apimpl

import (
	metric2 "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

var (
	meter        = metric.NewMeterProvider().Meter("go-rbac-proxy")
	metricAdd, _ = meter.Int64Counter(
		"rbac.role.add",
		metric2.WithDescription("Measures the total number of roles added."),
	)
	metricAddGlobal, _ = meter.Int64Counter(
		"rbac.role.add_global",
		metric2.WithDescription("Measures the total number of global roles added."),
	)
	metricCan, _ = meter.Int64Counter(
		"rbac.role.can",
		metric2.WithDescription("Measures the total number of RBAC checks."),
	)
)

const (
	attributeSourceKey = "role.source"

	sourceGlobal = "global_role"
	sourceRole   = "role"
	sourceNone   = "none"
)
