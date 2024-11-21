FROM ruby:3.3-slim

# nio4r dep need to be compiled
RUN apt update && apt install -y build-essential

RUN bundle config --global frozen 1
WORKDIR /app
COPY Gemfile Gemfile.lock ./
RUN gem install bundler:2.2.9

RUN bundle install

COPY main.rb .

CMD ["ruby","main.rb"]