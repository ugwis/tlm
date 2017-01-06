package preparation

import (
	"github.com/Goryudyuma/tlm/lib/Adlib"
	"github.com/Goryudyuma/tlm/lib/Follower"
)

type Preparation struct {
	Adlib    adlib.Adlibs
	Follower follower.Followers
}

func (p *Preparation) New(j JsonPreparation) {
	(*p).Adlib.New(j.Adlib)
	(*p).Follower.New(j.Follower)
}

type JsonPreparation struct {
	Adlib    []adlib.JsonAdlib       `json:"adlib"`
	Follower []follower.JsonFollower `json:"follower"`
}
