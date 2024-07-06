# frozen_string_literal: true

class ElasticsearchIndexingJob < ApplicationJob
  queue_as :elasticsearch

  def perform(message_id)
    message = ChatApplicationMessage.find(message_id)
    message.__elasticsearch__.index_document
  end
end
