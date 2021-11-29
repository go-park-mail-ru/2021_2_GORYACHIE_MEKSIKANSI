package orm

const (
	UnlimitedCount = -1
)

func ConvertInt32ToInt(i *int32) int {
	if i != nil {
		return int(*i)
	}
	return -1
}
