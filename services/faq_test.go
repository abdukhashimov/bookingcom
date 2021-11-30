package services_test

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/logger"
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func createFAQ(t *testing.T) *model.Faq {
	res, err := servcesObj.FaqService().CreateFaq(context.Background(), model.CreateFaq{
		Question: faker.Sentence(),
		Answer:   faker.Sentence(),
		Active:   false,
	})

	if err != nil {
		log.Fatal("failed to create faq", logger.Error(err))
		return nil
	}

	return res
}

func TestCreateFAQ(t *testing.T) {
	faq := createFAQ(t)
	deleteFaq(t, faq.Slug)
}

func deleteFaq(t *testing.T, slug string) {
	_, err := servcesObj.FaqService().DeleteFaq(context.Background(), slug)
	assert.NoError(t, err, "failed to delete faq")
	if err != nil {
		log.Fatal("error while deleting faq", logger.Error(err))
		return
	}
}
