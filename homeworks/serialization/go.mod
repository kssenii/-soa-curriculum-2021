module github.com/kssenii/soa-curriculum-2021/homeworks/serialization

go 1.15

replace github.com/kssenii/soa-curriculum-2021/homeworks/serialization/data => ./data

require (
	github.com/golang/protobuf v1.4.3
	github.com/vmihailenco/msgpack/v5 v5.2.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0
)
