require_relative './application_record'

class ChatApplication < ApplicationRecord
  has_many :chat_application_chats, dependent: :destroy

  validates :name, presence: true
  validates :token, presence: true, uniqueness: true

  before_validation :generate_token, on: :create

  private

  def generate_token
    self.token = SecureRandom.hex(16) while token.blank? || ChatApplication.exists?(token: token)
  end
end
