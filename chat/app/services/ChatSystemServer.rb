# frozen_string_literal: true
require 'grpc'
require_relative '../protobuf/chat_system_pb'
require_relative '../protobuf/chat_system_services_pb'
require_relative '../models/chat_application'

class ChatSystemServer < ChatSystemProto::ChatSystem::Service
  def create_application(create_application_req, _unused_call)
    app = ChatApplication.create!(name: create_application_req.name)
    ChatSystemProto::Application.new(token: app.token, name: app.name, chats_count: app.chat_application_chats_count)
  end

  def update_application(update_application_req, _unused_call)
    app = ChatApplication.find_by!(token: update_application_req.token)
    app.update!(name: update_application_req.name)
    ChatSystemProto::Application.new(token: app.token, name: app.name, chats_count: app.chat_application_chats_count)
  end

  def get_application(get_application_req, _unused_call)
    app = ChatApplication.find_by!(token: get_application_req.token)
    ChatSystemProto::Application.new(token: app.token, name: app.name, chats_count: app.chat_application_chats_count)
  end

  def create_chat(create_chat_req, _unused_call)
    app = ChatApplication.find_by!(token: create_chat_req.application_token)
    chat = app.chat_application_chats.create!
    CreateChatJob.perform_later(chat.id)
    ChatSystemProto::Chat.new(number: chat.number, application_token: app.token, messages_count: chat.chat_application_messages_count)
  end

  def get_chats(get_chats_req, _unused_call)
    app = ChatApplication.find_by!(token: get_chats_req.application_token)
    chats = app.chat_application_chats.map do |chat|
      ChatSystemProto::Chat.new(number: chat.number, application_token: app.token, messages_count: chat.chat_application_messages_count)
    end
    ChatSystemProto::ChatsList.new(chats: chats)
  end

  def create_message(create_message_req, _unused_call)
    app = ChatApplication.find_by!(token: create_message_req.application_token)
    chat = app.chat_application_chats.find_by!(number: create_message_req.chat_number)
    message = chat.chat_application_messages.create!(body: create_message_req.body)
    CreateMessageJob.perform_later(message.id)
    ChatSystemProto::Message.new(number: message.number, body: message.body)
  end

  def get_messages(get_messages_req, _unused_call)
    app = ChatApplication.find_by!(token: get_messages_req.application_token)
    chat = app.chat_application_chats.find_by!(number: get_messages_req.chat_number)
    messages = chat.chat_application_messages.map do |message|
      ChatSystemProto::Message.new(number: message.number, body: message.body)
    end
    ChatSystemProto::MessagesList.new(messages: messages)
  end

  def search_messages(search_messages_req, _unused_call)
    app = ChatApplication.find_by!(token: search_messages_req.application_token)
    chat = app.chat_application_chats.find_by!(number: search_messages_req.chat_number)
    results = ChatApplicationMessage.search(search_messages_req.query, chat.id)
    messages = results.map do |message|
      ChatSystemProto::Message.new(number: message.number, body: message.body)
    end
    ChatSystemProto::MessagesList.new(messages: messages)
  end
end
