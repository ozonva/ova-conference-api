package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	CreateConferenceEvent()
	MultiCreateConferenceEvent()
	UpdateConferenceEvent()
	DeleteConferenceEvent()
}

type metrics struct {
	createSuccessCounter      prometheus.Counter
	multiCreateSuccessCounter prometheus.Counter
	updateSuccessCounter      prometheus.Counter
	deleteSuccessCounter      prometheus.Counter
}

func NewMetrics() Metrics {
	return &metrics{
		createSuccessCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "create_count_total",
			Help: "Total count of successful created conferences",
		}),
		multiCreateSuccessCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "multi_create_count_total",
			Help: "Total count of successful multiCreated conferences",
		}),
		updateSuccessCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "update_count_total",
			Help: "Total count of successful updated conferences",
		}),
		deleteSuccessCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "delete_count_total",
			Help: "Total count of successful deleted conferences",
		}),
	}
}

func (m *metrics) CreateConferenceEvent() {
	m.createSuccessCounter.Inc()
}

func (m *metrics) MultiCreateConferenceEvent() {
	m.multiCreateSuccessCounter.Inc()
}

func (m *metrics) UpdateConferenceEvent() {
	m.updateSuccessCounter.Inc()
}

func (m *metrics) DeleteConferenceEvent() {
	m.deleteSuccessCounter.Inc()
}
