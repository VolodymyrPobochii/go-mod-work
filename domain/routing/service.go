// Package routing provides the routing domain service. It does not actually
// implement the routing service but merely acts as a proxy for a separate
// bounded context.
package routing

import (
	cargo2 "github.com/VolodymyrPobochii/go-mod-work/domain/cargo"
)

// Service provides access to an external routing service.
type Service interface {
	// FetchRoutesForSpecification finds all possible routes that satisfy a
	// given specification.
	FetchRoutesForSpecification(rs cargo2.RouteSpecification) []cargo2.Itinerary
}
