package services_test

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/logger"
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func createCategory(t *testing.T) *model.Category {
	information := faker.Sentence()
	res, err := servcesObj.CategoryService().CreateCategory(
		context.Background(),
		model.CreateCategory{
			Name:        faker.Word(),
			Information: &information,
			Image:       faker.Word(),
			Active:      false,
		},
	)

	if err != nil {
		log.Fatal("failed to create Category", logger.Error(err))
		return nil
	}

	return res
}

func TestCreateCategory(t *testing.T) {
	Category := createCategory(t)
	log.Info("assert created slug length")
	assert.GreaterOrEqual(t, len(Category.Slug), 1)
	deleteCategory(t, Category.Slug)
}

func TestGetCategory(t *testing.T) {
	Category := createCategory(t)
	for _, lang := range langs {
		CategoryObj, err := servcesObj.CategoryService().GetCategory(context.Background(), Category.Slug, lang)
		assert.NoErrorf(t, err, "failed to retreive obj in %s", lang)
		assert.Equal(t, CategoryObj.Name, Category.Name)
	}

	deleteCategory(t, Category.Slug)
}

func TestGetAllCategory(t *testing.T) {
	Category := createCategory(t)
	limit := 10
	offset := 0

	for _, lang := range langs {
		CategorysResp, err := servcesObj.CategoryService().GetAllCategory(context.Background(), &limit, &offset, lang)
		assert.NoError(t, err, "failed to retreive Categorys")
		assert.GreaterOrEqual(t, len(CategorysResp.Categories), 1, "Categorys array failed to be greater or equal to 1")
	}

	deleteCategory(t, Category.Slug)
}

func TestUpdateCategory(t *testing.T) {
	Category := createCategory(t)
	updated, err := servcesObj.CategoryService().UpdateCategory(context.Background(), model.UpdateCategory{
		Name:   "updated",
		Image:  "updated",
		Active: true,
		Slug:   Category.Slug,
		Lang:   Category.Lang,
	})
	assert.NoError(t, err, "failed to update Category")

	dbCategory, err := servcesObj.CategoryService().GetCategory(context.Background(), Category.Slug, Category.Lang)
	assert.NoError(t, err, "failed to retreive Category")
	assert.Equal(t, dbCategory.Name, updated.Name, "updated fields failed to match: question")
	assert.Equal(t, dbCategory.Image, updated.Image, "updated fields failed to match: answer")
	assert.Equal(t, dbCategory.Active, updated.Active, "updated fields failed to match: active")

	deleteCategory(t, Category.Slug)
}

func deleteCategory(t *testing.T, slug string) {
	_, err := servcesObj.CategoryService().DeleteCategory(context.Background(), slug)
	assert.NoError(t, err, "failed to delete Category")
	if err != nil {
		log.Fatal("error while deleting Category", logger.Error(err))
		return
	}
}
