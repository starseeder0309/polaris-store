package services

import "github.com/kimhyeonu/polaris-store/backend/models"

type Product interface {
	// 상품 목록 조회
	ReadProducts() ([]models.Product, error)
	// 프로모션 상품 목록 조회
	ReadPromotedProducts() ([]models.Product, error)
	// ID로 특정 상품 조회
	ReadProductById(int) (models.Product, error)
}
