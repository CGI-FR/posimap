package record

import "github.com/cgi-fr/posimap/refonte/api"

type Record interface {
	api.Record
	export(writer api.StructWriter, root Record) error
}
