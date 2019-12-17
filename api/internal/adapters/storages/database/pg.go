package database

import (
	"github.com/assizkii/calendar/api/internal/domain/entities"
	"github.com/assizkii/calendar/api/internal/domain/interfaces"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type PgStorage struct {
	connection *gorm.DB
}

func (pg *PgStorage) Validate(event entities.Event) error {
	panic("implement me")
}

func (pg *PgStorage) Get(id string) (entities.Event, error) {
	panic("implement me")
}

func (pg *PgStorage) Add(e entities.Event) (string, error) {
	err := pg.connection.Create(e).Error
	return e.Id, err
}

func (pg *PgStorage) Update(id string, e entities.Event) error {
	panic("implement me")
}

func (pg *PgStorage) Delete(id string) error {
	panic("implement me")
}

func (pg *PgStorage) List() map[string]entities.Event {
	panic("implement me")
}

func (pg *PgStorage) FilterByDate(start time.Time) []entities.Event {
	panic("implement me")
}

func New(connection *gorm.DB) interfaces.EventStorage {
	return &PgStorage{connection}
}

func GetConnection() *gorm.DB   {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", "172.18.0.2", "postgres", "event_calendar"))
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.DB().SetMaxOpenConns(100)
	return db
}

func RunMigrations(DB *gorm.DB)  {
	DB.AutoMigrate(
		&entities.Event{},
	)
}