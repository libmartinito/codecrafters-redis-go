package main

type Info struct {
	replication ReplicationInfo
}

type ReplicationInfo struct {
	masterReplid     string
	masterReplOffset int
	role             string
}

type Store struct {
	data map[string]ValueWithExpiry
	info Info
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

func (s *Store) Get(key string) string {
	return s.data[key].value
}

func (s *Store) InitInfo(replicaof string) {
	s.info.replication.masterReplid = "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
	s.info.replication.masterReplOffset = 0

	if replicaof != "" {
		s.info.replication.role = "slave"
	} else {
		s.info.replication.role = "master"
	}
}

func (s *Store) Set(key, value string) {
	s.data[key] = ValueWithExpiry{value: value, expiry: -1}
}

func (s *Store) SetWithExpiry(key, value string, expiry int64) {
	s.data[key] = ValueWithExpiry{value: value, expiry: expiry}
}
