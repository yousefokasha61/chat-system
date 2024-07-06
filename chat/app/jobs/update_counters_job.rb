# frozen_string_literal: true
class UpdateCountersJob < ApplicationJob
  queue_as :default

  def perform
    ChatApplication.find_each do |application|
      application.update_column(:chat_application_chats_count, application.chat_application_chats.count)
    end

    Chat.find_each do |chat|
      chat.update_column(:chat_application_messages_count, chat.chat_application_messages.count)
    end
  end
end
