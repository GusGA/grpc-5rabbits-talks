# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: imagexample.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "imagexample.ImageRequest" do
    optional :url, :string, 1
  end
  add_message "imagexample.ImageResponse" do
    optional :captcha, :string, 1
    optional :language, :string, 2
  end
end

module Imagexample
  ImageRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("imagexample.ImageRequest").msgclass
  ImageResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("imagexample.ImageResponse").msgclass
end
