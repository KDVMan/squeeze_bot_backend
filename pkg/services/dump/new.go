package services_dump

import (
	services_interface_dump "backend/pkg/services/dump/interface"
	"github.com/k0kubun/pp/v3"
)

type dumpServiceImplementation struct {
	pp *pp.PrettyPrinter
}

func NewDumpService() services_interface_dump.DumpService {
	prettyPrinter := pp.New()
	prettyPrinter.SetDecimalUint(true)

	return &dumpServiceImplementation{
		pp: prettyPrinter,
	}
}

func (object *dumpServiceImplementation) Dump(value interface{}) {
	object.pp.Println(value)
}

func (object *dumpServiceImplementation) DumpLabel(label string, value interface{}) {
	object.pp.Println(label, value)
}
