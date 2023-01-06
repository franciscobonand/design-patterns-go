package main

import "strings"

// To make the user properly create a structure (e.g. filling up all the required fields)
// you can create a private stricture (email) and a public builder for that structure (EmailBuilder)
type email struct {
    from, to, subject, body string
}

type EmailBuilder struct {
    email email
}

// Each builder method can contain a specific validation for the field it handles
func (b *EmailBuilder) From(from string) *EmailBuilder {
    if !strings.Contains(from, "@") {
        panic("Email should contain '@'")
    }

    b.email.from = from
    return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
    b.email.to = to
    return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
    b.email.subject = subject
    return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
    b.email.body = body
    return b
}

// sendMailImpl is responsible for sending the email received as a parameter 
func sendMailImpl(email *email) {
    // implementation
}

// The SendEmail exported function enforces the use of the EmailBuilder API,
// so the 'email' structure is properly created behind the scenes
// before sending the email itself
type build func(*EmailBuilder)
func SendEmail(action build) {
    builder := EmailBuilder{}
    action(&builder)
    sendMailImpl(&builder.email)
}

func mainBP() {
    SendEmail(func(b *EmailBuilder) {
        b.From("foo@bar.com").
            To("bar@baz.com").
            Subject("Meeting").
            Body("Hello, lets meet?")
    })
}
