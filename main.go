package main

import (
	"calendar/internal/calendar"
	"calendar/internal/pkg/mngtservice"
	"calendar/internal/pkg/storage"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

//go:generate protoc --go_out=paths=source_relative:. internal/pkg/mngtservice/Event.proto

var (
	appConf calendar.Config
	confPath string
)

func init()  {
	flag.StringVar(&confPath, "config", "", "configs/conf.yaml")
	flag.Parse()
	readConfig()
}

func main() {
	file, err := os.OpenFile(appConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	calendar.RunServer(appConf)

	generateEvents()
}

func generateEvents()  {
	eventStorage := storage.New()

	newYear := time.Date(2019, 12, 31, 23, 59, 59, 0, time.UTC)
	eventStart := new(timestamp.Timestamp)
	eventEnd := new(timestamp.Timestamp)

	for i := 1; i < 10; i++ {
		finish := newYear.AddDate(0, 0, 1)
		eventStart.Seconds = int64(newYear.Unix())
		eventEnd.Seconds = int64(finish.Unix())

		p := new(mngtservice.Event)
		p.Id = int32(i)
		p.Title = "Vacation day " + strconv.Itoa(i)
		p.Description = "Event_" + strconv.Itoa(i) + "_Description"
		p.From = eventStart
		p.To = eventEnd
		eventStorage.Add(p)
		newYear = newYear.AddDate(0, 0, 1)
	}

	eventList, err := eventStorage.List()

	var eventIds []int
	for k := range eventList {
		eventIds = append(eventIds, int(k))
	}

	sort.Ints(eventIds)

	if err != nil {
		log.Fatal(err)
	}
	for _, eventID := range eventIds {
		event := eventList[int32(eventID)]
		if err != nil {
			log.Fatal(err)
		}
		timeFrom := time.Unix(event.From.GetSeconds(), 0)
		timeTo := time.Unix(event.To.GetSeconds(), 0)
		fmt.Printf("%s starts %v and ends %v\n", event.Title, timeFrom.Format("2 Jan 2006"), timeTo.Format("2 Jan 2006"))
	}
}


func readConfig() {

	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &appConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}