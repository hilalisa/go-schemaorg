package corporation

import "github.com/dpb587/go-schemaorg"

var (
	// The exchange traded instrument associated with a Corporation object. The
	// tickerSymbol is expressed as an exchange and an instrument name separated by
	// a space character. For the exchange component of the tickerSymbol attribute,
	// we reccommend using the controlled vocaulary of Market Identifier Codes (MIC)
	// specified in ISO15022.
	TickerSymbol = schemaorg.NewProperty("tickerSymbol")
)
