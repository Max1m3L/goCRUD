package services

import (
	"cosmic-crud/models"
	"cosmic-crud/repositories"
	"fmt"
	"strings"
)

type SpaceshipService struct {
	repo *repositories.SpaceshipRepo
}

func NewSpaceshipService(repo *repositories.SpaceshipRepo) *SpaceshipService {
	return &SpaceshipService{repo: repo}
}

func (s *SpaceshipService) Create(ship *models.Spaceship) error {
	if strings.TrimSpace(ship.Name) == "" {
		return fmt.Errorf("имя корабля не может быть пустым")
	}
	if ship.Fuel < 0 {
		ship.Fuel = 0
	}
	if ship.CrewCount < 0 {
		ship.CrewCount = 0
	}

	ship.IsReady = ship.Fuel >= 100

	return s.repo.Create(ship)
}

func (s *SpaceshipService) GetAll() ([]models.Spaceship, error) {
	return s.repo.GetAll()
}

func (s *SpaceshipService) GetByID(id int) (*models.Spaceship, error) {
	ship, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if ship == nil {
		return nil, fmt.Errorf("корабль с ID %d не найден", id)
	}
	return ship, nil
}

func (s *SpaceshipService) Update(ship *models.Spaceship) error {
	if strings.TrimSpace(ship.Name) == "" {
		return fmt.Errorf("имя корабля не может быть пустым")
	}

	_, err := s.GetByID(ship.ID)
	if err != nil {
		return err
	}

	ship.IsReady = ship.Fuel >= 100

	return s.repo.Update(ship)
}

func (s *SpaceshipService) Delete(id int) error {
	deleted, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if !deleted {
		return fmt.Errorf("корабль с ID %d не найден", id)
	}
	return nil
}

func (s *SpaceshipService) Refuel(id int, amount int) (*models.Spaceship, error) {
	ship, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	ship.Fuel += amount
	ship.IsReady = ship.Fuel >= 100

	err = s.repo.Update(ship)
	if err != nil {
		return nil, err
	}

	return ship, nil
}
