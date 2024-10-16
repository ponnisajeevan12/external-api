# Objective: To interact with external APIs - Twitter API

The aim of this project is to generate Twitter API keys from developer portal, use these keys to authenticate and then perform tasks such as 'Create a Tweet' & 'Delete a Tweet' using Twitter APIs. 

### 1. How to setup Twitter Developer Account and Generate API Keys ?

- **Create Twitter Account and Twitter Developer Account:**

1) Goto https://x.com/i/flow/signup and key in the required details to create a Twitter Account.
2) Navigate to the developer portal available in the URL https://developer.x.com/en/portal/dashboard
3) Go with the Free plan for testing purposes, then proceed to create a project & app. The free subscription has a limit of 1 App per account. So, we can rename the default project and app name or even delete the default and proceed to create a new app following the on-screen instructions.

![DEV-Portal](https://github.com/ponnisajeevan12/external-api/blob/master/images/dev.png)
![DEV-Plan](https://github.com/ponnisajeevan12/external-api/blob/master/images/plan.png)

- **Generate API Keys:**

1) We have to create API Keys and Tokens for the App now.
2) To do this, click on "Projects & Apps" under your Developer Portal followed by click on the Project name and then the 'key' icon next to your Apps.

![DEV-Token](https://github.com/ponnisajeevan12/external-api/blob/master/images/Project.png)

3) Now, generate the below keys and save those to have it referred on your code.

_a) Consumer Keys - API Keys & Secret:_

This includes API Key and Secret. These are similar to username and password of an application and can be used to authenticate against the APIs.

![DEV-Token](https://github.com/ponnisajeevan12/external-api/blob/master/images/Token.png)

_b) Authentication Tokens - Bearer Token, Access Token and Secret:_

This includes Bearer Token, Access Token and Secret. Bearer Token authenticates the requests on behalf of your developer App. Access Token and Secrets are user-specific credentials which are used to authenticate OAuth 1.0a API requests.

- **Configure OAuth:**

Click on App Settings followed by User Authentication Settings and click on "Edit". Enable Read and Write Permissions for OAuth1.0a Authentication and then set callback url if you want to perform a local testing.

![OAuth](https://github.com/ponnisajeevan12/external-api/blob/master/images/perm.png)


### 2. How does the code code authenticate and communicate with the Twitter APIs to create and delete a Tweet?

1) Initialise go module and then install necessary dependencies for OAuth Authentication.

```
go get github.com/dghubble/oauth1
go get github.com/dghubble/go-twitter/twitter
```

2) We have created an .env file to keep the access keys and token safely.
3) When we run our main.go using 'go run main.go', it creates a tweet after authenticating against the endpoint for POST which is https://api.twitter.com/2/tweets.

![Tweet-EndPoints](https://github.com/ponnisajeevan12/external-api/blob/master/images/EndPoints.png)

4) Then OAuth creds are already referencing the .env file as per the code in main.go.
5) Once a tweet is created, it sleeps for 15 seconds and then deletes this same tweet using the delete function that we have created.

![Tweet-Created](https://github.com/ponnisajeevan12/external-api/blob/master/images/Tweets.png)
![Tweet-Deleted](https://github.com/ponnisajeevan12/external-api/blob/master/images/Create-Delete.png)

### 3: How does the code handles Authentication and Errors?

- **Authentication:**

The code uses OAuth1.0a credentials for authentication and this has been references in our go code.

- &#10060; IMPORTANT:   Always make sure NOT to push your valid secrets to GitHub. You can avoid this by adding the .env file to .gitignore file. 

- **Error Handling:**

Eror handling messages have been added to the code and few of the examples are as below.

_a) Failed to post tweet: 403 Forbidden & 401 UnAuthorised_

We will get these errors when we dont add the access tokens properly and so the authentication doesnt work properly.

![Error1](https://github.com/ponnisajeevan12/external-api/blob/master/images/Error.png)

_b) Error loading .env file_

We will get this error when the variable references are not proper in our code.

![Error2](https://github.com/ponnisajeevan12/external-api/blob/master/images/env.png)

# Conclusion: 

Here, we have learned how to communicate to Twitter API endpoints using Twitter API Token and secrets.

# References:

```
https://developer.x.com/en/docs/x-api/tweets/manage-tweets/introduction
https://developer.x.com/en/docs/platform-overview
```
