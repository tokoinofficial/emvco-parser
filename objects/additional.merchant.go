package objects

// As defined by [ISO 18245] and assigned by the Acquirer.
type MerchantCategoryCode struct {
	DataObject
}

// Indicates the country of the merchant acceptance device.
// A 2-character alpha value, as defined by [ISO 3166-1 alpha 2] and assigned by the Acquirer. The country may be displayed to the consumer by the mobile application when processing the transaction.
type MerchantCountryCode struct {
	DataObject
}

// The “doing business as” name for the merchant, recognizable to the consumer. This name may be displayed to the consumer by the mobile application when processing the transaction.
type MerchantName struct {
	DataObject
}

// City of operations for the merchant. This name may be displayed to the consumer by the mobile application when processing the transaction.
type MerchantCity struct {
	DataObject
}

// Zip code or Pin code or Postal code of the merchant. If present, this value may be displayed to the consumer by the mobile application when processing the transaction.
type MerchantPostalCode struct {
	DataObject
}

// Merchant Name and potentially other merchant related information in an alternate language, typically the local language.
type MerchantAlternateLanguage struct {
	DataObject
}
