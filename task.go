package main

import (
	"log"
	"net/url"
	"strconv"
	"sync"

	"github.com/bgpat/twtr"
	"github.com/davecgh/go-spew/spew"
)

type taglistnode struct {
	list []int64
	ok   bool
}

func getlist(client *twitter.Client, id int64) ([]int64, error) {

	data, err := client.GetListMembers(url.Values{
		"list_id": {strconv.FormatInt(id, 10)},
		"count":   {"5000"},
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	ret := make([]int64, len(data.Users))

	for _, v := range data.Users {
		ret = append(ret, int64(v.ID))
	}

	return ret, nil
}

func jobtask(jobparam Job, listarr *map[List][]int64, client *twitter.Client) {

}

func getalllist(jobs []Job, client *twitter.Client) (map[List][]int64, error) {
	var mutex sync.Mutex
	var syn sync.WaitGroup
	ret := make(map[List][]int64)
	chanerr := make(chan error, len(jobs)+1)
	defer close(chanerr)
	for _, v := range jobs {
		if _, ok := ret[v.List1]; !ok {
			go func() {
				syn.Add(1)
				if v.List1.ListID != 0 {
					list, err := getlist(client, v.List1.ListID)
					if err != nil {
						chanerr <- err
					} else {
						chanerr <- nil
						mutex.Lock()
						ret[v.List1] = list
						mutex.Unlock()
					}
				} else {
					chanerr <- nil
					mutex.Lock()
					ret[v.Listresult] = []int64{}
					mutex.Unlock()

				}
				syn.Done()
			}()
		}
		if _, ok := ret[v.List2]; !ok {
			go func() {
				syn.Add(1)
				if v.List2.ListID != 0 {
					list, err := getlist(client, v.List2.ListID)
					if err != nil {
						chanerr <- err
					} else {
						chanerr <- nil
						mutex.Lock()
						ret[v.List2] = list
						mutex.Unlock()
					}
				} else {
					chanerr <- nil
					mutex.Lock()
					ret[v.Listresult] = []int64{}
					mutex.Unlock()
				}
				syn.Done()
			}()
		}
		if _, ok := ret[v.Listresult]; !ok {
			go func() {
				syn.Add(1)
				if v.Listresult.ListID != 0 {
					list, err := getlist(client, v.Listresult.ListID)
					if err != nil {
						chanerr <- err
					} else {
						chanerr <- nil
						mutex.Lock()
						ret[v.Listresult] = list
						mutex.Unlock()
					}
				} else {
					chanerr <- nil
					mutex.Lock()
					ret[v.Listresult] = []int64{}
					mutex.Unlock()
				}
				syn.Done()
			}()
		}
	}

	for _, _ = range jobs {
		select {
		case v := <-chanerr:
			{
				if v != nil {
					return ret, v
				}
			}
		}
	}
	syn.Wait()
	return ret, nil
}

func querytask(queryparam Query, client *twitter.Client) error {

	listarr, err := getalllist(queryparam.Jobs, client)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	spew.Dump(listarr)

	for _, v := range queryparam.Jobs {
		jobtask(v, &listarr, client)
	}

	return nil
}
