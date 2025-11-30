package asset

import (
	"context"
)

type AssetRepository interface {
	FetchAssetsByTicketID(ctx context.Context, ticketID string) ([]Asset, error)
}
