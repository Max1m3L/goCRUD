package repositories

import (
	"cosmic-crud/models"
	"database/sql"
	"fmt"
)

type SpaceshipRepo struct {
	DB *sql.DB
}

func NewSpaceshipRepo(db *sql.DB) *SpaceshipRepo {
	return &SpaceshipRepo{DB: db}
}

// CREATE
func (r *SpaceshipRepo) Create(ship *models.Spaceship) error {
	query := `
    INSERT INTO spaceships (name, fuel, is_ready, engine_type, crew_count)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id`

	err := r.DB.QueryRow(
		query,
		ship.Name,
		ship.Fuel,
		ship.IsReady,
		ship.EngineType,
		ship.CrewCount,
	).Scan(&ship.ID)

	if err != nil {
		return fmt.Errorf("ошибка создания корабля: %v", err)
	}
	return nil
}

// READ ALL
func (r *SpaceshipRepo) GetAll() ([]models.Spaceship, error) {
	rows, err := r.DB.Query("SELECT id, name, fuel, is_ready, engine_type, crew_count FROM spaceships")
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка: %v", err)
	}
	defer rows.Close()

	var ships []models.Spaceship
	for rows.Next() {
		var ship models.Spaceship
		err := rows.Scan(&ship.ID, &ship.Name, &ship.Fuel, &ship.IsReady, &ship.EngineType, &ship.CrewCount)
		if err != nil {
			return nil, err
		}
		ships = append(ships, ship)
	}
	return ships, nil
}

// READ ONE
func (r *SpaceshipRepo) GetByID(id int) (*models.Spaceship, error) {
	var ship models.Spaceship
	query := "SELECT id, name, fuel, is_ready, engine_type, crew_count FROM spaceships WHERE id = $1"

	err := r.DB.QueryRow(query, id).Scan(
		&ship.ID,
		&ship.Name,
		&ship.Fuel,
		&ship.IsReady,
		&ship.EngineType,
		&ship.CrewCount,
	)

	if err == sql.ErrNoRows {
		return nil, nil // не найдено
	}
	if err != nil {
		return nil, fmt.Errorf("ошибка получения корабля: %v", err)
	}
	return &ship, nil
}

// UPDATE
func (r *SpaceshipRepo) Update(ship *models.Spaceship) error {
	query := `
    UPDATE spaceships 
    SET name = $1, fuel = $2, is_ready = $3, engine_type = $4, crew_count = $5
    WHERE id = $6`

	result, err := r.DB.Exec(
		query,
		ship.Name,
		ship.Fuel,
		ship.IsReady,
		ship.EngineType,
		ship.CrewCount,
		ship.ID,
	)

	if err != nil {
		return fmt.Errorf("ошибка обновления: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("корабль с ID %d не найден", ship.ID)
	}
	return nil
}

// DELETE
func (r *SpaceshipRepo) Delete(id int) (bool, error) {
	result, err := r.DB.Exec("DELETE FROM spaceships WHERE id = $1", id)
	if err != nil {
		return false, fmt.Errorf("ошибка удаления: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0, nil
}
