package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kimhyeonu/polaris-store/backend/controllers"
)

func RunApi(address string) error {
	router := gin.Default()

	controller, _ := controllers.GenerateController()

	membersGroup := router.Group("/members")
	{
		// 회원 생성
		membersGroup.POST("", controller.CreateMember)
		// 회원 접속
		membersGroup.POST("/connection", controller.ConnectMember)
		// 회원 접속 해제
		membersGroup.POST("/:id/disconnection", controller.DisconnectMember)

		// 회원 ID로 주문 목록 조회
		membersGroup.GET("/:id/orders", controller.ReadOrdersById)
		// 결제 처리
		membersGroup.POST("/payment", controller.ProcessPayment)
	}

	// 상품 목록 조회
	router.GET("/products", controller.ReadProducts)
	// 프로모션 상품 목록 조회
	router.GET("/products/promotion", controller.ReadPromotedProducts)

	return router.Run(address)
}
