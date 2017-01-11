package preparation

import (
	"github.com/Goryudyuma/tlm/lib/List"
	"github.com/Goryudyuma/tlm/lib/User"

	"github.com/bgpat/twtr"
)

func (p Preparation) Prepare(client *twtr.Client) (map[list.List]user.UserIDs, error) {
	ret := make(map[list.List]user.UserIDs)
	for _, v := range p.Adlib {
		ret[v.List] = v.UserIDs
	}

	//ここでフォロワー一覧のIDを取ってくる
	for _, v := range p.Follower {
		retval, err := v.GetFollowerIDs(client)
		if err != nil {
			return nil, err
		}
		ret[v.List] = retval
	}

	return ret, nil
}
