package stupidestCache

// Cache is the kind of cache we are making available over http.
// This one is the stupidest: there will be smarter ones.
type Cache interface {
	// Get returns a string and a presence flag
	Get(key string) (string, bool)
	// Put returns a sucess/failure indication
	Put(key, value string) error
	// Close shuts the cache down
	Close()
}

// Stupidest Cache -- a program to provide the simplest possible cache for a
//	particular task, to serve as a performance target

type sCache struct {
	m      map[string]string // the cache itself, not locked
	ask    chan string       // a channel for asking about keys
	answer chan ve           // a channel for returning values
	update chan kv           // a channel for updating keys & values
	done   chan bool         // and a channel for shutting down
}
type kv struct {
	k string
	v string
}

type ve struct {
	value  string
	exists bool
}

// New creates a new stupidest cache
func New() Cache {
	var s sCache

	if s.m != nil {
		panic("you can't New() me twice, Dave! Die")
	}
	s.m = make(map[string]string)
	s.ask = make(chan string)
	s.answer = make(chan ve)
	s.update = make(chan kv)
	s.done = make(chan bool)
	go s.stupid()
	return s
}

// Get takes a key and returns a value and an exists flag
func (s sCache) Get(key string) (string, bool) {
	var ans ve

	s.ask <- key
	ans = <-s.answer
	return ans.value, ans.exists
}

// Put takes a key and value and updates a cache line
func (s sCache) Put(key, value string) error {
	var up kv

	up.k = key
	up.v = value
	s.update <- up
	return nil
}

// Close closes the cache
func (s sCache) Close() {
	close(s.done) // tell the goroutine to exit
	s.m = nil     // smash the map, to force freeing
}

// stupid -- the implementation
func (s sCache) stupid() {

	for {
		select {
		case k := <-s.ask:
			var v ve
			// when given a key, reply with a value and an exists falg
			val, exists := s.m[k]
			v.value = val
			v.exists = exists
			s.answer <- v
		case kv := <-s.update:
			// when given a key and value, replace both
			s.m[kv.k] = kv.v
		case <-s.done:
			// when given a done indication, quit
			return
		}
	}
}
