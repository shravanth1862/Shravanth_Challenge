aws cloudformation deploy \
    --stack-name StaticWebStack \
    --template-file template.yaml \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides \
        StackName=StaticWebStack \
        KeyPair=dummytestpair \
        amiID=ami-093e79ec1e354151c \
        CertArn=arn:aws:acm:us-east-1:890418425978:certificate/8441b711-2d9f-49e8-85f9-ff3137f86a11