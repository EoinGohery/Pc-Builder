package builder

import (
	"io/ioutil"
	"log"
	"net/http"
)

type psBuilder struct {
	url []byte
}

func newPsBuilder() *psBuilder {
	return &psBuilder{}
}

func (b *psBuilder) sendRequest() {
	resp, err := http.Get("http://localhost:8080/api/parts/getPS")
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

func (b *psBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}
