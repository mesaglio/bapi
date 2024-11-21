require 'sinatra'
require 'sinatra/json'

set :port, 8080
set :bind, "0.0.0.0"

users = Array.new

helpers do
  def get_body
    JSON.parse(request.body.read) rescue nil
  end

  def valid_body?(body)
    body.has_key?("username") and body.has_key?("email")
  end

  def get_user_index(username, users)
    users.index { |usuario| usuario['username'] == username }
  end
end

get '/ping' do
  json('Pong')
end

get '/users' do
  json(users)
end

post '/users' do
  body = get_body
  if valid_body?(body)
    users.push(body)
    status 200
    json('Success')
  else
    status 400
  end
end

get '/users/:username' do |username|
  i = get_user_index(username, users)
  if i
    json(users[i])
  else
    status 404
  end
end

delete '/users/:username' do |username|
  i = get_user_index(username, users)
  if i
    users.delete_at(i)
  end
  status 200
end

patch '/users/:username' do |username|
  body = get_body
  i = get_user_index(username, users)
  unless i
    status 404
    return
  end
  if valid_body?(body)
    users.delete_at(i)
    users.push(body)
    status 200
    json('Success')
  else
    status 400
  end
end
