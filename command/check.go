package command


func CheckPort(port int,min,max int) bool{
	if (port >= min && port <= max) {
		return true
	}else{
		return false
	}
}
