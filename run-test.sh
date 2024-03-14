STACK_NAME="StaticWebStack"
LAMBDA_EXPORT_NAME="LambdaArn"
WEBSITE_URL_EXPORT_NAME="WebsiteURL"

echo -e "Lambda response\n"
LAMBDA_ARN=$(aws cloudformation describe-stacks --stack-name $STACK_NAME --query "Stacks[0].Outputs[?ExportName=='$LAMBDA_EXPORT_NAME'].OutputValue" --output text)

aws lambda invoke --function-name $LAMBDA_ARN output.txt
echo -e "\nChecking instance status: "
cat output.txt

WEBSITE_URL=$(aws cloudformation describe-stacks --stack-name $STACK_NAME --query "Stacks[0].Outputs[?ExportName=='$WEBSITE_URL_EXPORT_NAME'].OutputValue" --output text)
echo -e "\n\nSatic website status: "
echo -e "$WEBSITE_URL\n"

curl_output=$(curl -s -I -w "%{http_code}\n" "$WEBSITE_URL")
http_code=$(echo "$curl_output" | awk 'NR==1')
echo "HTTP Status Code: $http_code"

read -n 1 -s
echo "Continuing..."

