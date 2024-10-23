package application

import (
	application_communication "first-project/src/application/communication/emailService"
	"first-project/src/repository"
	"time"

	"github.com/robfig/cron/v3"
)

type CronJob struct {
	userRepository *repository.UserRepository
	emailService   *application_communication.EmailService
}

func NewCronJob(userRepository *repository.UserRepository, emailService *application_communication.EmailService) *CronJob {
	return &CronJob{
		userRepository: userRepository,
		emailService:   emailService,
	}
}

func (cronJob *CronJob) RunCronJob() {
	jobScheduler := cron.New()

	jobScheduler.AddFunc("@daily", func() {
		oneWeekAgo := time.Now().AddDate(0, 0, -7)
		users := cronJob.userRepository.FindUnverifiedUsersBeforeDate(oneWeekAgo)

		for _, user := range users {
			data := struct {
				Username string
			}{
				Username: user.Name,
			}
			cronJob.emailService.SendEmail(user.Email, "Activate your account", "activateAccount/en.html", data)
		}
	})

	jobScheduler.Start()
}
