class ChatApplicationMessage < ApplicationRecord
  include Elasticsearch::Model
  include Elasticsearch::Model::Callbacks

  belongs_to :chat_application_chat, counter_cache: true

  validates :number, presence: true, uniqueness: { scope: :chat_application_chat_id }
  validates :body, presence: true

  before_validation :set_number, on: :create
  after_commit :index_document, on: :create

  settings index: { number_of_shards: 1 } do
    mappings dynamic: 'false' do
      indexes :chat_application_chat_id, type: 'integer'
      indexes :body, type: 'text', analyzer: 'english'
    end
  end

  def as_indexed_json(options = {})
    self.as_json(
      only: [:chat_application_chat_id, :body]
    )
  end

  private

  def set_number
    self.number = (chat_application_chat.chat_application_messages.maximum(:number) || 0) + 1
  end

  def self.search(body, id)
    params = {
      query: {
        bool: {
          must: [
            {
              match: {
                chat_application_chat_id: id
              }
            }
          ],
          should: [
            {
              match: {
                body: body
              }
            }
          ]
        }
      }
    }

    self.__elasticsearch__.search(params).records.to_a
  end

  def index_document
    ElasticsearchIndexingJob.perform_later(self.id)
  end

end
