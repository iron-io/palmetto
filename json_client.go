package palmetto

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type jsonClient struct {
	srvURL string
}

func JSONClient(srvURL string) Client {
	return &jsonClient{srvURL: srvURL}
}

func (c *jsonClient) fullURL(pathFmt string, vals ...interface{}) string {
	path := fmt.Sprintf(pathFmt, vals...)
	return fmt.Sprintf("%s/%s", c.srvURL, path)
}

func (c *jsonClient) Get(key string) (*PlainResult, error) {
	u := c.fullURL("kv/rest/entry/%s", key)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("resp: %+v\n", resp)
	fmt.Printf("body: %s\n", string(b))
	return nil, ErrNotFound
}

func (c *jsonClient) Put(key string, val []byte) (*PlainResult, error) {

	return nil, ErrNotFound
}
