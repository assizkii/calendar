# calendar

install 

`make build`

create event:

`./api create -t "eventTitle" -d "eventDescription" --start 2019-12-10 --owner 1 --end 2019-12-22`

update event:

`./api update --id 1 -t "eventTitleAnother" -d "eventDescriptionAnother" --start 2019-12-10 --owner 1 --end 2019-12-22`

event list by period (day, week, month):

`./api list -p week`

run grpc server:

`./api run_server`

run scheduler:

`./scheduler`

run sender:

`./sender`
