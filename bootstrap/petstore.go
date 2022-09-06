package main

import (
	"context"
	petstore "github.com/qdriven/go-for-qa/petstore"
	"log"
	"net/http"
	"sync"
)

type petsService struct {
	pets map[int64]petstore.Pet
	id   int64
	mux  sync.Mutex
}

func (p *petsService) AddPet(ctx context.Context, req petstore.Pet) (petstore.Pet, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.pets[p.id] = req
	p.id++
	return req, nil
}

func (p *petsService) DeletePet(ctx context.Context, params petstore.DeletePetParams) (petstore.DeletePetOK, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	delete(p.pets, params.PetId)
	return petstore.DeletePetOK{}, nil
}

func (p *petsService) GetPetById(ctx context.Context, params petstore.GetPetByIdParams) (petstore.GetPetByIdRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet, ok := p.pets[params.PetId]
	if !ok {
		// Return Not Found.
		return &petstore.GetPetByIdNotFound{}, nil
	}
	return &pet, nil
}

func (p *petsService) UpdatePet(ctx context.Context, params petstore.UpdatePetParams) (petstore.UpdatePetOK, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet := p.pets[params.PetId]
	pet.Status = params.Status
	if val, ok := params.Name.Get(); ok {
		pet.Name = val
	}
	p.pets[params.PetId] = pet

	return petstore.UpdatePetOK{}, nil
}

func main() {
	// Create service instance.
	service := &petsService{
		pets: map[int64]petstore.Pet{},
	}
	// Create generated server.
	srv, err := petstore.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
