package bank

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/DATA-DOG/godog/gherkin"
)

var testAccount *account
var beneficiaryAccount *account
var errors error

var opt = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("account", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func createAccountWithBalance(balance int) *account {
	return &account{balance: balance}
}

func aUserBankAccountWith(balance int) error {
	testAccount = createAccountWithBalance(balance)
	return nil
}

func heDeposits(amount int) error {
	errors = testAccount.deposit(amount)
	return nil
}

func heWithdraws(amount int) error {
	errors = testAccount.withdraw(amount)
	return nil
}

func theUserShouldHaveABalanceOf(balance int) error {
	if testAccount.balance == balance {
		return nil
	}
	return fmt.Errorf("Incorrect account balance")
}

func aBeneficiaryBankAccountWith(balance int) error {
	beneficiaryAccount = createAccountWithBalance(balance)
	return nil
}

func heTransfers(amount int) error {
	errors = testAccount.transfer(beneficiaryAccount, amount)
	return nil
}

func theBeneficiaryShouldHaveABalanceOf(balance int) error {
	if beneficiaryAccount.balance == balance {
		return nil
	}
	return fmt.Errorf("Incorrect Beneficiary account balance")
}

func theSystemShouldThrowAnError(errMessage *gherkin.DocString) error {
	if errors != nil && errors.Error() == errMessage.Content {
		return nil
	}
	return fmt.Errorf("Expected error - %v, Acutal - %v ", errMessage.Content, errors)
}

func noErrorShouldBeFound() error {
	if errors != nil {
		return fmt.Errorf("Expected no errors. Found - %v ", errors)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a user bank account with (\d+)\$$`, aUserBankAccountWith)
	s.Step(`^he deposits (\d+)\$$`, heDeposits)
	s.Step(`^he withdraws (\d+)\$$`, heWithdraws)
	s.Step(`^the user bank account should have a balance of (\d+)\$$`, theUserShouldHaveABalanceOf)
	s.Step(`^a beneficiary bank account with (\d+)\$$`, aBeneficiaryBankAccountWith)
	s.Step(`^he transfers (\d+)\$$`, heTransfers)
	s.Step(`^the beneficiary bank account should have a balance of (\d+)\$$`, theBeneficiaryShouldHaveABalanceOf)
	s.Step(`^the system should throw an error$`, theSystemShouldThrowAnError)
	s.Step(`^no error should be found$`, noErrorShouldBeFound)

	s.BeforeScenario(func(interface{}) {
		testAccount = nil
		beneficiaryAccount = nil
		errors = nil
	})
}
