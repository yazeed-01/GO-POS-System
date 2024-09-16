package routes

import (
	"POSSystem/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")

	// this is the urls:
	/*
	   1. " /products-view " -> for inventory  management
	   2. " /sell " -> for sell product
	   3. " /returnSellProduct " -> for return product
	   4. " /ProductsCashier " -> for cashier to see the products on inventory
	   5. " /PurchaseHistory " -> purchase history
	   6. " /dashboard-data " ->  admin dashboard
	*/

	//----------------------------------------------------------------------------
	r.GET("/products-view", func(c *gin.Context) {
		c.File("./static/inventoryProducts.html")
	})
	r.GET("/products", controllers.GetProducts)                  // read all
	r.POST("/products", controllers.CreateProduct)               // create
	r.PATCH("/products/:product_id", controllers.UpdateProduct)  // update
	r.DELETE("/products/:product_id", controllers.DeleteProduct) // delete

	// -------------------------------------------------------------------------

	r.GET("/sell", func(c *gin.Context) {
		c.File("./static/sellProduct.html")
	})

	r.GET("/sell-product-table/:id", controllers.SellProductTable)

	r.POST("/invoice", controllers.CreateInvoice)

	// -------------------------------------------------------------------------

	r.GET("/returnSellProduct/:id", controllers.ReturnProduct)

	r.GET("/returnSellProduct", func(c *gin.Context) {
		c.File("./static/returnSellProduct.html")
	})

	r.POST("/returnPayment/:id", controllers.ReturnPayment)

	// -------------------------------------------------------------------------

	r.GET("/ProductsCashier", func(c *gin.Context) {
		c.File("./static/ProductsCashier.html")
	})

	// -------------------------------------------------------------------------

	r.GET("/PurchaseHistory", func(c *gin.Context) {
		c.File("./static/PurchaseHistory.html")
	})

	r.GET("/PurchaseHistoryAll", controllers.PurchaseHistoryAll) // read all

	// -------------------------------------------------------------------------
	r.GET("/dashboard-data", func(c *gin.Context) {
		c.File("./static/AdminDashboard.html")
	})
	r.GET("/admin/dashboard-data", controllers.AdminDashboardData)
	return r
}
