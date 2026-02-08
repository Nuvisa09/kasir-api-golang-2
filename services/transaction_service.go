package services

import (
	"errors"
	"kasir-api-golang-2/models"
	"kasir-api-golang-2/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService{
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	if len(items) == 0 {
		return nil, errors.New("checkout items cannot be empty")
	}

	for _, item := range items {
		if item.Quantity <= 0 {
			return nil, errors.New("quantity must be greater than 0")
		}
	}
	return s.repo.CreateTransaction(items)
}