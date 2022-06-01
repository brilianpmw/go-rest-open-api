package repository

import (
	"brilianpmw/go-rest-open-api/helper"
	"brilianpmw/go-rest-open-api/model/domain"
	"errors"

	"context"
	"database/sql"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(name) values (?)"
	res, err := tx.ExecContext(ctx, SQL, product.Name)
	helper.PanicIfError(err)
	id, err := res.LastInsertId()
	helper.PanicIfError(err)
	product.Id = int(id)
	return product
}
func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "update product set name=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete from product where id=?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)

}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "select * from product where id=?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	product := domain.Product{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)
		return product, nil

	} else {
		return product, errors.New("Product Not found")
	}

}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "select * from product "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)

		products = append(products, product)

	}

	return products
}
