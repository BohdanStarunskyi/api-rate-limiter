package cronManager

import (
	"api_tracker/config"
	"api_tracker/models"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type CronManager struct {
	scheduler *cron.Cron
}

func New() *CronManager {
	c := cron.New(cron.WithSeconds())
	return &CronManager{scheduler: c}
}

func (cm *CronManager) Start() {
	_, err := cm.scheduler.AddFunc("0 * * * * *", func() {
		if err := CleanupOldRequestLogs(); err != nil {
			log.Printf("Error cleaning up old requests: %v", err)
		} else {
			log.Println("Successfully cleaned up old requests.")
		}
	})

	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}

	cm.scheduler.Start()
}

func (cm *CronManager) Stop() {
	cm.scheduler.Stop()
}

func CleanupOldRequestLogs() error {
	expirationTime := time.Now().Add(-1 * time.Hour)

	if err := config.DB.Where("timestamp < ?", expirationTime).Delete(&models.RequestLog{}).Error; err != nil {
		return err
	}

	return nil
}
