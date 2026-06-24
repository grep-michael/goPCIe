package pcielookup

type LookupResult struct {
	Vendor PCIeEntitiy
	Device PCIeEntitiy
	Source string
}

type PCIeEntitiy struct {
	ID   string
	Name string
}
