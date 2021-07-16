module config

go 1.16

require (
	chainmaker.org/chainmaker-go/logger v0.0.0
	chainmaker.org/chainmaker-go/pb/protogo v0.0.0
	github.com/gogo/protobuf v1.3.2
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	chainmaker.org/chainmaker-go/common => ../common
	chainmaker.org/chainmaker-go/logger => ../logger
	chainmaker.org/chainmaker-go/pb/protogo => ../protogo
)
