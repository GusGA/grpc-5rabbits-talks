this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'imagexample_services_pb'
require 'rtesseract'
require 'mini_magick'
require 'open-uri'

# Don't allow downloaded files to be created as StringIO. Force a tempfile to be created.
OpenURI::Buffer.send :remove_const, 'StringMax' if OpenURI::Buffer.const_defined?('StringMax')
OpenURI::Buffer.const_set 'StringMax', 0

# GreeterServer is simple server that implements the Helloworld Greeter server.
class ImagexampleServer < Imagexample::ImageCaptchaService::Service
  # say_hello implements the SayHello rpc method.
  def resolve_captcha(image_req, _unused_call)
    puts 'Getting image'
    image_file = open(image_req.url)
    puts 'Resolving captcha'
    image = RTesseract.new(image_file.to_path, :processor => 'mini_magick')
    captcha = image.to_s.gsub(/\s+/,'')
    puts "Captcha resolve #{captcha}"
    Imagexample::ImageResponse.new(captcha: captcha, language: 'Ruby')
  end
end

# main starts an RpcServer that receives requests to GreeterServer at the sample
# server port.
def main
  port = 50051
  s = GRPC::RpcServer.new
  s.add_http2_port("0.0.0.0:#{port}", :this_port_is_insecure)
  puts "Listening on port #{port}"
  s.handle(ImagexampleServer)
  s.run_till_terminated
end

main