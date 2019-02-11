package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type DefaultBackend struct {
	data *types.StringC
}

func (s *DefaultBackend) Init() {
	s.data = nil
}

func (s *DefaultBackend) GetParserName() string {
	return "default_backend"
}

func (s *DefaultBackend) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.StringC{}
			return s.data, nil
		}
		return nil, errors.FetchError
	}
	return s.data, nil
}

func (p *DefaultBackend) GetOne(index int) (common.ParserData, error) {
	if index != 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (s *DefaultBackend) Set(data common.ParserData, index int) error {
	if data == nil {
		s.data = nil
		return nil
	}
	switch newValue := data.(type) {
	case *types.StringC:
		s.data = newValue
	case types.StringC:
		s.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (s *DefaultBackend) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "default_backend" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "DefaultBackend", Line: line, Message: "Parse error"}
		}
		s.data = &types.StringC{
			Comment: comment,
			Value:   parts[1],
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "default_backend", Line: line}
}

func (s *DefaultBackend) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("default_backend %s", s.data.Value),
			Comment: s.data.Comment,
		},
	}, nil
}