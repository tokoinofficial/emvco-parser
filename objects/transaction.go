package objects

// The transaction amount (excluding tips and convenience fees), if known. For instance, "99.34". If present, this value is displayed to the consumer by the mobile application when processing the transaction. If this data object is not present, the consumer is prompted to input the transaction amount to be paid to the merchant.
type TransactionAmount struct {
	DataObject
}

// Indicates the currency code of the transaction.
// A 3-digit numeric value, as defined by [ISO 4217]. This value will be used by the mobile application to display a recognizable currency to the consumer whenever an amount is being displayed or whenever the consumer is prompted to enter an amount.
type TransactionCurrency struct {
	DataObject
}

// Indicates whether the consumer will be prompted to enter a tip or whether the merchant has determined that a flat, or percentage convenience fee is charged.
type TransactionTip struct {
	DataObject
}

// The fixed amount convenience fee when 'Tip or Convenience Indicator' indicates a flat convenience fee.
// For example, "9.85", indicating that this fixed amount (in the transaction currency) will be charged on top of the transaction amount.
type TransactionValueOfConvenienceFeeFixed struct {
	DataObject
}

// The percentage convenience fee when 'Tip or Convenience Indicator' indicates a percentage convenience fee.
// For example, "3.00" indicating that a convenience fee of 3% of the transaction amount will be charged, on top of the transaction amount.
type TransactionValueOfConvenienceFeePercentage struct {
	DataObject
}
