package main

import (
	"fmt"
	"net/smtp"
	"sms-systems-monitoring/alerts"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/cpu"
)

// Configs
const (
	ComputerMonitoring = "Kevin's Mac Book Pro" // Device This will Monitor
	MemoryThreshold    = 70.0                   // Memory % Alert Threshold
	MemCheckInterval   = 5 * time.Second        // Check Memory every 5s
	BatteryThreshold   = 30.0                   // Batter % Alert Threshold
	BatCheckInterval   = 60 * time.Second       // Check Battery Every Min
)

func main() {
	log.Info().Msg("Now Monitoring CPU Usage and Battery Level")

	go monitorCPU()
	go monitorBattery()

}

func monitorCPU() {
	for {
		percentages, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Warn().Err(err).Msg("Error getting CPU usage:")
			continue
		}

		cpuUsage := percentages[0]
		log.Printf("CPU usage: %f", cpuUsage)

		if cpuUsage > MemoryThreshold {

		}
	}

}

func monitorBattery() {

}

func sendCpuAlert(cpuUsage float64) {
	subject := fmt.Sprintf("⚠️ CPU Usage Alert! ⚠️ Usage: %.2f", cpuUsage)
	body := fmt.Sprintf("The CPU usage on %s has exceeded %.2f. Current Usage: %.2f", ComputerMonitoring, MemoryThreshold, cpuUsage)
	message := fmt.Sprintf("Subject: %s\r\n\n\n%s", subject, body)

	smtpConfig, err := alerts.ConfigSMTP()
	if err != nil {
		log.Warn().Err(err).Msg("error getting the SMTP configurations")
	}

	auth := smtp.PlainAuth("", smtpConfig.SmtpUser, smtpConfig.SmtpPassword, smtpConfig.SmtpServer)
	addr := fmt.Sprintf("%s:%s", smtpConfig.SmtpServer, smtpConfig.SmtpPort)

	err = smtp.SendMail(addr, auth, smtpConfig.EmailFrom, []string{smtpConfig.EmailTo}, []byte(message))
	if err != nil {
		log.Warn().Err(err).Msg("Failed to send alert")
		return
	}

	log.Info().Msg("CPU Alert Successfully Sent!")
}
