package convert

import (
	"fmt"

	"convert/internal/config"

	"github.com/rs/zerolog/log"
)

type Parse struct {
	Results map[string]*Result
}

type Result struct {
	Rows *Rows
	Map  *Map

	Values []map[string]string
}

func (p *Parse) GetResults() map[string][]map[string]string {
	results := make(map[string][]map[string]string)
	for name, result := range p.Results {
		results[name] = result.Values
	}

	return results
}

func (p *Parse) Rows(rows [][]string) error {
	for i, row := range rows {
		for name, result := range p.Results {
			if result.Rows.IsInclude(i + 1) {
				log.Debug().Msgf("parsing row %d to %q", i, name)
				parsedValues := result.Map.Parse(row)
				log.Debug().Msgf("parsed values: %v", parsedValues)
				result.AddValue(parsedValues)
			}
		}
	}

	return nil
}

func (r *Result) AddValue(values map[string]string) {
	r.Values = append(r.Values, values)
}

func NewParse(cfg *config.Config) (*Parse, error) {
	parse := &Parse{
		Results: make(map[string]*Result),
	}

	maps := make(map[string]*Map)
	for name, m := range cfg.Map {
		mp, err := NewMap(m)
		if err != nil {
			return nil, fmt.Errorf("failed to create %q map: %w", name, err)
		}

		maps[name] = mp
	}

	for name, p := range cfg.Parse {
		rows, err := NewRows(p.Rows)
		if err != nil {
			return nil, fmt.Errorf("failed to create %q rows: %w", name, err)
		}

		mp, ok := maps[p.Map]
		if !ok {
			return nil, fmt.Errorf("map %q not found on %q", p.Map, name)
		}

		parse.Results[name] = &Result{
			Rows: rows,
			Map:  mp,
		}
	}

	return parse, nil
}
