package modules

type messageFcm struct {
	Message MessageContent `json:"message"`
}

type MessageContent struct {
	Notification NotificationContent `json:"notification"`
	Data         map[string]string   `json:"data"`
	Token        string              `json:"token"`
}

type NotificationContent struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
