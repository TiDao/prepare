module chainmaker.org/chainmaker-cryptogen

go 1.14

require (
	chainmaker.org/chainmaker-go/common v0.0.0-00010101000000-000000000000
	github.com/mr-tron/base58 v1.2.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	chainmaker.org/chainmaker-cryptogen/conf => ./conf
	chainmaker.org/chainmaker-go/common => ../common
)
