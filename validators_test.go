package wtforms

import (
	"testing"
)

type emailValidateTest struct {
	email string
	ok    bool
}

var emailvalidatetests = []emailValidateTest{
	{"username@example.com", true},
	{"aa.bb@example.com", true},
	{"aa_bb@example.com", true},
	{"aa-bb@example.com", true},
	{"aa123@example.com", true},
	{"AabB@example.com", true},
	{"Aa123.bb123@example.edu.cn", true},
	{"a@example.com", true},
	{"aa@example", false},
	{"<script>alert(1);</script>@test.com", false},
	{"aabbcc", false},
	{"aa@.com", false},
	{"aa.@xx.com", false},
	{"aa@", false},
	{"aa@bb@.example.com", false},
}

func TestEmailValidator(t *testing.T) {
	var email = Email{}

	for _, test := range emailvalidatetests {
		if ok, _ := email.CleanData(test.email); ok != test.ok {
			t.Error("test Email.CleanData:", test.email)
		}
	}
}
