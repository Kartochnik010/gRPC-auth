package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Kartochnik010/go-sso/internal/domain/models"
	"github.com/Kartochnik010/go-sso/internal/lib/jwt"
	"github.com/Kartochnik010/go-sso/internal/lib/logger/sl"
	"github.com/Kartochnik010/go-sso/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppId       = errors.New("invalid app_id")
)

type UserSaver interface {
	SaveUser(ctx context.Context, email string, hash []byte) (int64, error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appId int64) (models.App, error)
}

type Auth struct {
	userSaver    UserSaver
	userProvider UserProvider
	appProvider  AppProvider
	log          *slog.Logger
	tokenTTL     time.Duration
}

func New(log *slog.Logger, uSaver UserSaver, uProvider UserProvider, appProvider AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		log:          log,
		userSaver:    uSaver,
		userProvider: uProvider,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string, appId int64) (string, error) {
	const op = "auth.Login"
	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
		slog.Int64("app_id", appId),
	)
	log.Info("logging user")

	user, err := a.userProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("user not found", sl.Err(err))
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
			// return "", fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		log.Error("failed to get user", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password)); err != nil {
		log.Error("invalid credentials", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	app, err := a.appProvider.App(ctx, appId)
	if err != nil {
		log.Error("couldn't fint app with id", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	token, err := jwt.NewToken(user, app, a.tokenTTL)
	if err != nil {
		log.Error("counln't generate token", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user logged in successfully")
	return token, nil
}
func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string) (int64, error) {
	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)
	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.userSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			log.Warn("user already exists", sl.Err(err))
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}

		log.Error("failed to save user", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
func (a *Auth) IsAdmin(ctx context.Context, userId int64) (bool, error) {
	const op = "auth.IsAdmin"

	log := a.log.With(
		slog.String("op", op),
		slog.Int64("user_id", userId),
	)
	log.Info("checking if user is admin")

	if userId == 0 {
		a.log.Error("zero value for user_id")
		return false, fmt.Errorf("%s: %s", op, "zero value for user_id")
	}

	ok, err := a.userProvider.IsAdmin(ctx, userId)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			log.Warn("user not found")
			return false, fmt.Errorf("%s: %w", op, ErrInvalidAppId)
		}
		a.log.Error("error from userProvider", sl.Err(err))
		return false, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("user is admin")
	return ok, nil
}
