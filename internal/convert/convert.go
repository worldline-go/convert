package convert

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"

	"github.com/worldline-go/convert/internal/config"
	"github.com/worldline-go/convert/internal/template"
)

func Convert(cfg *config.Config) error {
	parse, err := NewParse(cfg)
	if err != nil {
		return err
	}

	log.Info().Msgf("reading file %q", cfg.Input)

	f, err := excelize.OpenFile(cfg.Input)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer f.Close()

	log.Info().Msgf("getting rows in sheet %q", cfg.Sheet)
	rows, err := f.GetRows(cfg.Sheet)
	if err != nil {
		return err
	}

	// /////////////////////////////////////////////////////////////
	// parse the rows and record the results
	if err := parse.Rows(rows); err != nil {
		return err
	}

	log.Info().Msg("parsing completed")
	parseResults := parse.GetResults()
	log.Debug().Msgf("results: %v", parseResults)

	// /////////////////////////////////////////////////////////////
	// run template and export the results
	for _, e := range cfg.Export {
		log.Info().Msgf("rendering template %q", e.Name)

		data, err := template.Render(e.Template, parseResults)
		if err != nil {
			return fmt.Errorf("failed to render template: %w", err)
		}

		if err := config.FileAPI.SetRaw(e.Output, data); err != nil {
			return fmt.Errorf("failed to save file: %w", err)
		}

		log.Info().Msgf("export saved to %q", e.Output)
	}

	return nil
}
