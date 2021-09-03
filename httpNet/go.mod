module httpnet

go 1.16

require cryptogen v0.0.0

replace (
	chainmaker.org/chainmaker-go/common => ../common
	chainmaker.org/chainmaker-go/logger => ../logger
	chainmaker.org/chainmaker-go/pb/protogo => ../protogen
	cryptogen => ../cryptogen
)
