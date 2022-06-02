package apimpl

import (
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/unit"
)

var (
	meter        = global.MeterProvider().Meter("go-rbac-proxy")
	metricAdd, _ = meter.SyncInt64().Counter(
		"rbac.role.add",
		instrument.WithUnit(unit.Dimensionless),
		instrument.WithDescription("Measures the total number of roles added."),
	)
	metricAddGlobal, _ = meter.SyncInt64().Counter(
		"rbac.role.add_global",
		instrument.WithUnit(unit.Dimensionless),
		instrument.WithDescription("Measures the total number of global roles added."),
	)
	metricCan, _ = meter.SyncInt64().Counter(
		"rbac.role.can",
		instrument.WithUnit(unit.Dimensionless),
		instrument.WithDescription("Measures the total number of RBAC checks."),
	)
)

const (
	attributeSourceKey = "role.source"

	sourceGlobal = "global_role"
	sourceRole   = "role"
	sourceNone   = "none"
)
