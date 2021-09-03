module command

go 1.16

require cryptogen v0.0.0

replace (
	chainmaker.org/chainmaker-go/common => ../common
	cryptogen => ../cryptogen
)
