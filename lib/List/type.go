package list

import (
	"github.com/Goryudyuma/tlm/lib/Tag"
	"github.com/Goryudyuma/tlm/lib/User"
)

type ListID int64

func (l *ListID) New(j int64) {
	*l = ListID(j)
}

type List struct {
	OwnerID user.UserID
	ListID  ListID
	Tag     tag.Tag
}

func (l *List) New(j JsonList) {
	(*l).OwnerID.New(j.OwnerID)
	(*l).ListID.New(j.ListID)
	(*l).Tag.New(j.Tag)
}

type JsonList struct {
	OwnerID int64  `json:"ownerid"`
	ListID  int64  `json:"listid"`
	Tag     string `json:"tag"`
}
