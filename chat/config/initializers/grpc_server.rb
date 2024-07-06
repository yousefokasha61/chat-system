# config/initializers/grpc_server.rb

require 'grpc'
require_relative '../../app/protobuf/chat_system_pb'
require_relative '../../app/protobuf/chat_system_services_pb'

module GrpcServer
  def self.start
    addr = "0.0.0.0:50051"
    s = GRPC::RpcServer.new
    s.add_http2_port(addr, :this_port_is_insecure)
    puts "Adding ChatSystem::Server to gRPC server"
    s.handle(ChatSystemServer)
    puts "ChatSystem::Server added successfully"

    Thread.new do
      puts "Starting gRPC server on #{addr}"
      s.run
    end
  end
end

if ENV["RUN_GRPC_SERVER"]
  Rails.application.config.after_initialize do
    puts "Initializing gRPC server"
    GrpcServer.start
    puts "gRPC server initialized"
  end
end

