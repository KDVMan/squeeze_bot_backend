package services_interface_dump

type DumpService interface {
	Dump(interface{})
	DumpLabel(string, interface{})
}
