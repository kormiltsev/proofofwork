package quote

import (
	"context"
	"log"
	"os/exec"
)

// QuoteService imitate some job after proof.
type QuoteService struct{}

// New returns job service.
func New() *QuoteService {
	return &QuoteService{}
}

// GetQuote returns a piece wisdom.
func (q *QuoteService) GetQuote(ctx context.Context) string {
	cmd := exec.CommandContext(ctx, "fortune")

	out, err := cmd.Output()
	if err != nil {
		log.Println("err", err)
		return "Life Is Too Short To Remove USB Safely."
	}
	return string(out)
}
