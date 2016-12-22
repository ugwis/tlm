package query

import "github.com/bgpat/twtr"

func (q Query) Querytask(client *twtr.Client) error {
	preparearr, err := q.preparation.Prepare(client)

	listarr, err := q.jobs.Getalllist(client)
	if err != nil {
		return err
	}

	for k, v := range listarr {
		preparearr[k] = v
	}

	commitlist, err := q.jobs.Task(client, preparearr)
	if err != nil {
		return err
	}

	err = commitlist.Commit(client)
	if err != nil {
		return err
	}
	return nil
}
