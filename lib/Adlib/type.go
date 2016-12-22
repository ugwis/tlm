package adlib

import (
	"github.com/Goryudyuma/tlm/lib/List"
	"github.com/Goryudyuma/tlm/lib/User"
)

type Adlibs []Adlib

func (a *Adlibs) New(j []JsonAdlib) {
	for _, v := range j {
		var one Adlib
		one.New(v)
		*a = append(*a, one)
	}
}

type Adlib struct {
	List    list.List
	UserIDs user.UserIDs
}

func (a *Adlib) New(j JsonAdlib) {
	(*a).List.New(j.List)
	(*a).UserIDs.New(j.UserIDs)
}

type JsonAdlib struct {
	List    list.JsonList `json:"list"`
	UserIDs []int64       `json:"userids"`
}
