package main

import (
	"context"
	"github.com/qdriven/go-for-qa/bootstrap/petstore"
	"log"
	"net/http"
	"sync"
)

type petsService struct {
	pets map[int64]api.Pet
	id   int64
	mux  sync.Mutex
}

func (p *petsService) AddPet(ctx context.Context, req api.Pet) (api.Pet, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.pets[p.id] = req
	p.id++
	return req, nil
}

func (p *petsService) DeletePet(ctx context.Context, params api.DeletePetParams) (api.DeletePetOK, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	delete(p.pets, params.PetId)
	return api.DeletePetOK{}, nil
}

func (p *petsService) GetPetById(ctx context.Context, params api.GetPetByIdParams) (api.GetPetByIdRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet, ok := p.pets[params.PetId]
	if !ok {
		// Return Not Found.
		return &api.GetPetByIdNotFound{}, nil
	}
	return &pet, nil
}

func (p *petsService) UpdatePet(ctx context.Context, params api.UpdatePetParams) (api.UpdatePetOK, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet := p.pets[params.PetId]
	pet.Status = params.Status
	if val, ok := params.Name.Get(); ok {
		pet.Name = val
	}
	p.pets[params.PetId] = pet

	return api.UpdatePetOK{}, nil
}

func main() {
	// Create service instance.
	service := &petsService{
		pets: map[int64]api.Pet{},
	}
	// Create generated server.
	srv, err := api.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
