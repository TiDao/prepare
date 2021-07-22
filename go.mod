module prepare

go 1.16

require (
	cryptogen v0.0.0
	localconf v0.0.0
	chainmaker.org/chainmaker-go/pb/protogo v0.0.0
)

replace (
	chainmaker.org/chainmaker-go/common => ./common
	chainmaker.org/chainmaker-go/logger => ./logger
	chainmaker.org/chainmaker-go/pb/protogo => ./protogo

	cryptogen => ./cryptogen
	localconf => ./localconf
)
