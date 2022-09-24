// Package inmem provides in-memory implementations of all the domain repositories.
package inmem

import (
	cargo2 "github.com/VolodymyrPobochii/go-mod-work/cargo"
	location2 "github.com/VolodymyrPobochii/go-mod-work/location"
	voyage2 "github.com/VolodymyrPobochii/go-mod-work/voyage"
	"sync"
)

type cargoRepository struct {
	mtx    sync.RWMutex
	cargos map[cargo2.TrackingID]*cargo2.Cargo
}

func (r *cargoRepository) Store(c *cargo2.Cargo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.cargos[c.TrackingID] = c
	return nil
}

func (r *cargoRepository) Find(id cargo2.TrackingID) (*cargo2.Cargo, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.cargos[id]; ok {
		return val, nil
	}
	return nil, cargo2.ErrUnknown
}

func (r *cargoRepository) FindAll() []*cargo2.Cargo {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	c := make([]*cargo2.Cargo, 0, len(r.cargos))
	for _, val := range r.cargos {
		c = append(c, val)
	}
	return c
}

// NewCargoRepository returns a new instance of a in-memory cargo repository.
func NewCargoRepository() cargo2.Repository {
	return &cargoRepository{
		cargos: make(map[cargo2.TrackingID]*cargo2.Cargo),
	}
}

type locationRepository struct {
	locations map[location2.UNLocode]*location2.Location
}

func (r *locationRepository) Find(locode location2.UNLocode) (*location2.Location, error) {
	if l, ok := r.locations[locode]; ok {
		return l, nil
	}
	return nil, location2.ErrUnknown
}

func (r *locationRepository) FindAll() []*location2.Location {
	l := make([]*location2.Location, 0, len(r.locations))
	for _, val := range r.locations {
		l = append(l, val)
	}
	return l
}

// NewLocationRepository returns a new instance of a in-memory location repository.
func NewLocationRepository() location2.Repository {
	r := &locationRepository{
		locations: make(map[location2.UNLocode]*location2.Location),
	}

	r.locations[location2.SESTO] = location2.Stockholm
	r.locations[location2.AUMEL] = location2.Melbourne
	r.locations[location2.CNHKG] = location2.Hongkong
	r.locations[location2.JNTKO] = location2.Tokyo
	r.locations[location2.NLRTM] = location2.Rotterdam
	r.locations[location2.DEHAM] = location2.Hamburg

	return r
}

type voyageRepository struct {
	voyages map[voyage2.Number]*voyage2.Voyage
}

func (r *voyageRepository) Find(voyageNumber voyage2.Number) (*voyage2.Voyage, error) {
	if v, ok := r.voyages[voyageNumber]; ok {
		return v, nil
	}

	return nil, voyage2.ErrUnknown
}

// NewVoyageRepository returns a new instance of a in-memory voyage repository.
func NewVoyageRepository() voyage2.Repository {
	r := &voyageRepository{
		voyages: make(map[voyage2.Number]*voyage2.Voyage),
	}

	r.voyages[voyage2.V100.Number] = voyage2.V100
	r.voyages[voyage2.V300.Number] = voyage2.V300
	r.voyages[voyage2.V400.Number] = voyage2.V400

	r.voyages[voyage2.V0100S.Number] = voyage2.V0100S
	r.voyages[voyage2.V0200T.Number] = voyage2.V0200T
	r.voyages[voyage2.V0300A.Number] = voyage2.V0300A
	r.voyages[voyage2.V0301S.Number] = voyage2.V0301S
	r.voyages[voyage2.V0400S.Number] = voyage2.V0400S

	return r
}

type handlingEventRepository struct {
	mtx    sync.RWMutex
	events map[cargo2.TrackingID][]cargo2.HandlingEvent
}

func (r *handlingEventRepository) Store(e cargo2.HandlingEvent) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	// Make array if it's the first event with this tracking ID.
	if _, ok := r.events[e.TrackingID]; !ok {
		r.events[e.TrackingID] = make([]cargo2.HandlingEvent, 0)
	}
	r.events[e.TrackingID] = append(r.events[e.TrackingID], e)
}

func (r *handlingEventRepository) QueryHandlingHistory(id cargo2.TrackingID) cargo2.HandlingHistory {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return cargo2.HandlingHistory{HandlingEvents: r.events[id]}
}

// NewHandlingEventRepository returns a new instance of a in-memory handling event repository.
func NewHandlingEventRepository() cargo2.HandlingEventRepository {
	return &handlingEventRepository{
		events: make(map[cargo2.TrackingID][]cargo2.HandlingEvent),
	}
}
