package porto

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"golang.org/x/net/context"

	"github.com/stretchr/testify/assert"
)

func TestCursor(t *testing.T) {
	fixtureBody := map[string][]byte{
		"A": []byte("A"),
		"B": []byte("B"),
		"C": []byte("C"),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		body, ok := fixtureBody[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(body)
	}))

	var urls []string
	var names []string
	for file := range fixtureBody {
		url := fmt.Sprintf("%s?name=%s", ts.URL, file)
		urls = append(urls, url)
		names = append(names, file)
	}
	sort.Strings(urls)
	sort.Strings(names)

	defer ts.Close()
	cur := newCursor(context.Background(), "testcursor", urls, nil)

	var name string
	for {
		resp, err := cur.Next()
		if err != nil {
			assert.Equal(t, err, ErrCursorNoResult)
			break
		}
		name, names = names[0], names[1:]

		body, err := ioutil.ReadAll(resp)
		assert.NoError(t, err)
		assert.Equal(t, fixtureBody[name], body, name)
		resp.Close()
	}
}
