package command


func checkPort(port int,min,max int) bool{
	if (port >= min && port <= max) {
		return true
	}else{
		return false
	}
}
