package fromHttp

// http -- do a daemon serving http

import (
	"errors"
	"fmt"
	"github.com/davecb/stupidestCache/src/fromFile"
	"github.com/davecb/stupidestCache/src/stupidestCache"
	"io"
	"log"
	"net/http"
)

func Run() {

	// run it as a deamon
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var c stupidestCache.Cache
		var h httpOp

		h.cache = stupidestCache.New()
		defer c.Close()
		h.getOperation(w, r)
	})
	err := http.ListenAndServe(":80", mux)
	if errors.Is(err, http.ErrServerClosed) {
		// successful exit
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}

// code for a web service

type httpOp struct {
	cache stupidestCache.Cache
}

// getOperation will do the server get stuff
func (h *httpOp) getOperation(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request\n")
	//hasKey := r.URL.Query().Has("key")
	key := r.URL.Query().Get("key")
	//hasValue := r.URL.Query().Has("value")
	value := r.URL.Query().Get("value")
	m := fmt.Sprintf("get, key=%q, value=%q\n", key, value)
	log.Printf("%s", m)
	x, present := h.cache.Get(key)
	fromFile.InterpretGet(present, x, key, value)

	io.WriteString(w, m)
}
