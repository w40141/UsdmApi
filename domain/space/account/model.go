// Package account defines the domain model for the account entity.
package account

import "github.com/w40141/UsdmApi/domain/space/vo"

// E is an entity object for account.
type E struct {
	id   vo.AccountID
	name vo.AccountName
}

// C is a struct for creating a new account.
type C struct {
	name vo.AccountName
}

// D is a struct for deleting an account.
type D struct {
	id vo.AccountID
}
