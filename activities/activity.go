package activities

import (
	"context"
	"net/http"

	"go.temporal.io/sdk/activity"
)

// GetCat is a workflow activity that gets cat facts
func GetCat(ctx context.Context) (fact string, err error) {
	logger := activity.GetLogger(ctx)
	response, err := http.Get("https://cat-fact.herokuapp.com/facts/random")
	if err != nil {
		logger.i
	}

	if response.
}
