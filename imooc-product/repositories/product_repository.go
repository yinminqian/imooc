package repositories

import (
	"database/sql"
	"imooc-product/common"
	"imooc-product/datamodels"
	"strconv"
)

//第一步,先开发对应的接口
//第二步,实现定义的接口

//开发接口
type IProduct interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

//结构体要通过构造函数显示的实现上面的接口
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

//构造函数
func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table: table, mysqlConn: db}
}

//以下逐步实现结构体的接口函数

//数据库链接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return nil
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return nil
}

//增
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	//判断数据库已经链接
	if err = p.Conn(); err != nil {
		return
	}

	//写sql语句
	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	//对sql语句进行转换,sql包工具
	stmt, errSql := p.mysqlConn.Prepare(sql)
	if errSql != nil {
		return 0, errSql
	}

	//进行插入动作
	result, errStmt := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if errStmt != nil {
		return 0, errStmt
	}

	return result.LastInsertId()
}

//删
func (p *ProductManager) Delete(productID int64) bool {
	//判断sql链接
	if err := p.Conn(); err != nil {
		return false
	}
	//写删除sql语句

	sql := "DELETE FROM product WHERE ID=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	//执行被转换后的sql语句,传入占位符
	_, err = stmt.Exec(productID)
	if err != nil {
		return false
	}
	return true
}￿

//改
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	//依旧判断链接是否存在
	if err = p.Conn(); err != nil {
		return err
	}
	//写sql语句
	sql := "UPdate product set productName=?,productNum=?,productImage=?,productUrl=? where ID=" +
		strconv.FormatInt(product.ID, 10)
	//转换sql语句
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	//执行转换后的sql语句
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return err
	}
	return nil
}


//查 单条
func (p *ProductManager) SelectByKey(productID int64) (productResult *datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}
	sql := "Select from " + p.table + "where ID=" + strconv.FormatInt(productID, 10)

	row, errStmt := p.mysqlConn.Query(sql)
	if errStmt != nil {
		return &datamodels.Product{}, errStmt
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}

	common.DataToStructByTagSql(result, productResult)
	return
}
