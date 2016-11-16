package main

type Config struct {
	ConsumerKey    string `yaml:"ConsumerKey"`
	ConsumerSecret string `yaml:"ConsumerSecret"`
	SeedString     string `yaml:"SeedString"`
	CallbackURL    string `yaml:"CallbackURL"`
}

type List struct {
	OwnerID int64  `json:"ownerid"`
	ListID  int64  `json:"listid"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
}

type Job struct {
	Operator   string `json:"operator"`
	List1      List   `json:"list1"`
	List2      List   `json:"list2"`
	Listresult List   `json:"listresult"`
	Saveflag   bool   `json:"saveflag"`
}

type Query struct {
	Jobs    []Job  `json:"jobs"`
	Regular string `json:"regular"`
}
