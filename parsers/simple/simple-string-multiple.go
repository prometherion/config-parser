package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleStringMultiple struct {
	Enabled    bool
	Value      []string
	Name       string
	SearchName string
}

func (s *SimpleStringMultiple) Init() {
	s.Enabled = false
	s.SearchName = s.Name
}

func (s *SimpleStringMultiple) GetParserName() string {
	return s.SearchName
}

func (s *SimpleStringMultiple) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.SearchName) {
		elements := strings.SplitN(line, " ", 2)
		if len(elements) < 2 {
			return "", &errors.ParseError{Parser: "SimpleStringMultiple", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Value = elements[1:]
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *SimpleStringMultiple) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *SimpleStringMultiple) String() []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %s", s.SearchName, strings.Join(s.Value, " "))}
	}
	return []string{}
}
