package template

import (
	"bytes"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/rytsh/mugo/pkg/fstore"
	"github.com/rytsh/mugo/pkg/templatex"
	"github.com/worldline-go/logz"
)

var Global = templatex.New(templatex.WithAddFuncsTpl(
	fstore.FuncMapTpl(
		fstore.WithLog(logz.AdapterKV{Log: log.Logger}),
		fstore.WithTrust(true),
		fstore.WithWorkDir("."),
	),
))

func Render(template string, data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := Global.Execute(templatex.WithContent(template), templatex.WithIO(&buf), templatex.WithData(data)); err != nil {
		return nil, fmt.Errorf("failed to render template: %w", err)
	}

	return buf.Bytes(), nil
}
