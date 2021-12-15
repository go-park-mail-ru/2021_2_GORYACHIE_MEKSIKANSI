package orm

const (
	UnlimitedCount        = -1
	PromoCodeFreeDelivery = 1
	PromoCodeSaleOverCost = 2
	PromoCodeSaleOverTime = 3
	PromoCodeFreeDishes   = 4
)

func ConvertInt32ToInt(i *int32) int {
	if i != nil {
		return int(*i)
	}
	return -1
}
