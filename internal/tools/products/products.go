package products

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/trenchesdeveloper/mcp-server-store/internal/client"
	"github.com/trenchesdeveloper/mcp-server-store/internal/mcp"
)

// ProductToolSet groups all product-related tools and shares the HTTP client.
type ProductToolSet struct {
	httpClient *client.RestClient
	logger     *logrus.Logger
}

// NewProductToolSet creates a new ProductToolSet with the given HTTP client and logger.
func NewProductToolSet(httpClient *client.RestClient, logger *logrus.Logger) *ProductToolSet {
	return &ProductToolSet{httpClient: httpClient, logger: logger}
}

// ---- List Products ----

// ListTool returns the tool definition for listing products.
func (p *ProductToolSet) ListTool() mcp.Tool {
	return mcp.Tool{
		Name:        "list_products",
		Description: "Lists products from the ecommerce store. Supports optional pagination with page and limit parameters.",
		InputSchema: mcp.InputSchema{
			Type: "object",
			Properties: map[string]mcp.Property{
				"page": {
					Type:        "string",
					Description: "Page number for pagination (default: 1)",
				},
				"limit": {
					Type:        "string",
					Description: "Number of products per page (default: 10)",
				},
			},
		},
	}
}

// ListHandler returns a handler that fetches products from the ecommerce API.
func (p *ProductToolSet) ListHandler() mcp.ToolHandler {
	return func(arguments map[string]interface{}) (*mcp.ToolCallResult, error) {
		p.logger.WithField("arguments", arguments).Info("Listing products")

		params := map[string]string{}

		if page, ok := arguments["page"].(string); ok && page != "" {
			params["page"] = page
		}
		if limit, ok := arguments["limit"].(string); ok && limit != "" {
			params["limit"] = limit
		}

		body, err := p.httpClient.Get("/products", params)
		if err != nil {
			p.logger.WithError(err).Error("Failed to list products")
			return nil, fmt.Errorf("failed to list products: %w", err)
		}

		var resp ProductResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			p.logger.WithError(err).Error("Failed to parse products response")
			return nil, fmt.Errorf("failed to parse products response: %w", err)
		}

		var sb strings.Builder
		fmt.Fprintf(&sb, "Found %d products\n\n", len(resp.Data))

		for i, product := range resp.Data {
			fmt.Fprintf(&sb, "%d. %s\n", i+1, formatProduct(product))
		}

		return &mcp.ToolCallResult{
			Content: []mcp.Content{
				{
					Type: "text",
					Text: sb.String(),
				},
			},
		}, nil
	}
}

func formatProduct(p Product) string {
	name := p.Name
	price := p.Price
	id := p.ID

	return fmt.Sprintf("**%s** (ID: %d) - $%.2f", name, id, price)
}

// ---- Search Products ----

// SearchTool returns the tool definition for searching products.
func (p *ProductToolSet) SearchTool() mcp.Tool {
	return mcp.Tool{
		Name:        "search_products",
		Description: "Full-text search products by name, SKU, and description with optional filters for category, price range, and pagination.",
		InputSchema: mcp.InputSchema{
			Type: "object",
			Properties: map[string]mcp.Property{
				"q": {
					Type:        "string",
					Description: "Search query (searches name, SKU, and description)",
				},
				"page": {
					Type:        "string",
					Description: "Page number for pagination (default: 1)",
				},
				"limit": {
					Type:        "string",
					Description: "Number of results per page (default: 10)",
				},
				"category_id": {
					Type:        "string",
					Description: "Filter by category ID",
				},
				"min_price": {
					Type:        "string",
					Description: "Minimum price filter",
				},
				"max_price": {
					Type:        "string",
					Description: "Maximum price filter",
				},
			},
			Required: []string{"q"},
		},
	}
}

// SearchHandler returns a handler that searches products via the ecommerce API.
func (p *ProductToolSet) SearchHandler() mcp.ToolHandler {
	return func(arguments map[string]interface{}) (*mcp.ToolCallResult, error) {
		p.logger.WithField("arguments", arguments).Info("Searching products")

		params := map[string]string{}

		for _, key := range []string{"q", "page", "limit", "category_id", "min_price", "max_price"} {
			if val, ok := arguments[key].(string); ok && val != "" {
				params[key] = val
			}
		}

		body, err := p.httpClient.Get("/products/search", params)
		if err != nil {
			p.logger.WithError(err).Error("Failed to search products")
			return nil, fmt.Errorf("failed to search products: %w", err)
		}

		var resp ProductResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			p.logger.WithError(err).Error("Failed to parse search response")
			return nil, fmt.Errorf("failed to parse search response: %w", err)
		}

		p.logger.WithField("count", len(resp.Data)).Info("Product search completed")

		var sb strings.Builder
		fmt.Fprintf(&sb, "Found %d products matching '%s'\n\n", len(resp.Data), params["q"])

		for i, product := range resp.Data {
			fmt.Fprintf(&sb, "%d. %s\n", i+1, formatProduct(product))
		}

		return &mcp.ToolCallResult{
			Content: []mcp.Content{
				{
					Type: "text",
					Text: sb.String(),
				},
			},
		}, nil
	}
}
