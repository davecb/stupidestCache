package mvp

import (
	"github.com/davecb/stupidestCache/src/common"
	"time"
)

// MVP Cache -- a library to provide a not-as-simple cache for a
//	particular task, to compare with a performance target

type mvpCache struct {
	m      map[string]ve // the cache itself, not locked
	ttl    time.Duration // time to live for cache entries
	ask    chan string   // a channel for asking about keys
	answer chan ve       // a channel for returning values
	update chan kv       // a channel for updating keys & values
	refill chan kv       // a channel to ask for refills on
	done   chan bool     // and a channel for shutting down
}

// kv is just a key and value.
type kv struct {
	k string
	v string
}

// ve is a value, a place to put the "exists" response from the map call
// and a last-updated time.
type ve struct {
	value   string
	exists  bool
	updated time.Time
}

// New creates a new mvp cache
func New() common.Cache {
	var s mvpCache

	if s.m != nil {
		panic("you can't New() me twice, Dave! Die")
	}
	s.ttl = time.Minute * 1
	s.m = make(map[string]ve)
	s.ask = make(chan string)
	s.answer = make(chan ve)
	s.update = make(chan kv)
	s.refill = make(chan kv)
	s.done = make(chan bool)
	go s.mvp()
	go s.refiller()
	return s
}

// Get takes a key and returns a value and an exists flag
func (s mvpCache) Get(key string) (string, bool) {
	var ans ve

	s.ask <- key
	ans = <-s.answer
	return ans.value, ans.exists
}

// Put takes a key and value and updates a cache line
func (s mvpCache) Put(key, value string) error {
	var up kv

	up.k = key
	up.v = value
	s.update <- up
	return nil
}

// Close closes the cache
func (s mvpCache) Close() {
	close(s.done) // tell the goroutine to exit
	s.m = nil     // smash the map, to force freeing
}

// mvp -- the implementation, running as a goroutine
func (s mvpCache) mvp() {

	for {
		select {
		case k := <-s.ask:
			// when given a key, reply with a value and an exists flag

			var v ve
			val, exists := s.m[k]
			v.value = val.value
			v.updated = val.updated
			v.exists = exists // if false, the values above will be the defaults
			s.answer <- v

			// After replying, check TTL and refresh if stale
			if v.updated.IsZero() || time.Since(v.updated) > s.ttl {
				s.getFromL2(k)
			}

		case kv := <-s.update:
			// when given a key and value, replace both
			var v ve
			v.value = kv.v
			v.updated = time.Now()
			s.m[kv.k] = v

		case <-s.done:
			// when given a done indication, quit
			return
		}
	}
}

func (s mvpCache) getFromL2(k string) {
	var x kv
	var y ve
	
	x.v = k
	s.refill <- x  // call the refiller
	x = <-s.refill // get the response
	y.value = x.v
	y.updated = time.Now()
	s.m[k] = y
}

// refiller, the call to the L2 cache, running as a goroutine
func (s mvpCache) refiller() {
	var v kv

	for {
		select {
		case k := <-s.refill:
			// pretend to ask somebody for a new value here
			v.k = k.k
			v.v = "bogus"
			s.refill <- v

		case <-s.done:
			// when given a done indication, quit
			return
		}
	}
}
