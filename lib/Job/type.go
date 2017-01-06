package job

import (
	"github.com/Goryudyuma/tlm/lib/List"
	"github.com/Goryudyuma/tlm/lib/ResultListConfig"
)

type Jobs []Job

func (j *Jobs) New(json []JsonJob) {
	for _, v := range json {
		var one Job
		one.New(v)
		*j = append(*j, one)
	}
}

type Job struct {
	Operator    string
	ListOne     list.List
	ListAnother list.List
	ListResult  list.List
	Config      resultlistconfig.ResultListConfig
}

func (j *Job) New(json JsonJob) {
	(*j).Operator = json.Operator
	(*j).ListOne.New(json.ListOne)
	(*j).ListAnother.New(json.ListAnother)
	(*j).ListResult.New(json.ListResult)
	(*j).Config.New(json.Config)
}

type JsonJob struct {
	Operator    string                                `json:"operator"`
	ListOne     list.JsonList                         `json:"listone"`
	ListAnother list.JsonList                         `json:"listanother"`
	ListResult  list.JsonList                         `json:"listresult"`
	Config      resultlistconfig.JsonResultListConfig `json:"config"`
}
