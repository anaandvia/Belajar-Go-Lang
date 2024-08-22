package repository

import (
	"belajar-goglang-database/entity"
	"context"
)

type CommentRepository interface {
	insert(ctx context.Context, comment entity.Comments) (entity.Comments, error)
	FindById(ctx context.Context, id entity.Comments) (entity.Comments, error)
	FindAll(ctx context.Context) ([]entity.Comments, error)
}
