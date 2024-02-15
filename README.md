# Convert

Helper repo to covert excel to custom file.

This tool uses [excelize](https://github.com/qax-os/excelize) and supports XLSM / XLSX / XLTM / XLTX files.

## Usage

Create a config file `config.yaml` and pass with the env variable `CONFIG_FILE`.

```yaml
input: /path/to/input.xlsx
sheet: Sheet1

map:
  custom_1: # convert to map of the column
    C: Sequence
    E: TD

parse:
  custom_value1:
    map: custom_1
    rows:
      - 31-32

export:
  - name: csv
    template: |
      {{ range $_, $element := .custom_value1 -}}
      {{ $element.Sequence }},{{ $element.TD }}
      {{ end }}
    output: /path/to/output.csv
```
