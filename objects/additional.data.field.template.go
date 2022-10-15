package objects

// The Additional Data Field Template includes information that may be provided by the Merchant or may be populated by the mobile application to enable or facilitate certain use cases.
// For the list of data objects that can be included in this template, please refer to Table 3.7.
type AdditionalDataFieldTemplate struct {
	DataObject
	BillNumber                     *AdditionalObjectBillNumber
	MobileNumber                   *AdditionalObjectMobileNumber
	StoreLabel                     *AdditionalObjectStoreLabel
	LoyaltyNumber                  *AdditionalObjectLoyaltyNumber
	ReferenceLabel                 *AdditionalObjectReferenceLabel
	CustomerLabel                  *AdditionalObjectCustomerLabel
	TerminalLabel                  *AdditionalObjectTerminalLabel
	PurposeOfTnx                   *AdditionalObjectPurposeOfTransaction
	AdditionalConsumerDataRequest  *AdditionalObjectConsumerDataRequest
	MerchantTaxID                  *AdditionalObjectMerchantTaxID
	MerchantChannel                *AdditionalObjectMerchantChannel
	PaymentSystemSpecificTemplates []*AdditionalObjectPaymentSystemSpecificTemplate
}
