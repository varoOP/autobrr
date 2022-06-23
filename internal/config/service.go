package config

import (
	"context"

	"github.com/autobrr/autobrr/internal/domain"
	"github.com/autobrr/autobrr/internal/logger"
	releasechecker "github.com/autobrr/autobrr/pkg/gh-release-checker"

	"github.com/rs/zerolog"
)

type Service interface {
}

type service struct {
	log    zerolog.Logger
	config domain.Config
}

func NewService(log logger.Logger, config domain.Config) Service {
	return &service{
		log:    log.With().Str("service", "config").Logger(),
		config: config,
	}
}

func (s *service) CheckUpdates() {
	r := releasechecker.GitHubReleaseChecker{Repo: "autobrr/autobrr"}
	r.CanUpdate(context.TODO(), s.config.Version)

}
