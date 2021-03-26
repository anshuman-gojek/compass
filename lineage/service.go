package lineage

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/odpf/columbus/models"
)

const (
	semUnlocked uint32 = iota
	semLocked
)

type TimeSource interface {
	Now() time.Time
}

type TimeSourceFunc func() time.Time

func (tsf TimeSourceFunc) Now() time.Time {
	return tsf()
}

type MetricsMonitor interface {
	Duration(string, int64)
}

// Service represents a high-level interface to the
// lineage package. Allows the client to configure
// an interval at which it'll construct the graph, while
// serving an old copy in between ticks.
type Service struct {
	typeRepo          models.TypeRepository
	recordRepoFactory models.RecordRepositoryFactory
	metricsMonitor    MetricsMonitor
	builder           Builder
	timeSource        TimeSource

	refreshInterval time.Duration
	lastBuilt       time.Time
	sem             uint32 // semaphore
	mu              sync.RWMutex
	graph           Graph
	err             error
}

func (srv *Service) build() {
	startTime := srv.timeSource.Now()
	graph, err := srv.builder.Build(srv.typeRepo, srv.recordRepoFactory)
	now := srv.timeSource.Now()

	srv.metricsMonitor.Duration("lineageBuildTime", int64(now.Sub(startTime)/time.Millisecond))

	srv.mu.Lock()
	defer srv.mu.Unlock()

	srv.graph = graph
	srv.err = err
	srv.lastBuilt = now
}

func (srv *Service) Graph() (Graph, error) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	srv.refreshIfNeeded()
	return srv.graph, srv.err
}

func (srv *Service) refreshIfNeeded() {
	threshold := srv.lastBuilt.Add(srv.refreshInterval)
	if srv.timeSource.Now().After(threshold) {
		srv.requestRefresh()
	}
}

func (srv *Service) requestRefresh() {
	// only one requestRefresh() call will be honored, and any and all following requests will be will be discarded until
	// the goroutine spawned by the former requestRefresh() is not finished.
	// WARN: do not touch this block of code unless you're _absolutely_ sure about what you're doing.
	if atomic.CompareAndSwapUint32(&srv.sem, semUnlocked, semLocked) {
		go func() {
			defer atomic.CompareAndSwapUint32(&srv.sem, semLocked, semUnlocked)
			srv.build()
		}()
	}
}

func NewService(er models.TypeRepository, rrf models.RecordRepositoryFactory, opts ...ServiceOpt) *Service {
	srv := &Service{
		builder:           DefaultBuilder,
		typeRepo:          er,
		recordRepoFactory: rrf,
		refreshInterval:   time.Minute,
		timeSource:        TimeSourceFunc(time.Now),
		metricsMonitor:    dummyMetricMonitor{},
	}
	for _, opt := range opts {
		opt(srv)
	}

	// TODO: Find a solution to solve memory issue

	// Temporarily disable building lineage on service creation.
	// Columbus's memory keeps spiking when app is starting
	// srv.build()

	return srv
}

type ServiceOpt func(*Service)

func WithRefreshInterval(d time.Duration) ServiceOpt {
	return func(srv *Service) {
		srv.refreshInterval = d
	}
}

func WithMetricMonitor(mm MetricsMonitor) ServiceOpt {
	return func(srv *Service) {
		srv.metricsMonitor = mm
	}
}

func WithBuilder(builder Builder) ServiceOpt {
	return func(srv *Service) {
		srv.builder = builder
	}
}

func WithTimeSource(ts TimeSource) ServiceOpt {
	return func(srv *Service) {
		srv.timeSource = ts
	}
}