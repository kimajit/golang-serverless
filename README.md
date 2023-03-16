# golang-serverless-setup


Deploying lambda function using server-less frame work:

This setup will retrive the image-name and repo-name from Aws Ecr  and Trigger the circle ci pipeline 

### Prerequisites 
1) setup secret manger with name 

```
dev/AWS/Setup
```

* Add parmeter in secret manager 

```

circleCIToken : "Personal Acces Token "

vcs: github

owner : "circle ci account id "
```
* Circle Ci parameter 

```
projectName : "pipelinename"

branchName: "Circle Ci Branch name "

serviceName :"circle Ci service name "

servicePath : "circle ci pram "

url : https://circleci.com/api/v2/project/
```



### Getting Started 

1) we need to create service first

```
sls create --template aws-go-mod --name lambda
```

2) we need to replace exmaple code with our go code

3) we need to configure serverless.yaml file with iam role and all the policies

4) setup event pattern in Serverless yaml 

```
source : "aws.ecr"
detail-type :"ECR Image Scan"
```

5) Downlode Chocolatey  

6) Downlode make 

```
choco install  make 
```

7) To deploy the serverless.yaml file we need to deploy make file 

``` 
make deploy
```
