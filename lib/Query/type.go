package query

import (
	"github.com/Goryudyuma/tlm/lib/Job"
	"github.com/Goryudyuma/tlm/lib/Preparation"
)

type Query struct {
	preparation preparation.Preparation
	jobs        job.Jobs
	regularflag bool
}

func (query *Query) New(j JsonQuery) {
	(*query).preparation.New(j.Preparation)
	(*query).jobs.New(j.Jobs)
	(*query).regularflag = j.Regularflag
}

type JsonQuery struct {
	Preparation preparation.JsonPreparation `json:"preparation"`
	Jobs        []job.JsonJob               `json:"jobs"`
	Regularflag bool                        `json:"regularflag"`
}
