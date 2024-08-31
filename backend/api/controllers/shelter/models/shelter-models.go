package models

import (
	"hack/core/models"
)

type ShelterPresenter struct {
	ID                   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name                 string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description          string `boil:"description" json:"description" toml:"description" yaml:"description"`
	Address              string `boil:"address" json:"address" toml:"address" yaml:"address"`
	Phone                string `boil:"phone" json:"phone" toml:"phone" yaml:"phone"`
	Capacity             int    `boil:"capacity" json:"capacity" toml:"capacity" yaml:"capacity"`
	CurrentOccupancy     int    `boil:"current_occupancy" json:"current_occupancy" toml:"current_occupancy" yaml:"current_occupancy"`
	CapacityPets         int    `boil:"capacity_pets" json:"capacity_pets" toml:"capacity_pets" yaml:"capacity_pets"`
	CurrentOccupancyPets int    `boil:"current_occupancy_pets" json:"current_occupancy_pets" toml:"current_occupancy_pets" yaml:"current_occupancy_pets"`
	Acessibility         bool   `boil:"acessibility" json:"acessibility" toml:"acessibility" yaml:"acessibility"`
	Email                string `boil:"email" json:"email" toml:"email" yaml:"email"`
}

func (presenter *ShelterPresenter) InitFromModel(shelter models.Shelter) {
	presenter.ID = shelter.ID
	presenter.Name = shelter.Name
	presenter.Description = shelter.Description
	presenter.Address = shelter.Address
	presenter.Phone = shelter.Phone
	presenter.Capacity = shelter.Capacity
	presenter.CurrentOccupancy = shelter.CurrentOccupancy
	presenter.CapacityPets = shelter.CapacityPets
	presenter.CurrentOccupancyPets = shelter.CurrentOccupancyPets
	presenter.Acessibility = shelter.Acessibility
	presenter.Email = shelter.Email

}
