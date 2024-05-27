package quote

import (
	"context"
	"log"
	"os/exec"
)

type QuoteService struct{}

func New() *QuoteService {
	return &QuoteService{}
}

func (q *QuoteService) GetQuote(ctx context.Context) string {
	cmd := exec.CommandContext(ctx, "fortune")

	out, err := cmd.Output()
	if err != nil {
		log.Println("err", err)
		return "Life Is Too Short To Remove USB Safely."
	}
	return string(out)
}
