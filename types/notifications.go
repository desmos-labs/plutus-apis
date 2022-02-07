package types

// Notification contains all the data that might be needed in order to send a notification
type Notification struct {
	// RecipientAppAccount contains the data of the application on which the donation was made (Twitter, Twitch, etc)
	RecipientAppAccount *ApplicationAccount

	// Tx contains the details of the donation transaction
	Tx *DonationTx

	// TipperUsername contains the username the tipper has chosen to use for this donation
	TipperUsername string

	// Donation message
	Message string
}

// NewNotification returns a new Notification instance
func NewNotification(
	recipientAppAccount *ApplicationAccount, donationTx *DonationTx, tipperUsername, message string,
) *Notification {
	return &Notification{
		RecipientAppAccount: recipientAppAccount,
		Tx:                  donationTx,
		TipperUsername:      tipperUsername,
		Message:             message,
	}
}
