package ads_enums

// AdCallToAction is the action button shown on an ad. It is the only thing
// that may be added to a promoted existing post, whose own copy and media
// cannot be changed.
type AdCallToAction string

const (
	// AdCallToActionShopNow sends the viewer to the shop.
	AdCallToActionShopNow AdCallToAction = "shop_now"
	// AdCallToActionLearnMore sends the viewer to a page with more detail.
	AdCallToActionLearnMore AdCallToAction = "learn_more"
	// AdCallToActionOrderNow sends the viewer straight into ordering.
	AdCallToActionOrderNow AdCallToAction = "order_now"
	// AdCallToActionGetOffer sends the viewer to a promotion or campaign.
	AdCallToActionGetOffer AdCallToAction = "get_offer"
	// AdCallToActionSignUp sends the viewer to registration.
	AdCallToActionSignUp AdCallToAction = "sign_up"
	// AdCallToActionSeeMore sends the viewer to a broader selection.
	AdCallToActionSeeMore AdCallToAction = "see_more"
)

// IsValid reports whether a is a known AdCallToAction.
func (a AdCallToAction) IsValid() bool {
	switch a {
	case AdCallToActionShopNow, AdCallToActionLearnMore, AdCallToActionOrderNow,
		AdCallToActionGetOffer, AdCallToActionSignUp, AdCallToActionSeeMore:
		return true
	}
	return false
}

func (a AdCallToAction) String() string { return string(a) }
