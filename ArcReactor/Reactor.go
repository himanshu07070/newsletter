package arcreactor

import (
	"context"
	"os"
	"strings"
	"sync"

	"net/smtp"

	database "github.com/himanshu07070/newsletter/Database"
	input "github.com/himanshu07070/newsletter/Input"
	utils "github.com/himanshu07070/newsletter/Utils"
)

func ReactorEngine(dataChannel chan input.MetaData, wg *sync.WaitGroup) {
	ctx := context.Background()
	defer wg.Done()

loop:
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			utils.Logger.Error(err.Error())
			break loop

		case metaData := <-dataChannel:

			finalTemplate := strings.Replace(utils.Body, "{link}", metaData.Link, -1)

			from := os.Getenv("from")
			password := os.Getenv("password")
			if from == "" || password == "" {
				utils.Logger.Error("credentials are invalid", from, password)
				return
			}
			if !utils.ValidateEmailAddress(from) || !utils.IsEmailValid(from) {
				utils.Logger.Info("Invalid email Id")
				return
			}
			var allUsers []database.User
			if err := database.GetAllUserEmails(&allUsers, ctx); err != nil {
				utils.Logger.Error(err.Error())
				return
			}

			var allEmails []string
			for _, value := range allUsers {
				allEmails = append(allEmails, value.Email)
			}
			to := allEmails
			smtpHost := utils.SmtpHost
			smtpPort := utils.SmtpPort
			subject := metaData.Subject
			mime := utils.Mime
			message := []byte(subject + mime + finalTemplate)
			auth := smtp.PlainAuth("", from, password, smtpHost)

			err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
			if err != nil {
				utils.Logger.Error(err.Error())
				return
			}
			utils.Logger.Info("Newsletter sent to all emails")

		}
	}

}
