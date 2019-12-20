# calendar

create event:

`go run api/main.go create -t "eventTitle" -d "eventDescription" --start 2019-12-10 --owner 1 --end 2019-12-22`

update event:

`go run api/main.go update --id 1 -t "eventTitleAnother" -d "eventDescriptionAnother" --start 2019-12-10 --owner 1 --end 2019-12-22`

event list by period (day, week, month):

`go run api/main.go list -p week`
