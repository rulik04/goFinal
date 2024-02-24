package homecontroller

import (
	"fmt"
	"go-web-native/entities"
	"go-web-native/models/productmodel"
	"net/http"
	"text/template"
) 
 
func Welcome(w http.ResponseWriter, r *http.Request) { 
 products := productmodel.Getall() 
 
 sortedProducts := make(map[entities.Category][]entities.Product) 
 for _, product := range products { 
  categoryName := product.Category 
  sortedProducts[categoryName] = append(sortedProducts[categoryName], product) 
 } 
 
 data := map[string]interface{}{ 
        "SortedProducts": sortedProducts, 
    } 
 
 fmt.Println(sortedProducts) 
 
 tmpl, err := template.ParseFiles("views/home/index.html") 
 if err != nil { 
  http.Error(w, err.Error(), http.StatusInternalServerError) 
  return 
 } 
 
 err = tmpl.Execute(w, data) 
 if err != nil { 
  http.Error(w, err.Error(), http.StatusInternalServerError) 
  return 
 } 
}