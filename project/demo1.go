package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryID"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)

	return products, nil
}

func AddProduct() (Product, error) {
	//product := Product{Id: 4, ProductName: "Telephone", CategoryId: 1, UnitPrice: 3000.0}

	////özel bir id vermezsen json id leri otomatik olarak sıralı verir
	product := Product{ProductName: "Mouse Pad", CategoryId: 1, UnitPrice: 25.0}

	jsonProduct, _ := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8",
		bytes.NewBuffer(jsonProduct))

	if err != nil {
		return Product{}, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productresponse Product

	json.Unmarshal(bodyBytes, &productresponse)
	return productresponse, nil
}

/*main içi

project.AddProduct()
	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
*/
