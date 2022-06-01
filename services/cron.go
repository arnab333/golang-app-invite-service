package services

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func CronInit() {
	c := cron.New()

	c.AddFunc("0 0 * * 0", cleanupInvitationToken) // This will run every 7 days at 00:00:00 hours

	// c.AddFunc("*/5 * * * *", cleanupInvitationToken) // This will run every 5 minutes

	c.Start()

	for {
		select {}
	}

}

func cleanupInvitationToken() {
	fmt.Println("Every 7 days")

	today := time.Now()

	for _, v := range InvitationTokens {
		if today.After(v.ExpiresAt) {
			DisableInviteToken(v.Token)
		}
	}
}
