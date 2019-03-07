#! /bin/bash

export consumerKey=`aws ssm get-parameter --name "consumerKey" --query 'Parameter.Value'`
export consumerSecret=`aws ssm get-parameter --name "consumerSecret" --query 'Parameter.Value'`
export token=`aws ssm get-parameter --name "token" --query 'Parameter.Value'`
export tokenSecret=`aws ssm get-parameter --name "tokenSecret" --query 'Parameter.Value'`

go build -o terraform-provider-twitter
terraform init