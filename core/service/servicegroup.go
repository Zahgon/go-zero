package service

type (
	Starter interface {
		Start()
	}

	Stopper interface {
		Stop()
	}

	Service interface {
		Starter
		Stopper
	}

	ServiceGroup struct {
		services []Service
		stopOnce func()
	}
)

func NewServiceGroup() *ServiceGroup { _ = "STUB: not implemented"; return nil }

func (sg *ServiceGroup) Add(service Service) { _ = "STUB: not implemented"; return }

func (sg *ServiceGroup) Start() { _ = "STUB: not implemented"; return }

func (sg *ServiceGroup) Stop() { _ = "STUB: not implemented"; return }

func (sg *ServiceGroup) doStart() { _ = "STUB: not implemented"; return }

func (sg *ServiceGroup) doStop() { _ = "STUB: not implemented"; return }

func WithStart(start func()) Service { _ = "STUB: not implemented"; return *new(Service) }

func WithStarter(start Starter) Service { _ = "STUB: not implemented"; return *new(Service) }

type (
	stopper struct{}

	startOnlyService struct {
		start func()
		stopper
	}

	starterOnlyService struct {
		Starter
		stopper
	}
)

func (s stopper) Stop() { _ = "STUB: not implemented"; return }

func (s startOnlyService) Start() { _ = "STUB: not implemented"; return }
