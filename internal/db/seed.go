package db

import (
	"context"
	"github/Alfian57/student-hub-api/internal/store"
	"log"
	"math/rand"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func Seed(storage store.Storage) {
	ctx := context.Background()

	generateCategory(ctx, storage, 100)

	log.Println("Seeder Success")
}

func generateCategory(ctx context.Context, storage store.Storage, amount int) []uuid.UUID {
	categoryIDs := make([]uuid.UUID, 0, amount)
	categoryType := []string{"blog", "project"}

	for i := 0; i < amount; i++ {
		name := faker.Word()
		category := store.Category{
			ID:   uuid.New(),
			Slug: slug.Make(name),
			Name: name,
			Type: categoryType[rand.Intn(len(categoryType))],
		}
		categoryIDs = append(categoryIDs, category.ID)

		storage.Category.Create(ctx, &category)
	}

	return categoryIDs
}
