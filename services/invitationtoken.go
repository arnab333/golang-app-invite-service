package services

import (
	"errors"
	"time"

	"github.com/arnab333/golang-app-invite-service/helpers"
)

type InviteToken struct {
	Token     string    `json:"token"`
	IsActive  bool      `json:"isActive"`
	ExpiresAt time.Time `json:"-"`
}

var InvitationTokens []InviteToken

func CreateInviteToken(authorID string) (string, error) {
	it := "INVITE-" + helpers.GetRandomString(16)

	expires := time.Now().Add(time.Hour * 24 * 7)

	InvitationTokens = append(InvitationTokens, InviteToken{
		Token:     it,
		IsActive:  true,
		ExpiresAt: expires,
	})

	return it, nil
}

func DisableInviteToken(key string) InviteToken {
	var token InviteToken
	for idx, v := range InvitationTokens {
		if v.Token == key {
			InvitationTokens[idx].IsActive = false
			token = InvitationTokens[idx]
			break
		}
	}
	return token
}

func GetAllInviteTokens() []InviteToken {
	return InvitationTokens
}

func GetInviteToken(key string) (*InviteToken, error) {
	var inviteToken InviteToken
	for _, v := range InvitationTokens {
		if v.Token == key && v.IsActive {
			inviteToken = v
			break
		}
	}
	if inviteToken.Token != "" {
		return &inviteToken, nil
	}

	return nil, errors.New("invalid invite token")
}

// func DeleteInviteToken(key string) string {
// 	for idx, v := range InvitationTokens {
// 		if v.Token == key {
// 			InvitationTokens = append(InvitationTokens[0:idx], InvitationTokens[idx+1:]...)
// 		}
// 	}
// 	return "Successfully deleted invite token"
// }
