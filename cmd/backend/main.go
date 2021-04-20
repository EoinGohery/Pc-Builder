package backend

import (
	"net/http"
)

//Run backend, check to see if the server has been started
func Run() bool {
	_, err := http.Get("http://localhost:8080/api/parts/getCPUs")
	if err != nil {
		return true
	}
	return false
}
