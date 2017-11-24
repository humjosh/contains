package contains

type StringSlice []string
type IntSlice []int
type Int32Slice []int32
type Int64Slice []int64

func (o StringSlice) Contains(search string) bool {
	for _, val := range o {
		if search == val {
			return true
		}
	}
	return false
}

func (o IntSlice) Contains(search int) bool {
	for _, val := range o {
		if search == val {
			return true
		}
	}
	return false
}

func (o Int32Slice) Contains(search int32) bool {
	for _, val := range o {
		if search == val {
			return true
		}
	}
	return false
}

func (o Int64Slice) Contains(search int64) bool {
	for _, val := range o {
		if search == val {
			return true
		}
	}
	return false
}
