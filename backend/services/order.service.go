package services

import "github.com/kimhyeonu/polaris-store/backend/models"

type Order interface {
	// 회원 ID로 주문 목록 조회
	ReadOrdersById(int) ([]models.Order, error)
}
