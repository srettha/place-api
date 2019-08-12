package models

import (
	"github.com/jinzhu/gorm"
)

// Place :: Place model
type Place struct {
	gorm.Model
	Name		string	`json:"name"`
	Latitude	string	`json:"latitude"`
	Longitude	string	`json:"longitude"`
}

// CreatePlace :: Create Place
type CreatePlace struct {
	Name		string	`json:"name" binding:"required"`
	Latitude	string	`json:"latitude" binding:"required"`
	Longitude	string	`json:"longitude" binding:"required"`
}

// UpdatePlace :: Update Place
type UpdatePlace struct {
	Name		string	`json:"name"`
	Latitude	string	`json:"latitude"`
	Longitude	string	`json:"longitude"`
}

// CreatePlace :: Create Place to DB
func (createPlace *CreatePlace) CreatePlace() (*Place, error) {
	place := &Place { Name: createPlace.Name, Latitude: createPlace.Latitude, Longitude: createPlace.Longitude }
	err := db.Create(place).Error
	if err != nil {
		return nil, err
	}

	return place, nil
}

// GetAllPlaces :: Get all Place from DB
func GetAllPlaces() ([] *Place, error) {
	places :=  make([]*Place, 0)
	err := db.Find(&places).Error
	if err != nil {
		return nil, err
	}

	return places, nil
}

// GetPlaceByID :: Get place by ID from DB
func GetPlaceByID(id int64) (*Place, error) {
	place := &Place {}
	err := db.First(&place, id).Error
	if err != nil {
		return nil, err
	}

	return place, nil
}

// UpdatePlaceByID :: Update place by ID from DB
func (updatePlace *UpdatePlace) UpdatePlaceByID(id int64) (error) {
	place := &Place {}
	err := db.First(&place, id).Update(&updatePlace).Error
	if err != nil {
		return err
	}

	return nil
}
