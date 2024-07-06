class CreateMessageJob < ApplicationJob
  queue_as :default

  def perform(message_id)
    message = ChatApplicationMessage.find(message_id)
    chat = message.chat_application_chat

    ChatApplicationChat.transaction do
      chat.lock!
      chat.increment!(:chat_application_messages_count)
    end
  end
end