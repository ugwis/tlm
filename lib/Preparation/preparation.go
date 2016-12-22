package preparation

import "github.com/Goryudyuma/tlm/lib/List"
import "github.com/Goryudyuma/tlm/lib/User"

import "github.com/bgpat/twtr"

func (p Preparation) Prepare(client *twtr.Client) (map[list.List]user.UserIDs, error) {
	ret := make(map[list.List]user.UserIDs)
	for _, v := range p.Adlib {
		ret[v.List] = v.UserIDs
	}

	//ここでフォロワー一覧のIDを取ってくる

	return ret, nil
}
