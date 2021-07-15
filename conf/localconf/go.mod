module chainmaker.org/chainmaker-go/localconf

go 1.15

require (
	chainmaker.org/chainmaker-go/common v0.0.0
	chainmaker.org/chainmaker-go/logger v0.0.0
	chainmaker.org/chainmaker-go/pb/protogo v0.0.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/spf13/viper v1.7.1
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace (
	chainmaker.org/chainmaker-go/common => ../../common
	chainmaker.org/chainmaker-go/logger => ../../logger
	chainmaker.org/chainmaker-go/pb/protogo => ../../protogo
)
