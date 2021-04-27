package data

type VendorValue interface {
}

type Vendor struct {
	Code  int8
	Value VendorValue
}
