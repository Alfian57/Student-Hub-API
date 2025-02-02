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

	userIDs := generateUser(ctx, storage, 100)

	blogCategoryIDs := generateCategory(ctx, storage, 100, store.CategoryTypeBlog)
	generateBlog(ctx, storage, 100, userIDs, blogCategoryIDs)

	projectCategoryIDs := generateCategory(ctx, storage, 100, store.CategoryTypeProject)
	generateProject(ctx, storage, 100, userIDs, projectCategoryIDs)

	log.Println("Seeder Success")
}

func generateUser(ctx context.Context, storage store.Storage, amount int) []uuid.UUID {
	userIDs := make([]uuid.UUID, 0, amount)
	userRole := []string{store.UserRoleAdmin, store.UserRoleUser}

	for i := 0; i < amount; i++ {
		user := store.User{
			ID:             uuid.New(),
			Name:           faker.Name(),
			Username:       faker.Username(),
			Email:          faker.Email(),
			Password:       faker.Password(),
			ProfilePicture: faker.URL(),
			Bio:            faker.Sentence(),
			Role:           userRole[rand.Intn(len(userRole))],
		}
		userIDs = append(userIDs, user.ID)

		storage.User.Create(ctx, &user)
	}

	log.Println("User Seeder Success")
	return userIDs
}

func generateCategory(ctx context.Context, storage store.Storage, amount int, categoryType string) []uuid.UUID {
	categoryIDs := make([]uuid.UUID, 0, amount)

	for i := 0; i < amount; i++ {
		name := faker.Word()
		category := store.Category{
			ID:   uuid.New(),
			Slug: slug.Make(name),
			Name: name,
			Type: categoryType,
		}
		categoryIDs = append(categoryIDs, category.ID)

		storage.Category.Create(ctx, &category)
	}

	log.Println("Category Seeder Success")
	return categoryIDs
}

func generateBlog(ctx context.Context, storage store.Storage, amount int, userIDs []uuid.UUID, categoryIDs []uuid.UUID) []uuid.UUID {
	blogIDs := make([]uuid.UUID, 0, amount)

	for i := 0; i < amount; i++ {
		title := faker.Word()
		blog := store.Blog{
			ID:         uuid.New(),
			UserID:     userIDs[rand.Intn(len(userIDs))],
			CategoryID: categoryIDs[rand.Intn(len(categoryIDs))],
			Slug:       slug.Make(title),
			Title:      title,
			Content:    faker.Sentence(),
			Thumbnail:  faker.URL(),
			IsPublish:  rand.Intn(2) == 1,
		}
		blogIDs = append(blogIDs, blog.ID)

		storage.Blog.Create(ctx, &blog)
	}

	log.Println("Blog Seeder Success")
	return blogIDs
}

func generateProject(ctx context.Context, storage store.Storage, amount int, userIDs []uuid.UUID, categoryIDs []uuid.UUID) []uuid.UUID {
	projectIDs := make([]uuid.UUID, 0, amount)

	for i := 0; i < amount; i++ {
		title := faker.Word()
		project := store.Project{
			ID:          uuid.New(),
			UserID:      userIDs[rand.Intn(len(userIDs))],
			CategoryID:  categoryIDs[rand.Intn(len(categoryIDs))],
			Slug:        slug.Make(title),
			Title:       title,
			Description: faker.Sentence(),
			Thumbnail:   faker.URL(),
			IsPublish:   rand.Intn(2) == 1,
			CodeLink:    faker.URL(),
			AppLink:     faker.URL(),
		}
		projectIDs = append(projectIDs, project.ID)

		storage.Project.Create(ctx, &project)
	}

	log.Println("Project Seeder Success")
	return projectIDs
}
