class ChatApplicationChat < ApplicationRecord
  belongs_to :chat_application, counter_cache: true
  has_many :chat_application_messages, dependent: :destroy

  validates :number, presence: true, uniqueness: { scope: :chat_application_id }

  before_validation :set_number, on: :create

  private

  def set_number
    self.number = (chat_application.chat_application_chats.maximum(:number) || 0) + 1
  end
end
