package objects

// The invoice number or bill number. This number could be provided by the merchant or could be an indication for the mobile application to prompt the consumer to input a Bill Number.
// For example, the Bill Number may be present when the QR Code is used for bill payment.
type AdditionalObjectBillNumber struct {
	DataObject
}

// The mobile number could be provided by the merchant or could be an indication for the mobile application to prompt the consumer to input a Mobile Number.
// For example, the Mobile Number to be used for multiple use cases, such as mobile top-up and bill payment.
type AdditionalObjectMobileNumber struct {
	DataObject
}

// A distinctive value associated to a store. This value could be provided by the merchant or could be an indication for the mobile application to prompt the consumer to input a Store Label.
// For example, the Store Label may be displayed to the consumer on the mobile application identifying a specific store.
type AdditionalObjectStoreLabel struct {
	DataObject
}

// Typically, a loyalty card number. This number could be provided by the merchant, if known, or could be an indication for the mobile application to prompt the consumer to input their Loyalty Number.
type AdditionalObjectLoyaltyNumber struct {
	DataObject
}

// Any value as defined by the merchant or acquirer in order to identify the transaction. This value could be provided by the merchant or could be an indication for the mobile app to prompt the consumer to input a transaction Reference Label.
// For example, the Reference Label may be used by the consumer mobile application for transaction logging or receipt display.
type AdditionalObjectReferenceLabel struct {
	DataObject
}

// Any value identifying a specific consumer. This value could be provided by the merchant (if known), or could be an indication for the mobile application to prompt the consumer to input their Customer Label.
// For example, the Customer Label may be a subscriber ID for subscription services, a student enrolment number, etc.
type AdditionalObjectCustomerLabel struct {
	DataObject
}

// A distinctive value associated to a terminal in the store. This value could be provided by the merchant or could be an indication for the mobile application to prompt the consumer to input a Terminal Label.
// For example, the Terminal Label may be displayed to the consumer on the mobile application identifying a specific terminal.
type AdditionalObjectTerminalLabel struct {
	DataObject
}

// Any value defining the purpose of the transaction. This value could be
// provided by the merchant or could be an indication for the mobile
// application to prompt the consumer to input a value describing the purpose
// of the transaction.
// For example, the Purpose of Transaction may have the value
// "International Data Package" for display on the mobile application.
type AdditionalObjectPurposeOfTransaction struct {
	DataObject
}

// Contains indications that the mobile application is to provide the requested
// information in order to complete the transaction. The information requested
// should be provided by the mobile application in the authorization without
// unnecessarily prompting the consumer.
// For example, the Additional Consumer Data Request may indicate that the
// consumer mobile number is required to complete the transaction, in which
// case the mobile application should be able to provide this number (that the
// mobile application has previously stored) without unnecessarily prompting
// the consumer
type AdditionalObjectConsumerDataRequest struct {
	DataObject
}

// The tax identification number of the merchant, assigned by the governmental body of the country in which the EMV merchant-presented QR code is being used/displayed.
// For example, the Merchant Tax ID may be used by the consumer mobile application for receipt display.
type AdditionalObjectMerchantTaxID struct {
	DataObject
}

// A merchant channel establishes the environment in which a QR Code is presented to the consumer. Covering use cases such as retail outlet, Ecommerce, bill payment with the purpose of improving transaction reporting.
type AdditionalObjectMerchantChannel struct {
	DataObject
}

type AdditionalObjectRFU struct {
	DataObject
}

type AdditionalObjectPaymentSystemSpecificTemplate struct {
	DataObject
}
