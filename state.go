package main

import (
	"encoding/gob"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

type State struct {
	sync.Mutex
	stateFile string
}

func NewState(path string) *State {
	return &State{
		stateFile: path,
	}
}

func (s *State) Save(sites []SiteResult) {
	file, err := os.Create(s.stateFile)
	if err != nil {
		log.WithField("StateFile", s.stateFile).WithError(err).Error("Unable to open state file")
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.WithField("StateFile", s.stateFile).WithError(err).Error("Unable to close state file")
		}
	}()

	enc := gob.NewEncoder(file)

	err = enc.Encode(sites)
	if err != nil {
		log.WithError(err).Error("Unable to write to state file")
	}
}

func (s *State) Load() ([]SiteResult, error) {
	var sites []SiteResult

	file, err := os.Open(s.stateFile)
	if err != nil {
		log.WithField("StateFile", s.stateFile).WithError(err).Error("Unable to open state file")
		return sites, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.WithField("StateFile", s.stateFile).WithError(err).Error("Unable to close state file")
		}
	}()

	dec := gob.NewDecoder(file)

	err = dec.Decode(&sites)
	if err != nil {
		log.WithError(err).Error("Unable to decode state file")
		return sites, err
	}

	return sites, nil
}
