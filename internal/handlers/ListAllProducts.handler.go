package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/product"
)

func ListAllProducts(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productListParams := &stripe.ProductListParams{}
		it := product.List(productListParams)
		// marshal to json as array?

		products := it.ProductList()

		w.Header().Set("content-type", "application/json")
		if err := json.NewEncoder(w).Encode(products); err != nil {
			log.Printf("[ERROR] (%s %s) could not send products response: %s\n", r.Method, r.URL, err)
			http.Error(w, "Could not get products list", http.StatusInternalServerError)
		}
	}
}
