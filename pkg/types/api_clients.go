package types

import (
	apiA "github.com/phev8/go_service_A/pkg/api"
	apiB "github.com/phev8/go_service_B/pkg/api"
)

// APIClients holds the service clients to the internal services
type APIClients struct {
	ServiceA apiA.ServiceAClient
	ServiceB apiB.ServiceBClient
}
