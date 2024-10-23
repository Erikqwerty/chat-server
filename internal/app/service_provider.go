package app

import (
	"context"
	"log"

	"github.com/erikqwerty/chat-server/internal/api"
	"github.com/erikqwerty/chat-server/internal/closer"
	"github.com/erikqwerty/chat-server/internal/config"
	"github.com/erikqwerty/chat-server/internal/repository"
	chatserverrepository "github.com/erikqwerty/chat-server/internal/repository/chat"
	"github.com/erikqwerty/chat-server/internal/service"
	chatservice "github.com/erikqwerty/chat-server/internal/service/chat"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool               *pgxpool.Pool
	chatServerRepository repository.ChatServerRepository

	chatService service.ChatService

	chatServerImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("ошибка загрузки конфигурации базы данных: %s", err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("ошибка загрузки конфигурации базы данных: %s", err.Error())
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgConfig == nil {
		s.PGConfig()
	}
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.pgConfig.DSN())
		if err != nil {
			log.Fatalf("ошибка подключения к базе данных: %v", err)
		}
		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping до базы данных не проходит: %v", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})
		s.pgPool = pool
	}
	return s.pgPool
}

func (s *serviceProvider) ChatServerRepository(ctx context.Context) repository.ChatServerRepository {
	if s.chatServerRepository == nil {
		s.chatServerRepository = chatserverrepository.NewRepo(s.PgPool(ctx))
	}
	return s.chatServerRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatservice.NewService(s.ChatServerRepository(ctx))
	}
	return s.chatService
}

func (s *serviceProvider) ChatServerImpl(ctx context.Context) *api.Implementation {
	if s.chatServerImpl == nil {
		s.chatServerImpl = api.NewImplementation(s.ChatService(ctx))
	}
	return s.chatServerImpl
}
