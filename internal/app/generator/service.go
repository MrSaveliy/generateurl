package generator

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"name/generator/internal/pkg/repository"
	"name/generator/internal/pkg/service"
	"name/generator/internal/pkg/store"
	in_memory "name/generator/internal/pkg/store/in-memory"
	"name/generator/pkg/api/generator"
)

type Implementation struct {
	generator.UnimplementedGeneratorServer

	r            repository.Repository
	urlGenerator *service.UrlGenerator
}

func NewImplementation() *Implementation {
	return &Implementation{
		r:            store.GetStore(),
		urlGenerator: service.NewUrlGenerator(),
	}
}

func (i *Implementation) GetUrl(ctx context.Context, req *generator.GetUrlRequest) (*generator.GetUrlResponse, error) {
	url, err := i.r.Get(req.GetShort())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, in_memory.ErrShortNotExist) {
			return nil, status.Errorf(codes.NotFound, "repository.Get: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "repository.Get: %v", err)
	}
	return &generator.GetUrlResponse{
		Url: url,
	}, nil
}

func (i *Implementation) SetUrl(ctx context.Context, req *generator.SetUrlRequest) (*generator.SetUrlResponse, error) {
	id, err := i.r.GetLastID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "repository.GetLastID: %v", err)
	}
	short := i.urlGenerator.GetShortByID(uint64(id))
	if err = i.r.Set(req.GetUrl(), short); err != nil {
		return nil, status.Errorf(codes.Internal, "repository.Set: %v", err)
	}
	return &generator.SetUrlResponse{
		Short: short,
	}, nil
}
