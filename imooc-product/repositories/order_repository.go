package repositories

import (
	"database/sql"
	"imooc-product/common"
	"imooc-product/datamodels"
	"strconv"
)

type IOrderRepository interface {
	Conn() error
	Insert(order *datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(order *datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderMangerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func NewOrderMAngerRepository(table string, sql *sql.DB) IOrderRepository {
	return &OrderMangerRepository{table: table, mysqlConn: sql}
}

//链接数据库
func (o *OrderMangerRepository) Conn() error {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return nil
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

//订单的插入
func (o *OrderMangerRepository) Insert(order *datamodels.Order) (productID int64, err error) {
	if err = o.Conn(); err != nil {
		return 0, err
	}
	sql := "INSERT " + o.table + " set userID=?,productID=?,orderStatus=?"
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return productID, errStmt
	}

	result, errResult := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if errResult != nil {
		return productID, errResult
	}
	productID, err = result.LastInsertId()
	return
}

//订单的删除
func (o *OrderMangerRepository) Delete(productID int64) (isOK bool) {
	if err := o.Conn; err != nil {
		return
	}
	sql := "delete from " + o.table + " where ID =?"
	stmt, errStmt := o.mysqlConn.Prepare(sql)

	if errStmt != nil {
		return
	}
	_, errResult := stmt.Exec(productID)
	if errResult != nil {
		return
	}
	return true
}

//订单的更新
func (o *OrderMangerRepository) Update(order *datamodels.Order) (err error) {
	if err = o.Conn(); err != nil {
		return err
	}
	sql := "Update " + o.table + " set userID=?,productID=?,orderStatus=? where ID=" +
		strconv.FormatInt(order.ID, 10)
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return errStmt
	}

	_, errResult := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)

	if errResult != nil {
		return errResult
	}
	return nil
}

//订单的查询 单个查询
func (o *OrderMangerRepository) SelectByKey(orderID int64) (order *datamodels.Order, err error) {
	if errConn := o.Conn; errConn != nil {
		return &datamodels.Order{}, errConn()
	}
	sql := "select * from " + o.table + "where ID=" + strconv.FormatInt(orderID, 10)
	row, errRow := o.mysqlConn.Query(sql)
	if errRow != nil {
		return &datamodels.Order{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Order{}, err
	}

	order = &datamodels.Order{}
	common.DataToStructByTagSql(result, order)
	return
}

//查询所有

func (o *OrderMangerRepository) SelectAll() (orderArray []*datamodels.Order, err error) {
	if errConn := o.Conn(); errConn != nil {
		return nil, errConn
	}
	sql := "select * from " + o.table
	rows, errRows := o.mysqlConn.Query(sql)
	if errRows != nil {
		return nil, errRows
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, err
	}
	for _, v := range result {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orderArray = append(orderArray, order)
	}
	return
}

func (o *OrderMangerRepository) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	if errConn := o.Conn(); errConn != nil {
		return nil, errConn
	}

	sql := "Select o.ID,p.productName,o.OrderStatus From imooc.order as o left join product as p on o.productID=p.ID"

	rows, errRows := o.mysqlConn.Query(sql)
	if errRows != nil {
		return nil, errRows
	}
	orderMap = common.GetResultRows(rows)
	return
}
