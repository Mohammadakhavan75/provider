package middleware

import (
	"context"
	// Import necessary Akash/Cosmos types (markettypes, sdkclient, providerstypes, etc.)
	markettypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
	providerstypes "github.com/akash-network/akash-api/go/node/provider/v1beta3"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// ChainClient defines the methods the bid engine needs
// to interact with the blockchain data, abstracting away the direct client context usage.
type ChainClient interface {
	// GetOrders fetches open orders matching specific filters from the market module.
	GetOrders(ctx context.Context, filters markettypes.OrderFilters) ([]markettypes.Order, error)

	// SubmitBid creates and broadcasts a MsgCreateBid transaction for a given order.
	// The provider address is implicitly determined from the client context/keyring used
	// during the middleware client's initialization.
	SubmitBid(ctx context.Context, orderID markettypes.OrderID, price sdktypes.DecCoin, deposit sdktypes.Coin) error

	// GetProviderInfo fetches the provider's information (including attributes) from the provider module.
	// The provider address is implicitly determined from the client context/keyring.
	GetProviderInfo(ctx context.Context) (*providerstypes.Provider, error)

	// GetAccountInfo fetches the account address and details (like sequence number)
	// for the provider's account, necessary for signing transactions.
	GetAccountInfo(ctx context.Context) (sdktypes.AccAddress, authtypes.AccountI, error)

	// GetMarketParams fetches the current parameters for the market module.
	GetMarketParams(ctx context.Context) (*markettypes.Params, error)

	// GetClientContext provides direct access to the underlying client context.
	// Use sparingly, prefer dedicated methods on this interface when possible.
	GetClientContext() sdkclient.Context
}
