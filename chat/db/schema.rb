# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[7.1].define(version: 2024_07_05_153227) do
  create_table "chat_application_chats", charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.bigint "chat_application_id", null: false
    t.integer "number"
    t.integer "chat_application_messages_count"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["chat_application_id"], name: "index_chat_application_chats_on_chat_application_id"
  end

  create_table "chat_application_messages", charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.bigint "chat_application_chat_id", null: false
    t.integer "number"
    t.text "body"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["chat_application_chat_id"], name: "index_chat_application_messages_on_chat_application_chat_id"
  end

  create_table "chat_applications", charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "name"
    t.string "token"
    t.integer "chat_application_chats_count"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["token"], name: "index_chat_applications_on_token", unique: true
  end

  add_foreign_key "chat_application_chats", "chat_applications"
  add_foreign_key "chat_application_messages", "chat_application_chats"
end
