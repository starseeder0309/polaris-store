package services

import "github.com/kimhyeonu/polaris-store/backend/models"

type Member interface {
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
}
