package avro

import "sync"

// programCache is a thread-safe cache for decodePrograms. I've added it because a decodeProgram is very expensive
// to construct, so we want to reuse them if we can
// The map key should be the string representation of the schema. This might seem a bit expensive but it's the only
// way to 100% ensure we're using the correct cached decoder for the given schema
type programCache struct {
	m   map[string]*decodeProgram
	mtx sync.Mutex
}

func (p *programCache) Set(key string, prog *decodeProgram) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	// because progamCache gets default initialised with an empty map (and we don't want to go around changing a bunch
	// of code), we initialise it on the first call to Set
	if p.m == nil {
		p.m = map[string]*decodeProgram{}
	}

	p.m[key] = prog
}

func (p *programCache) Get(key string) (*decodeProgram, bool) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.m == nil {
		return nil, false
	}

	prog, ok := p.m[key]
	return prog, ok
}
