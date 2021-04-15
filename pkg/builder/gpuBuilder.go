package builder

import (
	"io/ioutil"
	"log"
	"net/http"
)

type gpuBuilder struct {
	url []byte
}

func newGpuBuilder() *gpuBuilder {
	return &gpuBuilder{}
}

// func (b *gpuBuilder) setExtension() {
// 	b.extension = "http://localhost:8080/api/parts/getGPUs"
// }

func (b *gpuBuilder) sendRequest() {
	resp, err := http.Get("http://localhost:8080/api/parts/getGPUs")
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

func (b *gpuBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}
