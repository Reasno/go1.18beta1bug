package bar

type Data[K comparable, V any] struct {
	ch chan int
}

type Dataset[K comparable, V any] struct{}

func (m *Dataset[K, V]) GetData() *Data[K, V] {
	c := make(chan int, 1)
	d := &Data[K, V]{ch: c}
	return d
}

var DefaultMetadataSet = Dataset[string, any]{}

func GetData() *Data[string, any] {
	return DefaultMetadataSet.GetData()
}
