package services

import (
	"imooc-product/datamodels"
	"imooc-product/repositories"
)

//同 repositories一样 先定义接口
type IProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(product *datamodels.Product) (int64, error)
	UpdateProduct(product *datamodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

//生成构造函数
func NewProductService(repository repositories.IProduct) IProductService {
	return &ProductService{repository}
}

//在serve中调用repositories里面的方法,repositories写较底层的方法,供service调用
//查->直接调用repositories里面的方法
func (p *ProductService) GetProductByID(productID int64) (*datamodels.Product, error) {
	return p.productRepository.SelectByKey(productID)
}

//查->全部
func (p *ProductService) GetAllProduct() ([]*datamodels.Product, error) {
	return p.productRepository.SelectAll()
}

//删除
func (p *ProductService) DeleteProductByID(productID int64) bool {
	return p.productRepository.Delete(productID)
}

//插入一条
func (p *ProductService) InsertProduct(product *datamodels.Product) (productID int64, err error) {
	return p.productRepository.Insert(product)
}

//更新
func (p *ProductService) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}
