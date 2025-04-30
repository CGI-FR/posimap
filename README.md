# POSIMAP

Transform positional text into structured data — and back again.

*POSIMAP* is a blend of "POSItional MAPper".

## Features

- Parse fixed-width text files based on a declarative positional schema.
- Convert structured data back into fixed-width text files.
- Easy configuration via YAML.
- Minimal dependencies, fast and lightweight.
- Supports validation and field trimming (optional).
- Supports text encoding conversion (EBCDIC, Unicode, ...).
- Supports OCCURS and REDEFINES in schema definition.

## Why POSIMAP ?

While CSV and JSON handle separated or structured data easily, fixed-width files are still widely used in legacy systems, financial data exchanges, and large-scale batch processing. POSIMAP helps bridge the gap between positional text formats and structured modern data workflows.

## Usage

### Example fixed-width data file (`person.fixed-width`)

```text
JOHN    DOE     1234 ELM STREET          SPRINGFIELD, IL 62704
JANE    SMITH   56 MAPLE AVENUE          RIVERSIDE, CA 92501
```

### Example config (`schema.yaml`)

```yaml
schema:
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
$ posimap fold < person.fixed-width
{"FIRSTNAME":"JOHN    ","LASTNAME":"DOE     ","ADDRESS":{"LINE-1":"1234 ELM STREET          ","LINE-2":"SPRINGFIELD, IL 62704"}}
{"FIRSTNAME":"JANE    ","LASTNAME":"SMITH   ","ADDRESS":{"LINE-1":"56 MAPLE AVENUE          ","LINE-2":"RIVERSIDE, CA 92501"}}
```

### Transform from JSON file to fixed-width format

```bash
$ posimap unfold < person.json
JOHN    DOE     1234 ELM STREET          SPRINGFIELD, IL 62704
JANE    SMITH   56 MAPLE AVENUE          RIVERSIDE, CA 92501
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

Copyright (C) 2025 CGI France

Posimap is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Posimap is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with posimap. If not, see http://www.gnu.org/licenses/.
