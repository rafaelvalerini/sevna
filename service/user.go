package service

import (
	"agregador/model"
	"agregador/repository"
    "gopkg.in/gomail.v2"
	"crypto/tls"
)

const (
	MAIL_FROM = "rafael@futuroclick.com.br"
	MAIL_RESET_SUBJECT = "Reset de senha"
	MAIL_URL = "mail.futuroclick.com.br"
	MAIL_PORT = 587
	MAIL_USER = "rafael@futuroclick.com.br" 
	MAIL_PASSWORD = "Rafilkis1536*"
)

func ResetPassword(mail string) (user model.User){

    user = repository.FindUserByMail(mail)

    if user.Id <= 0 {
    	
    	return user

    }else{

    	return sendResetMail(user)

    }

}

func sendResetMail(user model.User) (userRetunr model.User) {

	m := gomail.NewMessage()
	m.SetHeader("From", MAIL_FROM)
	m.SetHeader("To", user.Mail)
	m.SetHeader("Subject", MAIL_RESET_SUBJECT)
	m.SetBody("text/plain", "Nome: "+user.Username+", E-Mail: "+user.Mail+", Senha: "+user.Password)

	d := gomail.NewPlainDialer(MAIL_URL, MAIL_PORT, MAIL_USER, MAIL_PASSWORD)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	d.SSL = false

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return user

}