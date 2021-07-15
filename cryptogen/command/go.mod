module command

go 1.16

require (
	chainmaker.org/chainmaker-cryptogen/config v0.0.0
	chainmaker.org/chainmaker-go/common v0.0.0
	github.com/mr-tron/base58 v1.2.0
	github.com/spf13/cobra v1.2.1
)

replace (
	chainmaker.org/chainmaker-cryptogen/config => ../config
	chainmaker.org/chainmaker-go/common => ../../common
)
