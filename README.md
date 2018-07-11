# Widgets API

Deploys a Golang API to manage Widgets using Lambda, DynamoDB and API Gateway.

## Prerequisites

* [Golang](https://golang.org/dl/)
* [AWS account](https://aws.amazon.com/)

## Setting up the environment

The initial idea was to use [Terraform](http://terraform.io) to do the deploy of API, but how I'm learning about it, I don't had time do to that. Sorry :/
But if you follow all the steps below the API will work perfectly.

### Creating the Lambda function

* Go to the [Lambda service](https://console.aws.amazon.com/lambda/home);
* Click in the **Create a function** button, give your function a name (`WidgetsAPI`)
* Choose **Go** as the Runtime;
* Create a new role with permission to manage your application: choose **Create new role from template**, give your role a name (`widgets-api-role`) and then choose **Simple Microservice permissions** as the Policy template;
* In the new page, you need to put these `Environment variables`, then click on `save` button at the top of the page:
    * REGION_AWS (Ex: `us-east-1`)
    * USERS_TABLE (Ex: `users`)
    * WIDGETS_TABLE (Ex: `widgets`)
    * JWT_TOKEN (Ex: `secret`)

### Creating a new DynamoDB table

* Go to the [DynamoDB service] (https://console.aws.amazon.com/dynamodb/home);
* You need to click in the `Create table` button to create the two tables that you put in the `Environment variables` of lambda configuration, and put `id` as our primary key.

### Creating a new API Gateway

* Go to the [API Gateway](https://console.aws.amazon.com/apigateway/home) service and click in the **Get Started** button (If you already have a API Gateway you click in the button **+ Create API**);
* Choose **New API** in the radio buttons and give you API a name (`WidgetsAPI`), then click in the button **Create API**.

#### Create User Resources
* In the **Resources** page, click in the resource that was already created for you (`/`), click in **Actions** and select **Create Resource**.
* Give your resource a name (`Users`) and specify it's path name (`users`).
* Now click in your `/users` resource and repeat the process: click in the **Actions** button, give it a name (**SingleUser**) and a path name (`{id}`, it's only these part, because it's a child of `/users`, so our resource path is going to be `/users/{id}`).
* Now, let's create the methods: click in the `/users` resource, click in the **Actions** button, select **Create Method** and choose `GET`. in the method setup, choose **Lambda Function** as the Integration type, check the **Use Lambda Proxy integration** checkbox, choose the Region where your Lambda is and enter its name (`WidgetsAPI`). Repeat this process to create a `GET` method to `/users/{id}`.

#### Create Widgets Resources
* In the **Resources** page, click in the resource that was already created for you (`/`), click in **Actions** and select **Create Resource**.
* Give your resource a name (`Widgets`) and specify it's path name (`widgets`).
* Now click in your `/widgets` resource and repeat the process: click in the **Actions** button, give it a name (**SingleWidget**) and a path name (`{id}`, it's only these part, because it's a child of `/widgets`, so our resource path is going to be `/widgets/{id}`).
* Now, let's create the methods: click in the `/widgets` resource, click in the **Actions** button, select **Create Method** and choose `GET`. in the method setup, choose **Lambda Function** as the Integration type, check the **Use Lambda Proxy integration** checkbox, choose the Region where your Lambda is and enter its name (`WidgetsAPI`). Repeat this process to create a `POST` method to `/widgets` and (`GET`, `PUT`) methods to `/widgets/{id}`.

## Deploying to AWS

To deploy our application we need to follow some steps. First, we're going to send our application to our Lambda function:

1. Compile your code using:

```shell
# Compile the code to linux OS
$ GOOS=linux go build -o widgets-api widgets-api.go

# Compress your executable into a zip file
$ zip widgets-api.zip widgets-api
```

2. Go to your Lambda function at your AWS Console.
3. In the **Function code** section, upload your `.zip` file and change your Handler to `widgets-api`.
4. Click in the **Save** button at the top of the page.

5. Now we can test our API: go to the API Gateway we created earlier, select the route `GET /users` and click in the **TEST** icon and a response with an empty array will be returned.

To all of our tests you need to put the `Authorization` header before do the test (`Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M`)

To have a valid resopnse in this endpoint you need to put these information in the `USERS_TABLE` in DynamoDB:
```
{
    "name": "John Doe",
    "id": 1,
    "gravatar": "http://www.gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterid"
}
```

Repeat the step `5` and you'll receive a list with this user.

To create some widget in the database you can select the `POST /widgets` resource, click in the **TEST** icon and fill the Request body with:
```
{
    "name": "My Widget",
    "color": "purple",
    "price": "23,49",
    "melts": true,
    "inventory": "50"
}
```

Execute the test and you may receive the Created response. Now if you do the request to `GET widgets` you'll receive a list with these widget.

To edit a widget you can select `PUT {id}` resource, click in the **TEST** icon and fill these fields:
* Request body:
    ```
    {
        "name": "My Widget",
        "color": "purple",
        "price": "23,49",
        "melts": true,
        "inventory": "50"
    }
    ```
* {id} with the ID of the widget that you receive in `GET /widgets` resource:
    ```
    cfc3a148-ab5e-4b51-a5bf-a69f125fbdfc
    ```

After the request, the widget update will be returned in the response.

If all tests are ok you can deploy the API. Click in the **Actions** button, select **Deploy API**, create a new stage (EX: `production`) and you'll get the **Invoke URL** that you'll use to call the API.
Now you can use `<invoke-url>/<endpoint>` to call the API from all places.

## Documentation

To see the API documentation you need to follow these steps:

1. Get Swagger client
```
$ go get github.com/yvasiyarov/swagger
```

2. Run Swagger
```
swagger -apiPackage="github.com/henriqueholanda/widgets-api" -mainApiFile="widgets-api.go" -format=markdown -output=./doc.md
```
