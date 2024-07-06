#!/bin/sh

set -e

#Remove pre-existing server.pid for rails
if [ -f tmp/pids/server.pid ]; then
  rm tmp/pids/server.pid
fi


bundle exec sidekiq