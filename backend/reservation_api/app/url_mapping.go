package app

import (
    "reservation_api/bd"
    "reservation_api/controller"
    "reservation_api/repository"
    "reservation_api/service"

    log "github.com/sirupsen/logrus"
)

func mapsUrls() {
    log.Info("Starting mapping configurations")

    // Repositorio con la DB ya inicializada
    reservationRepo := repository.NewReservationRepository(bd.DB)

    // Servicio
    reservationService := service.NewReservationService(reservationRepo)

    // Controller
    reservationController := controller.NewReservationController(reservationService)

    // Rutas con métodos del controller
    router.POST("/reservations", reservationController.Create)
    router.GET("/reservations/user/:id", reservationController.GetByUser)
    router.PUT("/reservations/:id/cancel", reservationController.Cancel)
}
