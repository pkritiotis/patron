package log

import "github.com/prometheus/client_golang/prometheus"

var (
	logCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "observability",
			Subsystem: "log",
			Name:      "counter",
			Help:      "Counts logger calls per level",
		},
		[]string{"level"},
	)
)

func init() {
	prometheus.MustRegister(logCounter)
}

type metricDecoratorLogger struct {
	logger Logger
}

func NewLoggerWithMetrics(logger Logger) Logger {
	return metricDecoratorLogger{logger: logger}
}

func (l metricDecoratorLogger) Sub(m map[string]interface{}) Logger {
	return l.logger.Sub(m)
}

func (l metricDecoratorLogger) Fatal(i ...interface{}) {
	l.logger.Fatal(i)
}

func (l metricDecoratorLogger) Fatalf(s string, i ...interface{}) {
	l.Fatalf(s, i)
}

func (l metricDecoratorLogger) Panic(i ...interface{}) {
	l.logger.Panic(i)
}

func (l metricDecoratorLogger) Panicf(s string, i ...interface{}) {
	l.logger.Panicf(s, i)
}

func (l metricDecoratorLogger) Error(i ...interface{}) {
	l.logger.Error(i)
}

func (l metricDecoratorLogger) Errorf(s string, i ...interface{}) {
	l.logger.Errorf(s, i)
}

func (l metricDecoratorLogger) Warn(i ...interface{}) {
	l.logger.Warn(i)
}

func (l metricDecoratorLogger) Warnf(s string, i ...interface{}) {
	l.logger.Warnf(s, i)
}

func (l metricDecoratorLogger) Info(i ...interface{}) {
	l.logger.Info(i)
}

func (l metricDecoratorLogger) Infof(s string, i ...interface{}) {
	l.logger.Infof(s, i)
}

func (l metricDecoratorLogger) Debug(i ...interface{}) {
	l.logger.Debug(i)
}

func (l metricDecoratorLogger) Debugf(s string, i ...interface{}) {
	l.logger.Debugf(s, i)
}

func (l metricDecoratorLogger) Level() Level {
	return l.logger.Level()
}
