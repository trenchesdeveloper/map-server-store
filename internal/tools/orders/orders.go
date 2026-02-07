package orders

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/trenchesdeveloper/mcp-server-store/internal/client"
	"github.com/trenchesdeveloper/mcp-server-store/internal/mcp"
)

// OrderToolSet groups all order-related tools and shares the HTTP client.
type OrderToolSet struct {
	httpClient *client.RestClient
	logger     *logrus.Logger
}

// NewOrderToolSet creates a new OrderToolSet with the given HTTP client and logger.
func NewOrderToolSet(httpClient *client.RestClient, logger *logrus.Logger) *OrderToolSet {
	return &OrderToolSet{httpClient: httpClient, logger: logger}
}

// ---- Create Order ----

// CreateOrderTool returns the tool definition for creating an order from the cart.
func (o *OrderToolSet) CreateOrderTool() mcp.Tool {
	return mcp.Tool{
		Name:        "place_order",
		Description: "Creates an order from the current shopping cart. Requires authentication.",
		InputSchema: mcp.InputSchema{
			Type: "object",
		},
	}
}

// CreateOrderHandler returns a handler that creates an order.
func (o *OrderToolSet) CreateOrderHandler() mcp.ToolHandler {
	return func(arguments map[string]interface{}) (*mcp.ToolCallResult, error) {
		o.logger.Info("Creating order from cart")

		body, err := o.httpClient.WithToken().Post("/orders", nil)
		if err != nil {
			o.logger.WithError(err).Error("Failed to create order")
			return nil, fmt.Errorf("failed to create order: %w", err)
		}

		var resp OrderResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			o.logger.WithError(err).Error("Failed to parse order response")
			return nil, fmt.Errorf("failed to parse order response: %w", err)
		}

		o.logger.WithFields(logrus.Fields{
			"order_id": resp.Data.ID,
			"total":    resp.Data.Total,
		}).Info("Order created")

		result := fmt.Sprintf("Order #%d created successfully!\n- Status: %s\n- Total: $%.2f",
			resp.Data.ID, resp.Data.Status, resp.Data.Total)

		return &mcp.ToolCallResult{
			Content: []mcp.Content{
				mcp.NewTextContent(result),
			},
		}, nil
	}
}
