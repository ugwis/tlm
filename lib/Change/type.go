package change

import (
	"github.com/Goryudyuma/tlm/lib/List"
	"github.com/Goryudyuma/tlm/lib/User"
)

type Changes map[list.ListID]Change

type Change struct {
	AddList user.UserIDs
	DelList user.UserIDs
}

func (c *Change) New(j JsonChange) {
	(*c).AddList.New(j.AddList)
	(*c).DelList.New(j.DelList)
}

type JsonChange struct {
	AddList []int64
	DelList []int64
}
