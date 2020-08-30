package manager

import "errors"

var (
	BidderParamRequiredError = errors.New(`"bidder" query param is required`)
	UnsupportedBidderError   = errors.New("bidder is not supported")
	EmptyUIDError            = errors.New("UID should be not empty")
)
