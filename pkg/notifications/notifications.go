package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cli/go-gh/pkg/api"
)

// List notifications.
// If all is true read notifications are returned too.
// If participating is true returns also notification in which the user is directly participating or mentioned.
func List(client api.RESTClient, all, participating bool) ([]NotificationsResponse, error) {
	path := fmt.Sprintf(
		"notifications?all=%t&participating=%t",
		all,
		participating,
	)

	res := []NotificationsResponse{}
	err := client.Get(path, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarkAsRead the whole notification thread specified by threadID.
func MarkAsRead(client api.RESTClient, threadID string) error {
	path := fmt.Sprintf("notifications/threads/%s", threadID)
	return client.Patch(path, nil, nil)
}

func setThreadSubscription(client api.RESTClient, threadID string, ignore bool) error {
	path := fmt.Sprintf("notifications/threads/%s/subscription", threadID)
	body, err := json.Marshal(map[string]string{
		"ignore": strconv.FormatBool(ignore),
	})
	if err != nil {
		return err
	}
	return client.Put(path, bytes.NewBuffer(body), nil)
}

// IgnoreThread adds the thread specified by threadID to the list of thread
// the user should not be notified about.
func IgnoreThread(client api.RESTClient, threadID string) error {
	return setThreadSubscription(client, threadID, true)
}

// SubscribeThread adds the thread specified by threadID to the list of thread
// the user should be notified about.
func SubscribeThread(client api.RESTClient, threadID string) error {
	return setThreadSubscription(client, threadID, false)
}

// List notifications from the repository specified.
// If all is true read notifications are returned too.
// If participating is true returns also notification in which the user is directly participating or mentioned.
func ListFromRepo(client api.RESTClient, owner, repo string, all, participating bool) ([]NotificationsResponse, error) {
	path := fmt.Sprintf(
		"/repos/%s/%s/?all=%t&participating=%t",
		owner,
		repo,
		all,
		participating,
	)

	res := []NotificationsResponse{}
	err := client.Get(path, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
