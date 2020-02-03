package restify

import (
	"regexp"
	"strings"
	"time"

	"github.com/SpaceStock/go-restify/enum"
)

var (
	replacable = regexp.MustCompile("\\{(.*?)\\}")
)

// Delay type string
type Delay string

//IsZero given boolean whether it's true or false
//Return false if delay not equal 0 or 0s or preceded by "-" or delay is error when parsing to time duration
func (d Delay) IsZero() bool {
	delay := string(d)
	if delay == "0" || delay == "0s" || strings.HasPrefix("-", delay) {
		return true
	}

	_, err := time.ParseDuration(delay)
	if err != nil {
		return true
	}

	return false
}

//Pipeline test pipeline as what to do with the response object
type Pipeline struct {
	Cache     bool           `json:"cache"`
	CacheAs   string         `json:"cache_as"`
	OnFailure enum.OnFailure `json:"on_failure"`
	Delay     Delay          `json:"delay"`
}

//TestCase struct
type TestCase struct {
	Order       uint     `json:"order"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Request     Request  `json:"request"`
	Expect      Expect   `json:"expect"`
	Pipeline    Pipeline `json:"pipeline"`
}
