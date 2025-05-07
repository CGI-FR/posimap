package record

import (
	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
)

type Record interface {
	api.Record
	export(writer document.Writer, root Record) error
}
