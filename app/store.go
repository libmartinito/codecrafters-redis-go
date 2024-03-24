package main

type Store struct {
	data map[string]ValueWithExpiry
}

type ValueWithExpiry struct {
	value  string
	expiry int64
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]ValueWithExpiry),
	}
}

func (s *Store) Set(key, value string) {
	s.data[key] = ValueWithExpiry{value: value, expiry: -1}
}

func (s *Store) SetWithExpiry(key, value string, expiry int64) {
	s.data[key] = ValueWithExpiry{value: value, expiry: expiry}
}

func (s *Store) Get(key string) string {
	return s.data[key].value
}
