package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Requests struct {
	Client *http.Client
}

func (r *Requests) NewPOSTRquest(url string, body []byte) error {

	request, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	resp, err := r.Client.Do(request)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Non Zero status code : %d", resp.StatusCode)
	}

	return nil
}

func (r *Requests) NewGETRequest(url string) ([]byte, error) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func buildURL(base string, route string) (string, error) {
	return fmt.Sprintf(base + route), nil //need to do build error checking with slashes and stuff
}
