FROM ruby:3.3

WORKDIR /srv/jekyll
VOLUME /srv/jekyll

# Copy Gemfile and Gemfile.lock to the container
COPY Gemfile Gemfile.lock ./

# Install Bundler and gems from the Gemfile
RUN gem install bundler && bundle install

# Expose port 4000 for Jekyll
EXPOSE 4000

# Command to serve the Jekyll site
CMD ["bundle", "exec", "jekyll", "serve", "--host", "0.0.0.0", "--port", "4000", "--livereload", "--drafts", "--future"]

