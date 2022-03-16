package controllers

import (
	"assignment_2/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetOrders(ctx *gin.Context) {
	var results = []models.Order{}

	sqlStatementOrder := `SELECT * FROM orders`
	rowsOrder, err := idb.DB.Query(sqlStatementOrder)

	if err != nil {
		panic(err.Error())
	}

	defer rowsOrder.Close()
	for rowsOrder.Next() {
		order := models.Order{}

		err = rowsOrder.Scan(
			&order.OrderID,
			&order.CustomerName,
			&order.OrderedAt,
		)

		if err != nil {
			panic(err.Error())
		}
		results = append(results, order)
	}

	for i, order := range results {
		items := []models.Item{}
		orderID := strconv.Itoa(order.OrderID)
		sqlStatementItem := `SELECT * FROM items WHERE order_id=$1`
		rowsItem, err := idb.DB.Query(sqlStatementItem, orderID)

		if err != nil {
			panic(err.Error())
		}

		defer rowsItem.Close()
		for rowsItem.Next() {
			item := models.Item{}

			err = rowsItem.Scan(
				&item.ItemID,
				&item.ItemCode,
				&item.Description,
				&item.Quantity,
				&item.OrderId,
			)
			if err != nil {
				panic(err.Error())
			}

			if order.OrderID == item.OrderId {
				items = append(items, item)
			}
		}
		results[i].Items = items
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})

}

func (idb *InDB) CreateOrder(ctx *gin.Context) {
	var (
		newOrder models.Order
	)

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
	INSERT INTO orders (
		customer_name, 
		ordered_at
	) VALUES ($1,$2)
	RETURNING order_id;`

	lastOrderID := 0
	err := idb.DB.QueryRow(sqlStatement, newOrder.CustomerName, newOrder.OrderedAt).Scan(&lastOrderID)
	if err != nil {
		panic(err)
	}

	sqlStatement = `
		INSERT INTO items (
			item_code, 
			description,
			quantity,
			order_id
		) VALUES ($1,$2,$3,$4);`

	for _, item := range newOrder.Items {
		res, err := idb.DB.Exec(sqlStatement,
			item.ItemCode,
			item.Description,
			item.Quantity,
			lastOrderID,
		)
		if err != nil {
			panic(err)
		}

		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Updated data amount : %+v\n", count)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": newOrder,
	})

}

func (idb *InDB) UpdateOrder(ctx *gin.Context) {

	var (
		order       models.Order
		updateOrder models.Order
	)

	orderID := ctx.Param("orderID")
	ID, err := strconv.Atoi(orderID)

	if err != nil {
		panic(err)
	}

	if err := ctx.ShouldBindJSON(&updateOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatementOrder := `SELECT * FROM orders WHERE order_id=$1`
	rowsOrder, err := idb.DB.Query(sqlStatementOrder, ID)

	if err != nil {
		panic(err)
	}

	defer rowsOrder.Close()
	for rowsOrder.Next() {

		err = rowsOrder.Scan(
			&order.OrderID,
			&order.CustomerName,
			&order.OrderedAt,
		)

		if err != nil {
			panic(err.Error())
		}
	}

	//Untuk update data orders berdasarkan ID dari parameter
	sqlStatement := `UPDATE 
		orders 
	SET 
		customer_name= $2, 
		ordered_at= $3
	WHERE 
		order_id = $1;`

	res, err := idb.DB.Exec(sqlStatement, order.OrderID, updateOrder.CustomerName, updateOrder.OrderedAt)
	if err != nil {
		panic(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	//Untuk menghapus data di table items berdasarkan order.ID=items.order_id
	sqlStatementItems := `DELETE FROM items WHERE order_id=$1`

	_, err = idb.DB.Exec(sqlStatementItems, ID)
	if err != nil {
		panic(err)
	}

	sqlStatement = `
		INSERT INTO items (
			item_code, 
			description,
			quantity,
			order_id
		) VALUES ($1,$2,$3,$4);`

	for _, item := range updateOrder.Items {
		res, err := idb.DB.Exec(sqlStatement,
			item.ItemCode,
			item.Description,
			item.Quantity,
			ID,
		)
		if err != nil {
			fmt.Println("ERRORRR =>>>>", err.Error())
			panic(err)
		}

		_, err = res.RowsAffected()
		if err != nil {
			panic(err)
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Data update successfully",
	})
}

func (idb *InDB) DeleteOrder(ctx *gin.Context) {

	orderID := ctx.Param("orderID")
	ID, err := strconv.Atoi(orderID)
	if err != nil {
		panic(err)
	}

	sqlStatementItems := `DELETE FROM items WHERE order_id=$1`

	_, err = idb.DB.Exec(sqlStatementItems, ID)
	if err != nil {
		panic(err)
	}

	sqlStatementOrders := `DELETE FROM orders WHERE order_id=$1`

	_, err = idb.DB.Exec(sqlStatementOrders, ID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Deleted successfully",
	})
}
