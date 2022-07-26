#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# shellcheck source=../.env
source "$DIR/../.env" || {
  echo "error while sourcing env file"
  exit
}

cd "$DIR/.." || exit

if ! docker build -t "${PSQL_IMAGE}" -f "$DIR/Dockerfile_Psql" .; then
  echo "Could not build image ${PSQL_IMAGE}"
  exit 1
fi

PS3="Push image ${PSQL_IMAGE}: "

select opt in yes no; do

  case $opt in
  yes)
    docker push "${PSQL_IMAGE}"
    break
    ;;
  *)
    echo "NO PUSH"
    break
    ;;
  esac
done
