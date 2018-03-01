package localbusiness

import "github.com/dpb587/go-schemaorg"

var (
	// The price range of the business, for example <code>$$$</code>.
	PriceRange = schemaorg.NewProperty("priceRange")

	// The larger organization that this local business is a branch of, if any. Not
	// to be confused with (anatomical)<a class="localLink"
	// href="http://schema.org/branch">branch</a>.
	BranchOf = schemaorg.NewProperty("branchOf")

	// Cash, credit card, etc.
	PaymentAccepted = schemaorg.NewProperty("paymentAccepted")

	// <p>The general opening hours for a business. Opening hours can be specified
	// as a weekly time range, starting with days, then times per day. Multiple days
	// can be listed with commas ',' separating each day. Day or time ranges are
	// specified using a hyphen '-'.</p>
	// 
	// <ul>
	// <li>Days are specified using the following two-letter combinations:
	// <code>Mo</code>, <code>Tu</code>, <code>We</code>, <code>Th</code>,
	// <code>Fr</code>, <code>Sa</code>, <code>Su</code>.</li>
	// <li>Times are specified using 24:00 time. For example, 3pm is specified as
	// <code>15:00</code>. </li>
	// <li>Here is an example: <code>&lt;time itemprop="openingHours"
	// datetime=&quot;Tu,Th 16:00-20:00&quot;&gt;Tuesdays and Thursdays
	// 4-8pm&lt;/time&gt;</code>.</li>
	// <li>If a business is open 7 days a week, then it can be specified as
	// <code>&lt;time itemprop=&quot;openingHours&quot;
	// datetime=&quot;Mo-Su&quot;&gt;Monday through Sunday, all
	// day&lt;/time&gt;</code>.</li>
	// </ul>
	// 
	OpeningHours = schemaorg.NewProperty("openingHours")

	// The currency accepted (in <a href="http://en.wikipedia.org/wiki/ISO_4217">ISO
	// 4217 currency format</a>).
	CurrenciesAccepted = schemaorg.NewProperty("currenciesAccepted")
)
