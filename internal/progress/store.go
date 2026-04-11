package progress

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Session struct {
	Date					string	`json:"date"`
	Level					int			`json:"level"`
	Reps					int			`json:"reps"`
	Squeeze				int			`json:"squeeze_seconds"`
	Rest					int			`json:"rest_seconds"`
}

type Store struct {
	CurrentLevel	int					`json:"current_level"`
	Sessions			[]Session		`json:"sessions"`
}

func (s *Store) AddSession(level, reps, squeeze, rest int) {
	s.Sessions = append(s.Sessions, Session{
		Date:			time.Now().Format("2006-01-02"),
		Level:		level,
		Reps:			reps,
		Squeeze:	squeeze,
		Rest:			rest,
	})
}

func (s *Store) SessionsAtCurrentLevel() int {
	count := 0
	for _, session := range s.Sessions {
		if session.Level == s.CurrentLevel {
			count ++
		}
	}
	return count
}

func (s *Store) ShouldAdvance(sessionsRequired int) bool {
	return s.SessionsAtCurrentLevel() >= sessionsRequired
}

func dataPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".kegel-trainer")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "progress.json"), nil
}

func Load() (Store, error) {
	path, err := dataPath()
	if err != nil {
		return Store{}, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return Store{}, nil
	}
	if err != nil {
		return Store{}, nil
	}

	var store Store
	err = json.Unmarshal(data, &store)
	return store, err
}

func Save(store Store) error {
	path, err := dataPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}