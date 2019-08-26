package datamodels

type Order struct {
	ID          int64 `sql:"ID" imooc:"ID"`
	UserId      int64 `sql:"userID" imooc:"userID"`
	ProductId   int64 `sql:"productID" imooc:"productID"`
	OrderStatus int64 `sql:"orderStatus" imooc:"orderStatus"`
}

const (
	OrderWait    = iota
	OrderSuccess  //1
	OrderFailed   //2
)
