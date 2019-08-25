package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"imooc-product/common"
	"imooc-product/datamodels"
	"imooc-product/services"
	"strconv"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
}

//在controller中调用service中的方法
func (p *ProductController) GetAll() mvc.View {
	productAll, _ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": productAll,
		},
	}
}

//修改商品
func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

//添加商品的方法
func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProductService.InsertProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetManager() mvc.View {
	//首先获取id->转换类型
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	fmt.Print("id==>", id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "product/manager.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

//删除商品
func (p *ProductController) GetDelete() {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	isOk := p.ProductService.DeleteProductByID(id)
	if isOk {
		p.Ctx.Application().Logger().Debug("删除成功,ID为" + idString)
	} else {
		p.Ctx.Application().Logger().Debug("删除失败,ID为" + idString)
	}
	p.Ctx.Redirect("/product/all")
}
