package resultlistconfig

type ResultListConfig struct {
	Name       string
	Publicflag bool
	Saveflag   bool
}

func (r *ResultListConfig) New(j JsonResultListConfig) {
	(*r).Name = j.Name
	(*r).Publicflag = j.Publicflag
	(*r).Saveflag = j.Saveflag
}

type JsonResultListConfig struct {
	Name       string `json:"name"`
	Publicflag bool   `json:"publicflag"`
	Saveflag   bool   `json:"saveflag"`
}
