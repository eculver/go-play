package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	CMD_SET = iota
	CMD_DEL
)

// keep track of each write operation for a transaction

func main() {
	// simple: 1
	// simple1 := NewMemDB()
	// simple1.Set("a", "10")
	// fmt.Printf("got %s, expected 10\n", simple1.Get("a")) // 10
	// simple1.Del("a")
	// fmt.Printf("got %s, expected NULL\n", simple1.Get("a")) // NULL

	// ---------

	// simple: 2
	// simple2 := NewMemDB()
	// simple2.Set("a", "10")
	// simple2.Set("b", "10")
	// fmt.Printf("got %s, expected 2\n", simple2.Count())

	// ---------

	// transactions: 1
	trans1 := NewMemDB()
	trans1.Begin()
	trans1.Set("a", "10")
	fmt.Printf("got %s, expected 10\n", trans1.Get("a")) // 10
	trans1.Begin()
	trans1.Set("a", "20")
	fmt.Printf("got %s, expected 20\n", trans1.Get("a")) // 20
	if err := trans1.Rollback(); err != nil {
		log.Fatalf("ROLLBACK ERR: %v", err)
	}
	fmt.Printf("got %s, expected 10\n", trans1.Get("a")) // 10
	if err := trans1.Rollback(); err != nil {
		log.Fatalf("ROLLBACK ERR: %v", err)
	}
	fmt.Printf("got %s, expected NULL\n", trans1.Get("a")) // NULL
}

type Op struct {
	cmd  int
	key  string
	oldv string
	newv string
}

// MemDB is an in-memory KV datastore
type MemDB struct {
	db     map[string]string
	counts map[string]int // count -> key
	trans  [][]*Op
}

func NewMemDB() *MemDB {
	return &MemDB{
		db:    map[string]string{},
		trans: [][]*Op{},
	}
}

func (m MemDB) Apply(op *Op) {
	op.oldv = m.db[op.key]
	switch op.cmd {
	case CMD_SET:
		m.db[op.key] = op.newv
	case CMD_DEL:
		delete(m.db, op.key)
	}
}

func (m MemDB) Unapply(op *Op) {
	// TODO: remove the op
	switch op.cmd {
	case CMD_SET:
		fmt.Println("unapply set")
	case CMD_DEL:
		fmt.Println("unapply delete")
	}
}

func (m *MemDB) Get(k string) string {
	if len(m.trans) > 0 {
		// apply all transactions
		v := "NULL"
		for _, t := range m.trans {
			for _, op := range t {
				if op != nil && k == op.key {
					switch op.cmd {
					case CMD_SET:
						v = op.newv
					case CMD_DEL:
						v = "NULL"
					}
				}
			}
		}
		return v
	}
	if v, ok := m.db[k]; ok {
		return v
	}
	return "NULL"
}

func (m MemDB) Set(k, v string) {
	if len(m.trans) > 0 {
		oldv := "NULL"
		if v, ok := m.db[k]; ok {
			oldv = v
		}
		// register op
		op := &Op{
			cmd:  CMD_SET,
			key:  k,
			oldv: oldv,
			newv: v,
		}
		tidx := len(m.trans) - 1
		m.trans[tidx] = append(m.trans[tidx], op)
		return
	}
	m.db[k] = v
}

func (m MemDB) Del(k string) {
	if len(m.trans) > 0 {
		oldv := "NULL"
		if v, ok := m.db[k]; ok {
			oldv = v
		}
		// register op
		op := &Op{
			cmd:  CMD_DEL,
			key:  k,
			oldv: oldv,
			newv: "NULL",
		}
		tidx := len(m.trans) - 1
		m.trans[tidx] = append(m.trans[tidx], op)
		return
	}
	delete(m.db, k)
}

func (m *MemDB) Count(v string) string {
	// apply all operations
	return fmt.Sprintf("%d", m.counts[v])
}

func (m MemDB) Begin() {
	m.trans = append(m.trans, []*Op{})
}

func (m MemDB) Commit() error {
	if len(m.trans) == 0 {
		return errors.New("NO TRANSACTION")
	}
	for _, t := range m.trans {
		for _, o := range t {
			m.db[o.key] = o.newv
		}
	}
	return nil
}

func (m MemDB) Rollback() error {
	tidx := len(m.trans) - 1
	if len(m.trans) == 0 || tidx < 0 {
		return errors.New("NO TRANSACTION")
	}
	// m.trans[tidx] = m.trans[:tidx]
	return nil
}
