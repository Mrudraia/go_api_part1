package newhandlers

import (
	"fmt"
	"net/http"
)

type Sample struct {
	prefix  string
	data    map[string]string
	message string
}

func TestSample(Prefix string) Sample {
	b := Sample{
		prefix:  Prefix,
		data:    map[string]string{"name": "siddu", "age": "26", "company": "Redhat"},
		message: "Welcome to Redhat, explore openshift....!",
	}
	return b

}

func (i *Sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := TestSample("Siddu")
	fmt.Fprintf(w, res.message)

}
