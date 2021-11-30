package services_test

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/logger"
	"context"
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
)

func createFAQ(t *testing.T) *model.Faq {
	res, err := svs.FaqService().CreateFaq(context.Background(), model.CreateFaq{
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
	fmt.Println(faq)
}
