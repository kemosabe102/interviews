#!/bin/bash
# This script will build the Docker container image for the application
appName='sensor-qc'

buildContainerImage() {
  # By default, the image tag used in the Helm deployment is the same as the value for `appVersion`
  local imageTag
  imageTag=$(grep 'appVersion' "./deploy/${appName}/Chart.yaml" | awk -F \" '{print $2}')

  docker build . -t "${appName}:${imageTag}"
}

main() {
  echo "Changing to application directory root"
  pushd "$(git rev-parse --show-toplevel || echo '.')/${appName}" || exit 1

  buildContainerImage

  popd || exit 1
}

main
exit 0
