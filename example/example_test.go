package example

import (
	"testing"

	. "gopkg.in/check.v1"
	"gopkg.in/maxwellhealth/bongo.v0"
)

const (
	testDatabase = "mongogen-test"
)

func Test(t *testing.T) { TestingT(t) }

type MongoSuite struct {
	conn *bongo.Connection
}

var _ = Suite(&MongoSuite{})

func (s *MongoSuite) SetUpTest(c *C) {
	s.conn, _ = bongo.Connect(&bongo.Config{
		ConnectionString: "localhost",
		Database:         testDatabase,
	})
}

func (s *MongoSuite) TestQuery_FindByFoo(c *C) {
	store := NewMyModelStore(s.conn)
	m := store.New()
	m.Foo = "foo"

	c.Assert(store.Insert(m), IsNil)

	q := store.Query()
	q.FindByFoo("foo")

	r, err := store.Find(q)
	c.Assert(err, IsNil)

	res, err := r.All()
	c.Assert(res, HasLen, 1)
	c.Assert(err, IsNil)

	q.FindByFoo("bar")
	r, err = store.Find(q)
	c.Assert(err, IsNil)

	one, err := r.One()
	c.Assert(one, IsNil)
	c.Assert(err, IsNil)
}