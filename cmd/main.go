package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/madacluster/gowaspsamm/pkg/owaspsamm"
	"gopkg.in/yaml.v3"
)

const owasp_path = "./owasp-core/model"

type DB struct {
	Questions          map[string]*owaspsamm.Question          `json:"questions"`
	BussinessFunctions map[string]*owaspsamm.BussinessFunction `json:"bussiness_functions"`
	SecurityPractices  map[string]*owaspsamm.SecurityPractice  `json:"security_practices"`
	Streams            map[string]*owaspsamm.Stream            `json:"streams"`
	Activities         map[string]*owaspsamm.Activity          `json:"activities"`
	MaturityLevels     map[string]*owaspsamm.MaturityLevel     `json:"maturity_levels"`
	Answers            map[string]*owaspsamm.Answer            `json:"answers"`
	PracticeLevels     map[string]*owaspsamm.PracticeLevel     `json:"practice_levels"`
	Path               string
}

func NewDB(path string) (*DB, error) {
	db := &DB{Path: path}
	db.Questions = make(map[string]*owaspsamm.Question)
	db.BussinessFunctions = make(map[string]*owaspsamm.BussinessFunction)
	db.SecurityPractices = make(map[string]*owaspsamm.SecurityPractice)
	db.Streams = make(map[string]*owaspsamm.Stream)
	db.Activities = make(map[string]*owaspsamm.Activity)
	db.MaturityLevels = make(map[string]*owaspsamm.MaturityLevel)
	db.Answers = make(map[string]*owaspsamm.Answer)
	db.PracticeLevels = make(map[string]*owaspsamm.PracticeLevel)
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
	if err := db.ReadFiles("bussiness_functions", db.readBussinessFunction); err != nil {
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
	q := owaspsamm.Question{}
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
	q := owaspsamm.BussinessFunction{}
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
	q := owaspsamm.SecurityPractice{}
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
	q := owaspsamm.Stream{}
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
	q := owaspsamm.Activity{}
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
	q := owaspsamm.PracticeLevel{}
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
	q := owaspsamm.MaturityLevel{}
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

func (db *DB) readAnswer(path string) error {
	q := owaspsamm.Answer{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	if err := yaml.Unmarshal(data, &q); err != nil {
		return err
	}
	db.Answers[q.ID] = &q
	return nil
}

func main() {
	db, err := NewDB(owasp_path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db.Questions["47c8fb0cae5944d090d7f73f7632dc9f"])

}
