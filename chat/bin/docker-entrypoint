#!/bin/bash -e

run_safely() {
  "$@" || {
    echo "Command failed: $*"
    return 0
  }
}

# If running the rails server then create or migrate existing database
if [ "${1}" == "./bin/rails" ] && [ "${2}" == "server" ]; then
  ./bin/rails db:prepare
fi

run_safely bundle exec rake db:create
run_safely bundle exec rake db:migrate

exec "${@}"
