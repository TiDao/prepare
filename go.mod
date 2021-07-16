module prepare

go 1.16

require(
		localconf v0.0.0
		cryptogen v0.0.0
)

replace(
		localconf => ./localconf
		cryptogen => ./cryptogen
		)
