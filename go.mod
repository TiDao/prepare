module prepare

go 1.16

require (
	chainmaker.org/chainmaker-go/pb/protogo v0.0.0
	command v0.0.0
	cryptogen v0.0.0
	k8s v0.0.0
	localconf v0.0.0
)

replace (
	chainmaker.org/chainmaker-go/common => ./common
	chainmaker.org/chainmaker-go/logger => ./logger
	chainmaker.org/chainmaker-go/pb/protogo => ./protogo

	command => ./command
	cryptogen => ./cryptogen
	k8s => ./k8s
	localconf => ./localconf
)
