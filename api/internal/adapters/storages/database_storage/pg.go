package database_storage

import (
	"calendar/entities"
	"calendar/api/internal/domain/interfaces"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)


type PgStorage struct {
	connection *sqlx.DB
}

type EventDb struct {
	Id int `db:"id"`
	Title string `db:"title"`
	Description string `db:"description"`
	OwnerId int `db:"owner_id"`
	Start time.Time `db:"start"`
	EndTime time.Time `db:"end_time"`
}

func (pg *PgStorage) Validate(event entities.Event) error {

	switch "" {
	case strings.TrimSpace(event.Description):
		return errors.New("description field cannot be empty string")
	case strings.TrimSpace(event.Title):
		return errors.New("title field cannot be empty string")
	}

	return nil
}

func (pg *PgStorage) Get(id string) (entities.Event, error) {

	var event entities.Event
	var eventRow EventDb
	query := `select *	from events where id=$1`

	err := pg.connection.Get(&eventRow, query, id)
	if err != nil {
		return event, err
	}
	eventStart, err := ptypes.TimestampProto(eventRow.Start)
	eventEnd, err :=  ptypes.TimestampProto(eventRow.EndTime)
	if err != nil {
		return event, err
	}

	event = entities.Event{
		Id:     strconv.Itoa(eventRow.Id),
		OwnerId:     int32(eventRow.OwnerId),
		Title:       eventRow.Title,
		Description: eventRow.Description,
		Start:      eventStart,
		EndTime: eventEnd,


	}

	return event, nil
}

func (pg *PgStorage) Add(e entities.Event) (string, error) {

	if err := pg.Validate(e); err != nil {
		return "", fmt.Errorf("validate error: %s", err)
	}

	query := `insert into events(owner_id, title, description, start, end_time)
				 values($1, $2, $3, $4, $5) RETURNING id`
	endTime, err := ptypes.Timestamp(e.GetEndTime())
	startTime, err := ptypes.Timestamp(e.GetStart())
	if err != nil {
		return "", err
	}

	var id int
	err = pg.connection.QueryRow(query, e.OwnerId, e.Title, e.Description, startTime, endTime).Scan(&id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(id), nil
}

func (pg *PgStorage) Update(id string, e entities.Event) error {

	if err := pg.Validate(e); err != nil {
		return fmt.Errorf("validate error: %s", err)
	}

	query := `update events set owner_id=$2, title=$3, description=$4, start=$5, end_time=$6 where id=$1`
	endTime, err := ptypes.Timestamp(e.GetEndTime())
	startTime, err := ptypes.Timestamp(e.GetStart())
	if err != nil {
		return err
	}

	_, err = pg.connection.Exec(query, id, e.OwnerId, e.Title, e.Description, startTime, endTime)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PgStorage) Delete(id string) error {

	query := `delete from events where id=$1`

	_, err := pg.connection.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PgStorage) List() (map[string]entities.Event, error){

	var result map[string]entities.Event
	var events []entities.Event

	query := `select * from events`
	err := pg.connection.Select(&events, query)
	if err != nil {
		return nil, err
	}

	for _, event := range events  {
		result[event.GetId()] = event
	}
	return result, nil
}

func (pg *PgStorage) FilterByDate(startTime time.Time, endTime time.Time) ([]entities.Event, error) {

	var events []entities.Event
	var eventsRows []EventDb
	query := `select id, title, description, owner_id, start, end_time
					from events where (start >= $1 and end_time <= $2) or (end_time >= $1 and end_time <= $2)`

	err := pg.connection.Select(&eventsRows, query, startTime, endTime)
	if err != nil {
		return nil, err
	}

	for _, eventRow := range eventsRows  {
		eventStart, err := ptypes.TimestampProto(eventRow.Start)
		eventEnd, err :=  ptypes.TimestampProto(eventRow.EndTime)
		if err != nil {
			return nil, err
		}
		event := entities.Event{
			Id:     strconv.Itoa(eventRow.Id),
			OwnerId:     int32(eventRow.OwnerId),
			Title:       eventRow.Title,
			Description: eventRow.Description,
			Start:      eventStart,
			EndTime: eventEnd,


		}
		events = append(events, event)
	}

	return events, nil
}

func New(dsn string) interfaces.EventStorage{

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	err = Migrate(db)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	return &PgStorage{db}
}


// sql migrations from migrations folder
func Migrate(db *sqlx.DB) error  {


	err := filepath.Walk("./migrations", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if !info.IsDir() {
			log.Println(path)
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			fileValue := string(fileData)
			_, err = db.Exec(fileValue)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}