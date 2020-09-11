require 'pinger_services_pb.rb'
require 'pinger_pb.rb'

class ApplicationController < ActionController::API
  def ping
    pinger_stub = Pinger::Pinger::Stub.new('localhost:5300', :this_channel_is_insecure)
    pong = pinger_stub.ping(Pinger::Reqest.new(text: params['msg']))
    render json: { text: pong.text, count: pong.count }
  end
end
