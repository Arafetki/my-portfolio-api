package repository

import (
	"github.com/Arafetki/my-portfolio-api/internal/db/store"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type Repository struct {
	Article interface {
		Create(ar *models.Article) error
		Delete(id int) error
	}
}

func NewRepo(store *store.Store) *Repository {
	return &Repository{
		Article: ArticleRepo{store},
	}
}
