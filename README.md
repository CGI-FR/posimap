# POSCH

Transform positional text into structured data — and back again.

*POSCH* is a blend of "POsitional" and "SCHema".

## Features

- Parse fixed-width text files based on a declarative positional schema.
- Convert structured data back into fixed-width text files.
- Easy configuration via YAML.
- Minimal dependencies, fast and lightweight.
- Supports validation and field trimming (optional).
- Supports encoding conversion (EBCDIC, Unicode, ...)

## Why ?

While CSV and JSON handle separated or structured data easily, fixed-width files are still widely used in legacy systems, financial data exchanges, and large-scale batch processing. POSCH helps bridge the gap between positional text formats and structured modern data workflows.

## Usage

### Example schema (`schema.yaml`)

```yaml
- name: FIRSTNAME
  length: 8
- name: LASTNAME
  length: 8
- name: ADDRESS
  schema:
    - name: LINE-1
      length: 25
    - name: LINE-2
      length: 25
```

### Transform from fixed-width file to JSON format

```bash
$ posch fold < person.fixed-width
{ "FIRSTNAME": "JOHN    ", "LASTNAME": "DOE     ", "ADDRESS": { "LINE-1": "1234 ELM STREET          ", "LINE-2": "SPRINGFIELD, IL 62704    " } }
```

### Transform from JSON file to fixed-width format

```bash
$ posch unfold < person.json
JOHN    DOE     1234 ELM STREET          SPRINGFIELD, IL 62704
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

Copyright (C) 2025 CGI France

posch is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

posch is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with posch. If not, see http://www.gnu.org/licenses/.
