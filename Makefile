build:
	@echo "Build start"
	go build -o api ./pkg/api/main.go
	go build -o scheduler ./pkg/scheduler/main.go
	go build -o sender ./pkg/sender/main.go