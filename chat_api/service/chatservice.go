package service

import (
	"chat/ctx"
	"chat/grpc/client"
	"chat/model"
	proto2 "chat/proto"
	"context"
	"github.com/sirupsen/logrus"
)

type ChatSystemService struct {
	railsGrpcClient proto2.ChatSystemClient
	logger          *logrus.Logger
}

func NewChatSystemService(serviceContext ctx.ServiceContext) *ChatSystemService {
	return &ChatSystemService{
		railsGrpcClient: client.NewRailsGrpcClient(serviceContext.RailsGrpcClient()).Client,
		logger:          serviceContext.Logger(),
	}
}

func (s *ChatSystemService) CreateApplication(req *model.CreateApplicationRequest) (*model.CreateApplicationResponse, error) {
	backgroundContext := context.Background()

	res, err := s.railsGrpcClient.CreateApplication(backgroundContext, &proto2.CreateApplicationRequest{Name: req.Name})
	if err != nil {
		s.logger.Errorf("CreateApplication: error while invoking command %v", err)
		return nil, err
	}

	return &model.CreateApplicationResponse{
		Application: model.Application{
			Token: res.Token,
			Name:  res.Name,
			Chats: res.ChatsCount,
		}}, nil
}

func (s *ChatSystemService) UpdateApplication(req *model.UpdateApplicationRequest, id string) (*model.UpdateApplicationResponse, error) {
	backgroundContext := context.Background()

	res, err := s.railsGrpcClient.UpdateApplication(backgroundContext, &proto2.UpdateApplicationRequest{Name: req.Name, Token: id})
	if err != nil {
		s.logger.Errorf("UpdateApplication: error while invoking command %v", err)
		return nil, err
	}

	return &model.UpdateApplicationResponse{
		Application: model.Application{
			Token: res.Token,
			Name:  res.Name,
			Chats: res.ChatsCount,
		}}, nil
}

func (s *ChatSystemService) GetApplication(id string) (*model.GetApplicationResponse, error) {
	backgroundContext := context.Background()

	res, err := s.railsGrpcClient.GetApplication(backgroundContext, &proto2.GetApplicationRequest{Token: id})
	if err != nil {
		s.logger.Errorf("GetApplication: error while invoking command %v", err)
		return nil, err
	}

	return &model.GetApplicationResponse{
		Application: model.Application{
			Token: res.Token,
			Name:  res.Name,
			Chats: res.ChatsCount,
		}}, nil
}

func (s *ChatSystemService) CreateChat(id string) (*model.CreateChatResponse, error) {
	backgroundContext := context.Background()

	res, err := s.railsGrpcClient.CreateChat(backgroundContext, &proto2.CreateChatRequest{ApplicationToken: id})
	if err != nil {
		s.logger.Errorf("CreateChat: error while invoking command %v", err)
		return nil, err
	}

	return &model.CreateChatResponse{
		Chat: model.Chat{
			Number:           res.Number,
			ApplicationToken: res.ApplicationToken,
			Messages:         res.MessagesCount,
		}}, nil
}

func (s *ChatSystemService) ReadChats(id string) (*model.GetChatsResponse, error) {
	backgroundContext := context.Background()

	res, err := s.railsGrpcClient.GetChats(backgroundContext, &proto2.GetChatsRequest{ApplicationToken: id})
	if err != nil {
		s.logger.Errorf("ReadChats: error while invoking command %v", err)
		return nil, err
	}

	chats := make([]model.Chat, 0)
	for _, chat := range res.Chats {
		chats = append(chats, model.Chat{
			Number:           chat.Number,
			ApplicationToken: chat.ApplicationToken,
			Messages:         chat.MessagesCount,
		})
	}

	return &model.GetChatsResponse{
		Chats: chats,
	}, nil
}

func (s *ChatSystemService) CreateMessage(id string, number int32) (*model.CreateMessageResponse, error) {
	backgroundContext := context.Background()

	res, err := s.railsGrpcClient.CreateMessage(backgroundContext, &proto2.CreateMessageRequest{ApplicationToken: id, ChatNumber: number})
	if err != nil {
		s.logger.Errorf("CreateMessage: error while invoking command %v", err)
		return nil, err
	}

	return &model.CreateMessageResponse{
		Message: model.Message{
			Number: res.Number,
			Body:   res.Body,
		}}, nil
}

func (s *ChatSystemService) GetMessages(id string, number int32) (*model.GetMessagesResponse, error) {
	backgroundContext := context.Background()

	res, err := s.railsGrpcClient.GetMessages(backgroundContext, &proto2.GetMessagesRequest{ApplicationToken: id, ChatNumber: number})
	if err != nil {
		s.logger.Errorf("GetMessages: error while invoking command %v", err)
		return nil, err
	}

	messages := make([]model.Message, 0)
	for _, message := range res.Messages {
		messages = append(messages, model.Message{
			Number: message.Number,
			Body:   message.Body,
		})
	}

	return &model.GetMessagesResponse{
		Messages: messages,
	}, nil
}

func (s *ChatSystemService) SearchMessages(id string, number int32, req *model.SearchMessagesRequest) (*model.SearchMessagesResponse, error) {
	backgroundContext := context.Background()

	res, err := s.railsGrpcClient.SearchMessages(backgroundContext, &proto2.SearchMessagesRequest{ApplicationToken: id, ChatNumber: number, Query: req.Query})
	if err != nil {
		s.logger.Errorf("SearchMessages: error while invoking command %v", err)
		return nil, err
	}

	messages := make([]model.Message, 0)
	for _, message := range res.Messages {
		messages = append(messages, model.Message{
			Number: message.Number,
			Body:   message.Body,
		})
	}

	return &model.SearchMessagesResponse{
		Messages: messages,
	}, nil
}
