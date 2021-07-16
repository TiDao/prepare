module cryptogen

go 1.16

require (
	chainmaker.org/chainmaker-go/common v0.0.0
	github.com/mr-tron/base58 v1.2.0
	github.com/spf13/cobra v1.2.1
	gopkg.in/yaml.v2 v2.4.0
)

replace chainmaker.org/chainmaker-go/common => ../common
