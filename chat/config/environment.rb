# Load the Rails application.
require_relative "application"

# Explicitly require ChatSystemServer
require_relative '../app/services/ChatSystemServer'

# Initialize the Rails application
Rails.application.initialize!

# Initialize the Rails application.