//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/orm WrapperPromocodeInterface,ConnectionInterface,TransactionInterface

package orm

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/myerror"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type WrapperPromocodeInterface interface {
	GetTypePromoCode(promoCode string, restaurantId int) (int, error)
	ActiveFreeDelivery(promoCode string, restaurantId int) (bool, error)
	ActiveCostForFreeDish(promoCode string, restaurantId int) (int, int, error)
	ActiveCostForSale(promoCode string, amount int, restaurantId int) (int, error)
	ActiveTimeForSale(promoCode string, amount int, restaurantId int) (int, error)
}

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type TransactionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Conn() *pgx.Conn
}

type Wrapper struct {
	Conn ConnectionInterface
}

func (db *Wrapper) GetTypePromoCode(promoCode string, restaurantId int) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PGetTypePromoCodeTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var typePromoCode int
	err = tx.QueryRow(contextTransaction,
		"SELECT type FROM promocode WHERE code = $1 AND restaurant = $2 AND end_date > NOW()",
		promoCode, restaurantId).Scan(&typePromoCode)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, &errPkg.Errors{
				Alias: errPkg.PGetTypePromoCodeRestaurantsNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.PGetTypePromoCodeRestaurantsNotSelect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PGetTypePromoCodeNotCommit,
		}
	}
	return typePromoCode, nil
}

func (db *Wrapper) ActiveFreeDelivery(promoCode string, restaurantId int) (bool, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.PActiveFreeDeliveryTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var freeDelivery bool
	err = tx.QueryRow(contextTransaction,
		"SELECT free_delivery FROM promocode WHERE code = $1 AND restaurant = $2",
		promoCode, restaurantId).Scan(&freeDelivery)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, &errPkg.Errors{
				Alias: errPkg.PActiveFreeDeliveryRestaurantsNotFound,
			}
		}
		return false, &errPkg.Errors{
			Alias: errPkg.PActiveFreeDeliveryRestaurantsNotSelect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.PActiveFreeDeliveryNotCommit,
		}
	}
	return freeDelivery, nil
}

func (db *Wrapper) ActiveCostForFreeDish(promoCode string, restaurantId int) (int, int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, 0, &errPkg.Errors{
			Alias: errPkg.PActiveCostForFreeDishTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var costForFreeDish int
	var dishId int
	err = tx.QueryRow(contextTransaction,
		"SELECT cost_for_free_dish, free_dish_id FROM promocode WHERE code = $1 AND restaurant = $2",
		promoCode, restaurantId).Scan(&costForFreeDish, &dishId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, 0, &errPkg.Errors{
				Alias: errPkg.PActiveCostForFreeDishRestaurantsNotFound,
			}
		}
		return 0, 0, &errPkg.Errors{
			Alias: errPkg.PActiveCostForFreeDishRestaurantsNotSelect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, 0, &errPkg.Errors{
			Alias: errPkg.PActiveCostForFreeDishNotCommit,
		}
	}
	return costForFreeDish, dishId, nil
}

func (db *Wrapper) ActiveCostForSale(promoCode string, amount int, restaurantId int) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PActiveCostForSaleTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var costForFreeDelivery int
	var salePercent, saleAmount *int32
	err = tx.QueryRow(contextTransaction,
		"SELECT cost_for_sale, sale_percent, sale_amount FROM promocode WHERE code = $1 AND restaurant = $2",
		promoCode, restaurantId).Scan(&costForFreeDelivery, &salePercent, &saleAmount)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, &errPkg.Errors{
				Alias: errPkg.PActiveCostForSaleRestaurantsNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.PActiveCostForSaleRestaurantsNotSelect,
		}
	}

	var newSum int
	if salePercent != nil {
		newSum = amount - amount*int(*salePercent)/100
	} else {
		newSum = amount - int(*saleAmount)
		if newSum < 0 {
			newSum = 0
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PActiveCostForSaleNotCommit,
		}
	}
	return newSum, nil
}

func (db *Wrapper) ActiveTimeForSale(promoCode string, amount int, restaurantId int) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PActiveTimeForSaleTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var timeSale time.Time
	var salePercent, saleAmount *int32
	err = tx.QueryRow(contextTransaction,
		"SELECT time_for_sale, sale_in_time_percent, sale_in_time_amount FROM promocode WHERE code = $1 AND restaurant = $2",
		promoCode, restaurantId).Scan(&timeSale, &salePercent, &saleAmount)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, &errPkg.Errors{
				Alias: errPkg.PActiveTimeForSaleRestaurantsNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.PActiveTimeForSaleRestaurantsNotSelect,
		}
	}

	var newSum int
	if time.Now().Before(timeSale) {
		if salePercent != nil {
			newSum = amount - amount*int(*salePercent)/100
		} else {
			newSum = amount - int(*saleAmount)
			if newSum < 0 {
				newSum = 0
			}
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PActiveTimeForSaleNotCommit,
		}
	}
	return newSum, nil
}
