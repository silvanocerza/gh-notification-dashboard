package notifications

// NotificationResponse stores all the data that we need to
// show a notification information to a user.
type NotificationsResponse struct {
	ID              string
	Unread          bool
	Reason          string
	UpdatedAt       string
	LastReadAt      string
	Subject         subject
	Repository      repository
	URL             string
	SubscriptionURL string
}

type subject struct {
	Title            string
	URL              string
	LatestCommentURL string
	Type             string
}

type repository struct {
	ID       int
	Name     string
	FullName string
	HtmlURL  string
	APIURL   string
}
