class CreateChatApplications < ActiveRecord::Migration[7.1]
  def change
    create_table :chat_applications do |t|
      t.string :name
      t.string :token
      t.integer :chat_application_chats_count

      t.timestamps
    end
    add_index :chat_applications, :token, unique: true
  end
end
