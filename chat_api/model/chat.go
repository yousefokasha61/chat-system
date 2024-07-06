package model

const (
	InvalidRequest = "invalid_request"
)

type HttpMainResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type EmptyData struct{}

type Application struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Chats int32  `json:"chats"`
}

type Chat struct {
	Number           int32  `json:"number"`
	ApplicationToken string `json:"applicationToken"`
	Messages         int32  `json:"messages"`
}

type Message struct {
	Number int32  `json:"number"`
	Body   string `json:"body"`
}

type CreateApplicationRequest struct {
	Name string `json:"name" required:"true"`
}

type CreateApplicationResponse struct {
	Application Application `json:"application"`
}

type UpdateApplicationRequest struct {
	Name string `json:"name" required:"true"`
}

type UpdateApplicationResponse struct {
	Application Application `json:"application"`
}

type GetApplicationResponse struct {
	Application Application `json:"application"`
}

type CreateChatResponse struct {
	Chat Chat `json:"chat"`
}

type GetChatsResponse struct {
	Chats []Chat `json:"chats"`
}

type CreateMessageRequest struct {
	Body string `json:"body" required:"true"`
}

type CreateMessageResponse struct {
	Message Message `json:"message"`
}

type GetMessagesResponse struct {
	Messages []Message `json:"messages"`
}

type SearchMessagesRequest struct {
	Query string `json:"messages"`
}
type SearchMessagesResponse struct {
	Messages []Message `json:"messages"`
}
