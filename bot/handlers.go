package bot

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"tg-alerter/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleUpdate(update tgbotapi.Update) {
	command, payload := getCommand(strings.ToLower(update.Message.Text))
	if command == "" {
		return
	}
	logger.GetLogger().Infoln("Command recieved:", command)

	switch command {
	case COMMAND_ALL_PREFIX:
		b.handleAllCommand(payload, update)
	}
}

func getCommand(input string) (string, string) {
	re := regexp.MustCompile(`^[/@](\w+)\s*(.*)`)
	match := re.FindStringSubmatch(input)
	if len(match) > 2 {
		return match[1], strings.TrimSpace(match[2])
	}
	return "", ""
}

func (b *Bot) handleAllCommand(userText string, update tgbotapi.Update) {
	log := logger.GetLogger()

	admins, err := b.API.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: update.Message.Chat.ID}})
	if err != nil {
		log.Printf("Error getting list of admins: %v", err)
		return
	}
	userName := update.Message.From.UserName
	log.Info("User that trigered command: ", userName)

	var mentions []string
	for _, admin := range admins {
		if admin.User.UserName != "" && admin.User.UserName != b.API.Self.UserName && admin.User.UserName != userName {
			mentions = append(mentions, fmt.Sprintf("@%s", admin.User.UserName))
		}
	}

	var messageToSend string = userText
	if messageToSend == "" {
		messageToSend = ListOfMessages[rand.Intn(len(ListOfMessages))]
	}

	mentionText := fmt.Sprintf("@%s сообщает:\n\n%s",
		update.Message.From.UserName,
		strings.TrimSpace(strings.Join(
			[]string{strings.Join(mentions, " "), messageToSend},
			" ",
		),
		),
	)

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           update.Message.Chat.ID,
			ReplyToMessageID: update.Message.MessageID,
		},
		Text:      mentionText,
		ParseMode: "Markdown",
	}

	if _, err := b.API.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
