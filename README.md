# Chat System

This project implements a chat system with a Ruby on Rails backend using gRPC and a separate Go REST API that communicates with the Rails app.

## Technologies and Versions

- Ruby: 3.3.3
- Rails: 7.1.3
- Go: 1.22
- MySQL: 9.0.0
- Redis: 7.4
- Elasticsearch: 8.8.0
- Sidekiq: 7.0.3
- gRPC: 1.65.0
- Protocol Buffers: 3.17.3

## Components

1. Rails gRPC Server:
   - Handles core business logic
   - Uses MySQL for data storage
   - Implements Sidekiq for background job processing
   - Uses Elasticsearch for message searching

2. Go REST API:
   - Provides a RESTful interface for clients
   - Communicates with the Rails gRPC server

## Key Features

- Application management (create, update, retrieve)
- Chat creation and retrieval
- Message creation, retrieval, and search
- Asynchronous processing of chats and messages
- Race condition handling

## Setup

1. Clone the repository
2. Install dependencies:
   - run `make run` or run `docker network create shared_network && docker-compose -f chat/docker-compose.yml up -d && docker-compose -f chat_api/docker-compose.yml up -d`