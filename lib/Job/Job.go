package job

import (
	"sync"

	"github.com/Goryudyuma/tlm/lib/Change"
	"github.com/Goryudyuma/tlm/lib/List"
	"github.com/Goryudyuma/tlm/lib/User"

	"github.com/bgpat/twtr"
)

func (v Job) dojob(client *twtr.Client, result, origin *map[list.List]user.UserIDs,
	ret *map[list.ListID]change.Change, listids *map[list.List]list.ListID) error {
	l1 := (*result)[v.List1]
	l2 := (*result)[v.List2]
	switch v.Operator {
	case "*":
		(*result)[v.Listresult] = l1.Intersect(l2)
	case "+":
		(*result)[v.Listresult] = l1.Union(l2)
	case "-":
		(*result)[v.Listresult] = l1.Except(l2)
	}

	if v.Config.Saveflag {
		addval := change.Change{
			AddList: (*result)[v.Listresult].Except((*origin)[v.Listresult]),
			DelList: (*origin)[v.Listresult].Except((*result)[v.Listresult])}

		listid, ok := (*listids)[v.Listresult]
		if !ok {
			listid = v.Listresult.ListID

			if listid == 0 {
				var mode string
				if v.Config.Publicflag {
					mode = "public"
				} else {
					mode = "private"
				}

				createlist, err := client.CreateList(twtr.Values{
					"name": v.Config.Name,
					"mode": mode,
				})
				if err != nil {
					return err
				}

				listid = list.ListID(createlist.ID)
			}
			(*listids)[v.Listresult] = listid
		}
		(*ret)[listid] = addval
	}
	return nil
}

func (jobs Jobs) Task(client *twtr.Client, origin map[list.List]user.UserIDs) (
	change.Changes, error) {
	listids := make(map[list.List]list.ListID)
	ret := make(map[list.ListID]change.Change)

	result := make(map[list.List]user.UserIDs, len(origin))
	for k, v := range origin {
		result[k] = v
	}

	for _, v := range jobs {
		v.dojob(client, &result, &origin, &ret, &listids)
	}
	return ret, nil
}

func (j Jobs) Getalllist(client *twtr.Client) (map[list.List]user.UserIDs, error) {
	var mutex sync.Mutex
	ret := make(map[list.List]user.UserIDs)
	chanerr := make(chan error, len(j)*3+1)
	defer close(chanerr)
	for _, v := range j {
		if _, ok := ret[v.List1]; !ok {
			go v.List1.GetListMembers(client, chanerr, &ret, &mutex)
		} else {
			chanerr <- nil
		}
		if _, ok := ret[v.List2]; !ok {
			go v.List2.GetListMembers(client, chanerr, &ret, &mutex)
		} else {
			chanerr <- nil
		}
		if _, ok := ret[v.Listresult]; !ok {
			go v.Listresult.GetListMembers(client, chanerr, &ret, &mutex)
		} else {
			chanerr <- nil
		}
	}

	var err error
	for i := 0; i < len(j)*3; i++ {
		select {
		case v := <-chanerr:
			{
				if v != nil {
					err = v
				}
			}
		}
	}
	return ret, err
}
