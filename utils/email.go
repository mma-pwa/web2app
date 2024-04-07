package utils

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

type MessageType string

type MailSender struct {
	ReceiveAddress []MailInfo
	Cc             []MailInfo
	Mails          []MailMessage
	MailConfig
}

const (
	TextType     MessageType = "text"
	HtmlType     MessageType = "html"
	LinkType     MessageType = "link"
	MarkdownType MessageType = "markdown"
)

type Message struct {
	Title   string
	Message string
	Type    MessageType
}
type MailInfo struct {
	Name    string
	Address string
}

type MailMessage struct {
	Message
	AttachesPath []string
}

func NewMailReceiver(name, address string) MailInfo {
	return MailInfo{
		Name:    name,
		Address: address,
	}
}

func NewMailMessage(title, message string, messageType MessageType, attaches []string) MailMessage {
	return MailMessage{
		Message: Message{
			Title:   title,
			Message: message,
			Type:    messageType,
		},
		AttachesPath: attaches,
	}
}

func NewMailConfig(host string, port int, username, password string, tls bool) MailConfig {
	return MailConfig{
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  password,
		TLSSecure: tls,
	}
}

func NewMailSenderDefault(mailConfig MailConfig, receivers []string, Ccs []string, mails []MailMessage) *MailSender {
	to := make([]MailInfo, 0, 1)
	for _, receiver := range receivers {
		to = append(to, NewMailReceiver("", receiver))
	}
	cc := make([]MailInfo, 0, 1)
	for _, s := range Ccs {
		cc = append(cc, NewMailReceiver("", s))
	}
	return &MailSender{
		ReceiveAddress: to,
		Cc:             cc,
		Mails:          mails,
		MailConfig:     mailConfig,
	}
}

func NewMailSender(mailConfig MailConfig, to []MailInfo, cc []MailInfo, mails []MailMessage) *MailSender {
	return &MailSender{
		ReceiveAddress: to,
		Cc:             cc,
		Mails:          mails,
		MailConfig:     mailConfig,
	}
}

func (m *MailSender) LoadConfig(config MailConfig) *MailSender {
	m.Host = config.Host
	m.Port = config.Port
	m.Username = config.Username
	m.Password = config.Password
	m.TLSSecure = config.TLSSecure
	return m
}

func (m *MailSender) AddReceiver(receiver string) *MailSender {
	mailReceivers := m.ReceiveAddress
	if mailReceivers == nil {
		mailReceivers = make([]MailInfo, 0, 1)
	}
	mailReceivers = append(mailReceivers, NewMailReceiver("", receiver))
	m.ReceiveAddress = mailReceivers
	return m
}

func (m *MailSender) AddReceivers(receivers []string) *MailSender {
	mailReceivers := m.ReceiveAddress
	if mailReceivers == nil {
		mailReceivers = make([]MailInfo, 0, 1)
	}
	for _, receiver := range receivers {
		mailReceivers = append(mailReceivers, NewMailReceiver("", receiver))
	}
	m.ReceiveAddress = mailReceivers
	return m
}

func (m *MailSender) AddMailReceiver(receiver MailInfo) *MailSender {
	mailReceivers := m.ReceiveAddress
	if mailReceivers == nil {
		mailReceivers = make([]MailInfo, 0, 1)
	}
	mailReceivers = append(mailReceivers, receiver)
	m.ReceiveAddress = mailReceivers
	return m
}

func (m *MailSender) AddMailReceivers(receivers []MailInfo) *MailSender {
	mailReceivers := m.ReceiveAddress
	if mailReceivers == nil {
		mailReceivers = make([]MailInfo, 0, 1)
	}
	mailReceivers = append(mailReceivers, receivers...)
	m.ReceiveAddress = mailReceivers
	return m
}

func (m *MailSender) AddCc(Cc string) *MailSender {
	cc := m.Cc
	if cc == nil {
		cc = make([]MailInfo, 0, 1)
	}
	cc = append(cc, NewMailReceiver("", Cc))
	m.Cc = cc
	return m
}

func (m *MailSender) AddCcs(Ccs []string) *MailSender {
	cc := m.Cc
	if cc == nil {
		cc = make([]MailInfo, 0, 1)
	}
	for _, s := range Ccs {
		cc = append(cc, NewMailReceiver("", s))
	}
	m.Cc = cc
	return m
}

func (m *MailSender) AddMailCc(Cc MailInfo) *MailSender {
	cc := m.Cc
	if cc == nil {
		cc = make([]MailInfo, 0, 1)
	}
	cc = append(cc, Cc)
	m.Cc = cc
	return m
}

func (m *MailSender) AddMailCcs(Ccs []MailInfo) *MailSender {
	cc := m.Cc
	if cc == nil {
		cc = make([]MailInfo, 0, 1)
	}
	cc = append(cc, Ccs...)
	m.Cc = cc
	return m
}

func (m *MailSender) AddMail(message MailMessage) *MailSender {
	mails := m.Mails
	if mails == nil {
		mails = make([]MailMessage, 0, 1)
	}
	mails = append(mails, message)
	m.Mails = mails
	return m
}

func (m *MailSender) AddMails(messages []MailMessage) *MailSender {
	mails := m.Mails
	if mails == nil {
		mails = make([]MailMessage, 0, 1)
	}
	mails = append(mails, messages...)
	m.Mails = mails
	return m
}

func (m *MailSender) Send() (int, error) {
	if len(m.Mails) > 0 && m.Mails != nil {
		count := 0
		// create dialer
		dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)
		if !m.TLSSecure {
			dialer.TLSConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		var errs []error
		for _, mail := range m.Mails {
			// build message
			message := gomail.NewMessage()
			message.SetHeader("From", m.Username)
			message.SetHeader("To", m.getReceiversList())
			if m.Cc != nil {
				for _, mailReceiver := range m.Cc {
					message.SetAddressHeader("Cc", mailReceiver.Address, mailReceiver.Name)
				}
			}
			message.SetHeader("Subject", mail.Title)
			switch mail.Type {
			case TextType:
				message.SetBody("text/plain", mail.Message.Message)
			case HtmlType:
				message.SetBody("text/html", mail.Message.Message)
			default:
				message.SetBody("text/plain", mail.Message.Message)
			}
			if mail.AttachesPath != nil {
				for _, attach := range mail.AttachesPath {
					message.Attach(attach)
				}
			}
			if err := dialer.DialAndSend(message); err != nil {
				errs = append(errs, err)
				continue
			}
			count += 1
		}
		if count < len(m.Mails) {
			return count, fmt.Errorf("发送邮件出现异常, errors: %+v", errs)
		} else {
			return count, nil
		}
	} else {
		return 0, nil
	}

}

func (m *MailSender) getReceiversList() string {
	result := ""
	if m.ReceiveAddress == nil || len(m.ReceiveAddress) < 1 {
		return ""
	}
	for _, receiver := range m.ReceiveAddress {
		if result == "" {
			result = receiver.Address
		} else {
			result = result + " " + receiver.Address
		}
	}
	return result
}

type MailConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	TLSSecure bool   `json:"tls_secure"`
}
