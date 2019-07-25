#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

cat <<EOF
STABLE_KUBE_NAMESPACE ${NAMESPACE}
EOF