#!/bin/bash

set -eux

if [ $# -ne 1 ]; then
  echo "Usage: $0 <STACK_NAME>"
  exit 1
fi

STACK_NAME=$1

# Determine region: first try aws configure get region, then fallback to AWS_REGION environment variable
REGION=$(aws configure get region || echo "")
if [ -z "${REGION}" ]; then
  if [ -z "${AWS_REGION}" ]; then
    echo "Error: AWS_REGION environment variable is not set and no region is configured in AWS CLI." >&2
    exit 1
  fi
  REGION=${AWS_REGION}
fi

# Function to resolve SSM parameters
resolve_ssm_parameter() {
  local param_value=$1
  if [[ "$param_value" == "{{resolve:ssm:"* ]]; then
    local ssm_param_name=$(echo "$param_value" | sed -e 's/^{{resolve:ssm://' -e 's/}}$//')
    echo "$(aws ssm get-parameter --region ${REGION} --name "${ssm_param_name}" --query Parameter.Value --output text)"
  else
    echo "${param_value}"
  fi
}

# Get UserPoolClientId from stack outputs
USER_POOL_CLIENT_ID=$(aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Outputs[?OutputKey=='AdminUserPoolClientId'].OutputValue" \
  --output text)

if [ -z "${USER_POOL_CLIENT_ID}" ]; then
  echo "Error: UserPoolClientId not found in stack outputs for stack: ${STACK_NAME}" >&2
  exit 1
fi

# Get CallbackURL and LogoutURL from stack parameters
AUTH_REDIRECT_URL=$(aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='CallbackURL'].ParameterValue" \
  --output text)

LOGOUT_REDIRECT_URL=$(aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='LogoutURL'].ParameterValue" \
  --output text)

# Get DomainName from stack parameters
IDP=$(aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name "${STACK_NAME}" \
  --query "Stacks[0].Parameters[?ParameterKey=='DomainName'].ParameterValue" \
  --output text)

CONFIG_FILE="$(dirname "$0")"/../static/config.json

cat << EOF > "${CONFIG_FILE}"
{
  "UserPoolClientId": "${USER_POOL_CLIENT_ID}",
  "AuthRedirectURL": "${AUTH_REDIRECT_URL}",
  "LogoutRedirectURL": "${LOGOUT_REDIRECT_URL}",
  "IDP": "${IDP}"
}
EOF

echo "Generated ${CONFIG_FILE}"
