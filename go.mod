module github.com/assizkii/calendar

go 1.13

replace github.com/assizkii/calendar/internal/pkg/storage v0.0.0 => ./internal/pkg/storage

replace github.com/assizkii/calendar/internal/pkg/mngtservice v0.0.0 => ./internal/pkg/mngtservice

require (
	github.com/golang/protobuf v1.3.2
	gopkg.in/yaml.v2 v2.2.4
)
