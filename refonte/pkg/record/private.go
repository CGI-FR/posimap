package record

import (
	"github.com/cgi-fr/posimap/refonte/api"
	"github.com/cgi-fr/posimap/refonte/driven/document"
)

type Record interface {
	api.Record
	export(writer document.Writer, root Record) error
}
