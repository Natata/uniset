package uniset

// Uniset is the unique set
type Uniset map[interface{}]struct{}

// Add adds an element to the set
// true for adding success
// false for adding fail, the element is exist
func (u Uniset) Add(v interface{}) bool {
	_, ok := u[v]
	if ok {
		return false
	}

	u[v] = struct{}{}
	return true
}

// Remove removes an element from the set
func (u Uniset) Remove(v interface{}) {
	delete(u, v)
}

// List lists all element of the set
func (u Uniset) List() []interface{} {
	l := []interface{}{}
	for k := range u {
		l = append(l, k)
	}

	return l
}
