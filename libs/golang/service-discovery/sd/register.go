package service_discovery

import "sync"

type Service struct {
	Name    string
	Address string
	Port    int
}

type Registry struct {
	mu       sync.RWMutex
	services map[string]*Service
}

func NewRegistry() *Registry {
	return &Registry{
		services: make(map[string]*Service),
	}
}

func (r *Registry) RegisterService(name, address string, port int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services[name] = &Service{Name: name, Address: address, Port: port}
}

func (r *Registry) GetService(name string) *Service {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.services[name]
}

func (r *Registry) GetAllServices() []*Service {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var services []*Service
	for _, service := range r.services {
		services = append(services, service)
	}
	return services
}
