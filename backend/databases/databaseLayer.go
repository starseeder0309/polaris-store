package databaseLayer

import (
	"github.com/kimhyeonu/polaris-store/backend/models"
)

type DatabaseLayer interface {
	// ID로 특정 회원 조회
	ReadMemberById(int) (models.Member, error)
	// 이름으로 특정 회원 조회
	ReadMemberByName(string, string) (models.Member, error)
	// 회원 생성
	CreateMember(models.Member) (models.Member, error)
	// 회원 접속
	ConnectMember(email, password string) (models.Member, error)
	// 회원 접속 해제
	DisconnectMember(int) error

	// 상품 목록 조회
	ReadProducts() ([]models.Product, error)
	// 프로모션 상품 목록 조회
	ReadPromotedProducts() ([]models.Product, error)
	// ID로 특정 상품 조회
	ReadProductById(int) (models.Product, error)

	// 회원 ID로 주문 목록 조회
	ReadOrdersById(int) ([]models.Order, error)
}
