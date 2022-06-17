package owaspsamm

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type DB struct {
	Questions          map[string]*Question          `json:"questions"`
	BussinessFunctions map[string]*BussinessFunction `json:"bussiness_functions"`
	SecurityPractices  map[string]*SecurityPractice  `json:"security_practices"`
	Streams            map[string]*Stream            `json:"streams"`
	Activities         map[string]*Activity          `json:"activities"`
	MaturityLevels     map[string]*MaturityLevel     `json:"maturity_levels"`
	AnswersTypes       map[string]*AnswerTypeSet     `json:"answers"`
	PracticeLevels     map[string]*PracticeLevel     `json:"practice_levels"`
	Responses          map[string]*Answer            `json:"responses"`
	Path               string
}

func NewDB(path string) (*DB, error) {
	db := &DB{Path: path}
	db.Questions = make(map[string]*Question)
	db.BussinessFunctions = make(map[string]*BussinessFunction)
	db.SecurityPractices = make(map[string]*SecurityPractice)
	db.Streams = make(map[string]*Stream)
	db.Activities = make(map[string]*Activity)
	db.MaturityLevels = make(map[string]*MaturityLevel)
	db.AnswersTypes = make(map[string]*AnswerTypeSet)
	db.PracticeLevels = make(map[string]*PracticeLevel)
	db.Responses = make(map[string]*Answer)
	if err := db.ReadFolder(); err != nil {
		return nil, err
	}
	return db, nil
}

func (db *DB) ReadFolder() error {
	if err := db.ReadFiles("questions", db.readQuestion); err != nil {
		return err
	}

	if err := db.ReadFiles("activities", db.readActivity); err != nil {
		return err
	}
	if err := db.ReadFiles("business_functions", db.readBussinessFunction); err != nil {
		return err
	}
	if err := db.ReadFiles("maturity_levels", db.readMaturityLevel); err != nil {
		return err
	}
	if err := db.ReadFiles("practice_levels", db.readPracticeLevel); err != nil {
		return err
	}
	if err := db.ReadFiles("security_practices", db.readSecurityPractice); err != nil {
		return err
	}
	if err := db.ReadFiles("streams", db.readStream); err != nil {
		return err
	}
	if err := db.ReadFiles("answer_sets", db.readAnswerTypeSet); err != nil {
		return err
	}
	return nil
}

func (db *DB) ReadFiles(folder string, read func(path string) error) error {
	files, err := filepath.Glob(db.Path + "/" + folder + "/*.yml")
	if err != nil {
		return err
	}
	for _, file := range files {
		err := read(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return nil
}

func (db *DB) readQuestion(path string) error {
	q := Question{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.Questions[q.ID] = &q
	return nil
}

func (db *DB) readBussinessFunction(path string) error {
	q := BussinessFunction{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.BussinessFunctions[q.ID] = &q
	return nil
}

func (db *DB) readSecurityPractice(path string) error {
	q := SecurityPractice{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.SecurityPractices[q.ID] = &q
	return nil
}

func (db *DB) readStream(path string) error {
	q := Stream{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.Streams[q.ID] = &q
	return nil
}

func (db *DB) readActivity(path string) error {
	q := Activity{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.Activities[q.ID] = &q
	return nil
}

func (db *DB) readPracticeLevel(path string) error {
	q := PracticeLevel{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.PracticeLevels[q.ID] = &q
	return nil
}

func (db *DB) readMaturityLevel(path string) error {
	q := MaturityLevel{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.MaturityLevels[q.ID] = &q
	return nil
}

func (db *DB) readAnswerTypeSet(path string) error {
	q := AnswerTypeSet{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.AnswersTypes[q.ID] = &q
	return nil
}

func (db *DB) GetActivityQuestions(id string) []*Question {
	var questions []*Question
	for _, q := range db.Questions {
		if q.Activity == id {
			questions = append(questions, q)
		}
	}
	return questions
}

func (db *DB) GetStreamQuestions(id string) []*Question {
	var questions []*Question
	for _, a := range db.Activities {
		if a.Stream == id {
			questions = append(questions, db.GetActivityQuestions(a.ID)...)
		}
	}
	return questions
}

func (db *DB) GetSecurityPracticeQuestions(id string) []*Question {
	var questions []*Question
	for _, s := range db.Streams {
		if s.Practice == id {
			questions = append(questions, db.GetStreamQuestions(s.ID)...)
		}
	}
	return questions
}

func (db *DB) GetBussinesFunctionsQuestions(id string) []*Question {

	var questions []*Question
	for _, s := range db.SecurityPractices {
		if s.Function == id {
			questions = append(questions, db.GetSecurityPracticeQuestions(s.ID)...)
		}
	}
	return questions
}

func (db *DB) GetRadarChartFunctions() map[string]float64 {
	var radarChart = make(map[string]float64)
	for _, b := range db.BussinessFunctions {
		result := 0.0
		for _, q := range db.GetBussinesFunctionsQuestions(b.ID) {
			resp := db.Responses[q.ID]
			result += db.AnswersTypes[q.Answerset].Values[resp.Value].Value
		}
		radarChart[b.Name] = result
	}
	return radarChart
}

func (db *DB) GetRadarChartPractices() map[string]float64 {
	var radarChart = make(map[string]float64)
	for _, b := range db.SecurityPractices {
		result := 0.0
		for _, q := range db.GetSecurityPracticeQuestions(b.ID) {
			resp := db.Responses[q.ID]
			result += db.AnswersTypes[q.Answerset].Values[resp.Value].Value
		}
		radarChart[b.Name] = result
	}
	return radarChart
}
