# lib/tasks/compile_protos.rake
namespace :protos do
  desc "Compile proto files"
  task compile: :environment do
    proto_files = Dir[Rails.root.join('protos', '*.proto')]
    proto_files.each do |proto_file|
      output_dir = Rails.root.join('app', 'protobuf')
      FileUtils.mkdir_p(output_dir)
      system("grpc_tools_ruby_protoc -I #{Rails.root.join('protos')} --ruby_out=#{output_dir} --grpc_out=#{output_dir} #{proto_file}")
    end
  end
end
