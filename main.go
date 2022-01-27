package bar

type Data[K comparable, V any] struct {
	field int
}

type Dataset[K comparable, V any] struct{}

func (m *Dataset[K, V]) GetData() *Data[K, V] {
	d := &Data[K, V]{field: 1}
	return d
}

var DefaultDataset = Dataset[string, any]{}

func GetData() *Data[string, any] {
	return DefaultDataset.GetData()
}
