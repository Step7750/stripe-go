package stripe

import (
	"encoding/json"
)

// A set of available payout methods for this blockchain address. Only values from this set should be passed as the `method` when creating a payout.
type BlockchainAddressAvailablePayoutMethod string

// List of values that BlockchainAddressAvailablePayoutMethod can take.
const (
	BlockchainAddressAvailablePayoutMethodStandard BlockchainAddressAvailablePayoutMethod = "standard"
)

// The supported network for this blockchain network.
type BlockchainAddressNetwork string

// List of values that BlockchainAddressNetwork can take.
const (
	BlockchainAddressNetworkPolygon BlockchainAddressNetwork = "polygon"
)

// Blockchain addresses are used as payout methods for [Connect Express accounts](https://stripe.com/docs/connect/express-accounts).
//
// This is only available to Stripe [Crypto Payouts](https://stripe.com/docs/connect/crypto-payouts) beta users.
//
// Related FAQ: [Crypto Payouts](https://support.stripe.com/express/questions/crypto-payouts).
type BlockchainAddress struct {
	APIResource
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ID of the account that the bank account is associated with.
	Account *Account `json:"account"`
	// A set of available payout methods for this blockchain address. Only values from this set should be passed as the `method` when creating a payout.
	AvailablePayoutMethods []BlockchainAddressAvailablePayoutMethod `json:"available_payout_methods"`
	// Three-letter [ISO code for the currency](https://stripe.com/docs/payouts) paid out to the blockchain address.
	Currency Currency `json:"currency"`
	// Whether this blockchain address is the default external account for its currency.
	DefaultForCurrency bool `json:"default_for_currency"`
	// Uniquely identifies this particular blockchain address. You can use this attribute to check whether two blockchain addresses are the same.
	Fingerprint string `json:"fingerprint"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Network of the blockchain address on-chain.
	Network BlockchainAddressNetwork `json:"network"`
}

// UnmarshalJSON handles deserialization of a BlockchainAddress.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BlockchainAddress) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type blockchainAddress BlockchainAddress
	var v blockchainAddress
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BlockchainAddress(v)
	return nil
}
