package database

import (
	"context"
	"database/sql"
	"log"
	"github.com/ikiselewskii/avito-test-task/models"
	"github.com/ikiselewskii/avito-test-task/utils"
)

func CreateTables(ctx context.Context) {
	err := createUsersTable(ctx)
	if err != nil {
		panic(err)
	}
	err = createReservationsTable(ctx)
	if err != nil {
		panic(err)
	}

}

func createUsersTable(ctx context.Context) error {
	_, err := DB.NewCreateTable().
		Model((*models.Customer)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Println("Failed to create Users Table ", err)
	}
	return err
}

func createReservationsTable(ctx context.Context) error {
	_, err := DB.NewCreateTable().
		Model((*models.Transaction)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Println("Failed to create ReservationsTable ", err)
	}
	return err
}




func AddMoney(to models.Customer, ctx context.Context) error {
	_, err := DB.NewInsert().
		Model(&to).
		On("CONFLICT (id) DO UPDATE").
		Set("balance = customer.balance + EXCLUDED.balance").
		Exec(ctx)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetBalance(usr models.Customer, ctx context.Context) (int, error){
	err := DB.NewSelect().
	Model(&usr).
	ColumnExpr("balance").
	Where("id = ?", usr.ID).
	Scan(ctx, &usr)
	if err != nil {
		log.Println(err, " no such user")
		return 0, err
	}
	return usr.Balance, err
}

func Reserve(tr models.Transaction, ctx context.Context) error {
	tx, err := DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println("Failed to begin transaction ", err)
		tx.Rollback()
		return err
	}
	var customers []models.Customer
	err = tx.NewSelect().
		Model(&customers).
		ColumnExpr("balance").
		Where("id = ?", tr.FromID).
		For("UPDATE").
		Scan(ctx, &customers)
	if len(customers) == 0 {
		log.Println("user doesn`t exist")
		tx.Rollback()
		return sql.ErrNoRows
	}
	customer := customers[0]
	log.Println(customer)
	if customer.Balance < tr.Amount {
		log.Printf("inssuficient balance on account id:%d", customer.ID)
		tx.Rollback()
		return &utils.InsufficientBalanceError{}
	}
	if err != nil {
		tx.Rollback()
		log.Println("Something went wrong ", err)
		return err
	}
	_, err = tx.NewUpdate().
		Model(&customer).
		Set("balance = ?", customer.Balance-tr.Amount).
		Where("id = ?", tr.FromID).
		Exec(ctx)

	if err != nil {
		log.Println("Failed to reduce balance ", err)
		tx.Rollback()
		return err
	}
	tr.Status = "reserved"
	_, err = tx.NewInsert().
		Model(&tr).
		Exec(ctx)
	if err != nil {
		log.Println("Can not reserve ", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func Approve(tr models.Transaction, ctx context.Context) error{
	tx, err := DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println("Failed to begin transaction ", err)
		tx.Rollback()
		return err
	}
	
	_, err = tx.NewUpdate().
				Model(&tr).
					Where("order_id = ?", tr.OrderID).
						Set("status = ?", "accepted").Exec(ctx)
	if err != nil{
		log.Println("unable to approve transaction due to ", err)
		return err
	}
	return err
}
