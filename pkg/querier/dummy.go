package querier

import (
	"net/url"

	"github.com/prometheus/prometheus/scrape"
)

// DummyTargetRetriever implements TargetRetriever.
type DummyTargetRetriever struct{}

// Targets implements TargetRetriever.
func (r DummyTargetRetriever) Targets() []*scrape.Target { return nil }

// DummyAlertmanagerRetriever implements AlertmanagerRetriever.
type DummyAlertmanagerRetriever struct{}

// Alertmanagers implements AlertmanagerRetriever.
func (r DummyAlertmanagerRetriever) Alertmanagers() []*url.URL { return nil }
