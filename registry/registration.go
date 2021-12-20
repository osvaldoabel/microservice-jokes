package registry

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}
