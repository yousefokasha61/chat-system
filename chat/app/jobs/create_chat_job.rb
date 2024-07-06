class CreateChatJob < ApplicationJob
  queue_as :default

  def perform(chat_id)
    chat = ChatApplicationChat.find(chat_id)
    application = chat.chat_application

    ChatApplication.transaction do
      application.lock!
      application.increment!(:chat_application_chats_count)
    end
  end
end