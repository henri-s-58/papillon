package common

import (
	"fmt"
	"gopkg.in/guregu/null.v4"
	"papillon/support"
	"reflect"
	"strconv"
)

var noNode = NewNode(-1, "", -1)

func NoNode() *Node {
	return noNode
}

type Node struct {
	id       int
	idString string
	host     string
	port     int
	rack     null.String
	hash     null.Int // for cache
}

func NewNode(id int, host string, port int) *Node {
	return NewNodeWithRack(id, host, port, null.NewString("", false))
}

func NewNodeWithRack(id int, host string, port int, rack null.String) *Node {
	return &Node{
		id:       id,
		idString: strconv.Itoa(id),
		host:     host,
		port:     port,
		rack:     rack,
		hash:     null.NewInt(-1, false),
	}
}

func (n *Node) IsEmpty() bool {
	return n.host == "" || n.port < 0
}

func (n *Node) Id() int {
	return n.id
}

func (n *Node) IdString() string {
	return n.idString
}

func (n *Node) Host() string {
	return n.host
}

func (n *Node) Port() int {
	return n.port
}

func (n *Node) HasRack() bool {
	return n.rack.Valid
}

func (n *Node) Rack() string {
	return n.rack.ValueOrZero()
}

func (n *Node) HashCode() int64 {
	hash := n.hash
	if hash.Valid {
		return n.hash.Int64
	}
	var result int64 = 31
	if n.host != "" {
		result += support.StrToHashCode(n.host)
	}
	result = 31*result + int64(n.id)
	result = 31*result + int64(n.port)
	result = 31 * result
	if n.rack.Valid {
		result += support.StrToHashCode(n.rack.String)
	}
	n.hash = null.NewInt(result, true)
	return result
}

func (n *Node) Equals(other interface{}) bool {
	return reflect.DeepEqual(n, other)
}

func (n *Node) String() string {
	return fmt.Sprintf(
		"%s:%d(id: %s rack: %s)",
		n.host, n.port, n.idString, n.rack.ValueOrZero(),
	)
}
