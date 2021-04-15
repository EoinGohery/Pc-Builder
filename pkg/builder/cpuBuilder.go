package builder

import (
	"io/ioutil"
	"log"
	"net/http"
)

type cpuBuilder struct {
	url []byte
}

func newCpuBuilder() *cpuBuilder {
	return &cpuBuilder{}
}

func (b *cpuBuilder) sendRequest() {
	resp, err := http.Get("http://localhost:8080/api/parts/getCPUs")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Printf("%s\n", body)
	b.url = body
}

func (b *cpuBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}
