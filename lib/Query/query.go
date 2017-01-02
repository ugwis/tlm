package query

import (
	"github.com/bgpat/twtr"
	"github.com/davecgh/go-spew/spew"
)

func (q Query) Querytask(client *twtr.Client) error {
	preparearr, err := q.preparation.Prepare(client)

	listarr, err := q.jobs.Getalllist(client, &preparearr)
	if err != nil {
		return err
	}

	commitlist, err := q.jobs.Task(client, listarr)
	if err != nil {
		return err
	}

	err = commitlist.Commit(client)
	if err != nil {
		return err
	}
	return nil
}
