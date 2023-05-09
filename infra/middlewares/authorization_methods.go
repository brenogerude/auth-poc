package middleware

import "errors"

var authorizationsMethods = map[string]func(s Session, subject string) (bool, error){
	"users":         checkUserAccess,
	"bank_accounts": checkBankAccountAccess,
}

func checkUserAccess(s Session, subject string) (bool, error) {
	switch s.Role {
	case MemberRole:
		return s.isSessionOwner(subject), nil
	case ClientRole:
		return s.isConnectRelated(subject)
	}
	return false, errors.New("unable to authorize")
}

func checkBankAccountAccess(s Session, subject string) (bool, error) {
	switch s.Role {
	case MemberRole:
		return s.isBankAccountMemberOwner(subject), nil
	case ClientRole:
		return s.isBankAccountClientOwnerOrRelated(subject)
	}
	return false, errors.New("unable to authorize")
}

func (s Session) isSessionOwner(subject string) bool {
	return s.UserID == subject
}
func (s Session) isConnectRelated(subject string) (bool, error) {
	//check if session client and subject user are linked by a common connection
	return false, nil
}

func (s Session) isBankAccountMemberOwner(subject string) bool {
	//getBankAccountUserID
	bankAccountUserID := "2"
	return s.UserID == bankAccountUserID
}

func (s Session) isBankAccountClientOwnerOrRelated(subject string) (bool, error) {
	//getBankAccountCompanyID and userID
	bankAccountCompanyID := ""
	bankAccountUserID := "2"
	//check it the bank_account is holded by this company
	if s.CompanyID == bankAccountCompanyID {
		return true, nil
	}
	//if not check also if this company have relation with the user holder
	return s.isConnectRelated(bankAccountUserID)
}
