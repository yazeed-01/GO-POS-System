# GO Point of Sale (POS) System

This is a full-stack Point of Sale (POS) system developed using Go with the Gin framework for the backend and PostgreSQL as the database. The front end is built using HTML and Bootstrap. The system is designed to be used by any business that sells products and needs to manage inventory, handle sales transactions, return products, and analyze sales data through an admin dashboard. It provides seamless management of products, purchases, and performance tracking for businesses of all sizes.


## Features
1. **Sell Products**: Facilitate sales transactions and update inventory accordingly.
2. **Return Products**: Process product returns and adjust inventory.
3. **Inventory Management**: Full CRUD operations for products (add, update, delete, view).
4. **Purchase History**: Track and view the history of all purchases made in the system.
5. **Admin Dashboard**: Provides key business metrics including:
   - Total Revenue
   - Total Loss
   - Total Profit
   - Most Sold Product


## Routes
Here are the main URLs for the system:

1. `/products-view` - Manage and view inventory
2. `/sell` - Sell products
3. `/returnSellProduct` - Return products
4. `/ProductsCashier` - Cashier view of products in inventory
5. `/PurchaseHistory` - View purchase history
6. `/dashboard-data` - Admin dashboard to view business metrics


## To install the necessary dependencies for this project, run the following commands:

```bash
- go get github.com/githubnemo/CompileDaemon
- go install github.com/githubnemo/CompileDaemon
- go get github.com/joho/godotenv
- go get -u github.com/gin-gonic/gin
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/postgres
- go get github.com/jackc/pgx/v5
- go get -u golang.org/x/crypto/bcrypt
```
and for run use:
```bash
CompileDaemon -command="./POSSystem"
```
- `/PurchaseHistory` - View the purchase history.
- `/dashboard-data` - Admin dashboard displaying key metrics.




