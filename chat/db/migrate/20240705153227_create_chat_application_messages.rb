class CreateChatApplicationMessages < ActiveRecord::Migration[7.1]
  def change
    create_table :chat_application_messages do |t|
      t.references :chat_application_chat, null: false, foreign_key: true
      t.integer :number
      t.text :body

      t.timestamps
    end
  end
end
