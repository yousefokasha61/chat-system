class CreateChatApplicationChats < ActiveRecord::Migration[7.1]
  def change
    create_table :chat_application_chats do |t|
      t.references :chat_application, null: false, foreign_key: true
      t.integer :number
      t.integer :chat_application_messages_count

      t.timestamps
    end
  end
end
