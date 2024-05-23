package token

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	cb = context.Background()
	TK *JWTAccessGenerate
)

type TokenGenerate interface {
	Token(userid int64, role, username string, expired time.Duration) (string, error)
	Verify(ctx context.Context, token string) (*model.UserInfo, bool)
}

// NewJWTAccessGenerate create to generate the jwt access token instance
func NewJWTAccessGenerate(cli *redis.Client, method jwt.SigningMethod) TokenGenerate {
	return &JWTAccessGenerate{
		Client:       cli,
		SignedMethod: method,
	}
}

// JWTAccessGenerate generate the jwt access token
type JWTAccessGenerate struct {
	Client       *redis.Client
	SignedMethod jwt.SigningMethod
}

// Token based on the UUID generated token
func (a *JWTAccessGenerate) Token(userid int64, role, username string, expired time.Duration) (string, error) {
	id := uuid.NewString()

	userinfo := &model.UserInfo{
		UserID: userid,
		Role:   role,
		Name:   username,
		UUID:   id,
	}
	b, _ := json.Marshal(userinfo)

	token := jwt.NewWithClaims(a.SignedMethod, &jwt.RegisteredClaims{
		Audience:  jwt.ClaimStrings{id},
		Subject:   string(b),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
	})

	key := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}
	a.Client.Set(cb, id, hex.EncodeToString(key), expired)
	access, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return access, nil
}

func (a *JWTAccessGenerate) Verify(ctx context.Context, token string) (*model.UserInfo, bool) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		aud, err := t.Claims.GetAudience()
		if err != nil {
			global.GVB_LOGGER.Errorln(err)
			return nil, err
		}
		id := aud[0]

		keyHex, err := a.Client.Get(ctx, id).Result()
		if err != nil {
			err := errors.New("cannot find key,user has logged out ")
			global.GVB_LOGGER.Errorln(err)
			return nil, err
		}
		key, _ := hex.DecodeString(keyHex)
		return key, nil
	})

	if err != nil || !tk.Valid {
		if err != nil {
			global.GVB_LOGGER.Errorln(err)
		}
		return nil, false
	}

	subject, _ := tk.Claims.GetSubject()
	var userinfo model.UserInfo
	json.Unmarshal([]byte(subject), &userinfo)

	return &userinfo, true
}
