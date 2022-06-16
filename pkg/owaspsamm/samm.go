package owaspsamm

type Samm struct {
	BussinesFunctions map[string]BussinessFunction `json:"bussines_functions"`
}

type BussinessFunction struct {
	Model       string      `yaml:"model"`
	ID          string      `yaml:"id"`
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Color       interface{} `yaml:"color"`
	Logo        string      `yaml:"logo"`
	Order       int         `yaml:"order"`
	Type        string      `yaml:"type"`
}

type SecurityPractice struct {
	Function         string `yaml:"function"`
	ID               string `yaml:"id"`
	Name             string `yaml:"name"`
	ShortName        string `yaml:"shortName"`
	ShortDescription string `yaml:"shortDescription"`
	LongDescription  string `yaml:"longDescription"`
	Order            int    `yaml:"order"`
	Assignee         string `yaml:"assignee"`
	Progress         int    `yaml:"progress"`
	Type             string `yaml:"type"`
}

type Stream struct {
	Function         string `yaml:"function"`
	ID               string `yaml:"id"`
	Name             string `yaml:"name"`
	ShortName        string `yaml:"shortName"`
	ShortDescription string `yaml:"shortDescription"`
	LongDescription  string `yaml:"longDescription"`
	Order            int    `yaml:"order"`
	Assignee         string `yaml:"assignee"`
	Progress         int    `yaml:"progress"`
	Type             string `yaml:"type"`
}

type Activity struct {
	Stream            string      `yaml:"stream"`
	Level             string      `yaml:"level"`
	ID                string      `yaml:"id"`
	Title             string      `yaml:"title"`
	Benefit           string      `yaml:"benefit"`
	ShortDescription  string      `yaml:"shortDescription"`
	LongDescription   string      `yaml:"longDescription"`
	Results           interface{} `yaml:"results"`
	Metrics           interface{} `yaml:"metrics"`
	Costs             interface{} `yaml:"costs"`
	Personnel         interface{} `yaml:"personnel"`
	Notes             interface{} `yaml:"notes"`
	RelatedActivities interface{} `yaml:"relatedActivities"`
	Type              string      `yaml:"type"`
}

type PracticeLevel struct {
	Practice      string `yaml:"practice"`
	Maturitylevel string `yaml:"maturitylevel"`
	ID            string `yaml:"id"`
	Objective     string `yaml:"objective"`
	Type          string `yaml:"type"`
}

type MaturityLevel struct {
	ID          string `yaml:"id"`
	Number      int    `yaml:"number"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
}

type Question struct {
	Activity  string   `yaml:"activity"`
	Answerset string   `yaml:"answerset"`
	ID        string   `yaml:"id"`
	Text      string   `yaml:"text"`
	Order     int      `yaml:"order"`
	Quality   []string `yaml:"quality"`
	Type      string   `yaml:"type"`
}

type Answer struct {
	ID     string `yaml:"id"`
	Values []struct {
		Text   string `yaml:"text"`
		Value  int    `yaml:"value"`
		Weight int    `yaml:"weight"`
		Order  int    `yaml:"order"`
	} `yaml:"values"`
	Type string `yaml:"type"`
}
