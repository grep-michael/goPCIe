package pcielookup

type LookupResult struct {
	Vendor PCIeEntitiy
	Device PCIeEntitiy
	Class  string
	Source string
}

type PCIeEntitiy struct {
	ID   string
	Name string
}
