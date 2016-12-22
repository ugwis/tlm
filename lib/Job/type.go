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
	Operator   string
	List1      list.List
	List2      list.List
	Listresult list.List
	Config     resultlistconfig.ResultListConfig
}

func (j *Job) New(json JsonJob) {
	(*j).Operator = json.Operator
	(*j).List1.New(json.List1)
	(*j).List2.New(json.List2)
	(*j).Listresult.New(json.Listresult)
	(*j).Config.New(json.Config)
}

type JsonJob struct {
	Operator   string                                `json:"operator"`
	List1      list.JsonList                         `json:"list1"`
	List2      list.JsonList                         `json:"list2"`
	Listresult list.JsonList                         `json:"listresult"`
	Config     resultlistconfig.JsonResultListConfig `json:"config"`
}
