# IMT2681 - Cloud Technologies - Assignment 2

---

## Service Requirements 

The service will allow a user to monitor a currency ticker, and notify a webhook upon a tick where certain conditions are met, such as the price falling below or going above a given threshold. The API must allow the user to specify the base currency (for simplicity, if the base is not EURO, your service should respond that this is not yet implemented), the target currency (arbitrary, from all currencies supported by [fixer.io](http://fixer.io/)), and the min and max price for the event to trigger the notification. The notification will be provided via a webhook specified by the user, and multiple webhooks should be provided (predefined types). 

 

In addition, the service will be able to monitor the currencies (all from the http://api.fixer.io/latest?base=EUR query) at regular time intervals (once a day) and store the results in a MongoDB database. The system will allow the user to query for the "latest" ticker of given currency pair between EUR/xxx(See Highlights), and also, to query for the "running average" of the last 3 days.  



Clarification: The service will only need to support the use of EUR as base currency. Other currencies could be supported, if you want, but are not required. If the user specifies base currency other than EURO, your service can respond: "not implemented", along with the corresponding status code.

---

## Service Specification

Base URL: http://<app_name>.herokuapps.com/

#### Registration of new webhook

New webhooks can be registered using POST requests with the following schema.

**Request:**

​	**Method:** `POST`

​	**Path:** `api/`

​	**Payload specification: **

```json
{
    "webhookURL":      {"type": "string"},
    "baseCurrency":    {"type": "string"},
    "targetCurrency":  {"type": "string"},
    "minTriggerValue": {"type": "number"}, 
    "maxTriggerValue": {"type": "number"}
}
```

​	**Example:**

```json
{
    "webhookURL":      "http://remoteUrl:8080/randomWebhookPath",
    "baseCurrency":    "EUR",
 	"targetCurrency":  "NOK",
    "minTriggerValue": 1.50, 
    "maxTriggerValue": 2.55
}
```

**Response:**

The response body contains the id of the created webhook, as string. Note, the response body will contain only the created id, as string, not the entire path; no json encoding. Response code upon success is _201 - Created_.

---

#### Invoking a registered webhook

When the service invokes a registered webhook, it uses following payload specification:

**Request:**

​	**Method:** `POST`

​	**URL:** `<webhookUrl>`

​	**Payload Specification:**

```json 
{
    "baseCurrency":    {"type": "string"},
 	"targetCurrency":  {"type": "string"},
    "currentRate":     {"type": "number"},
    "minTriggerValue": {"type": "number"},
    "maxTriggerValue": {"type": "number"}
}
```

​	**Example:**

```json
{
    "baseCurrency": 	"EUR",
 	"targetCurrency": 	"NOK",
    "currentRate": 		2.75,
    "minTriggerValue": 	1.50, 
    "maxTriggerValue":	2.55
}
```

**Response:**

Upon successful notification you will receive either status code 200 (for trigger) or 204 (when no trigger).

---

#### Accessing registered webhooks

Registered webhooks can be accessed with the webhook ID generated during registration**.**

**Request:**

​	**Method:** `GET`

​	**Path:** `/api/{id}` 

**Response:**

​	Upon entered invalid ID

Status _404 - Not Found_

​	Upon entered valid ID

Status _200 - OK_

​	**Body:**

```json
{
  "webhookURL":			"http://remoteUrl:8080/randomWebhookPath",
  "baseCurrency":		"EUR",
  "targetCurrency":		"NOK",
  "minTriggerValue":	1.50, 
  "maxTriggerValue":	2.55
}
```

---

#### Deleting registered webhooks 

Registered webhooks can also be deleted using the webhook id.

**Method:** `DELETE`

**Path:** `/api/{id}`

**Response:**

​	Upon deletion:

Status _202 - Accepted_

​	Upon failed deletion:

Status _404 - Not Found_

---

#### Retrieving the latest currency exchange rates

**Request:**

​	**Method:** `POST`

​	**Path:** `/api/latest`

​	**Payload Specification:**

```json 
{
    "baseCurrency":    {"type": "string"},
 	"targetCurrency":  {"type": "string"},
}
```

​	**Example:**

```json
{
    "baseCurrency": 	"USD",
 	"targetCurrency": 	"NZD",
}
```

**Response:** 

The response contains only the latest exchange rate value (no json tags). 

​	**Example:** 1.56

---

#### Retrieving the running average over the past three days 

**Request:**

​	**Method:** `POST`

​	**Path:** `/api/average`

​	**Payload Specification:**

```json 
{
    "baseCurrency":    {"type": "string"},
 	"targetCurrency":  {"type": "string"},
}
```

​	**Example:**

```json
{
    "baseCurrency": 	"USD",
 	"targetCurrency": 	"NZD",
}
```

**Response:** 

The response contains only the average (of the last three days) exchange rate value (no json tags). 

​	**Example:** 1.89

---

#### Triggering webhooks for testing purposes - FOR DEVS ONLY

 

This trigger invokes all webhooks (i.e. bypasses the timed trigger) and sends the payload as specified under 'Invoking a registered webhook'. This functionality is meant for testing and evaluation purposes.

**Request:**

​	**Method:** `GET`

​	**Path:** `/api/evaluationtrigger` 

**Response:**

If all invocations ran successfully

Status _200 - OK_ 

---

## Highlights

_Parts of the assignment I'm particularly happy about_

- Testing standards - taking use of the same type of testing generally throughout all of my code:

  I'm taking use of nested testing patterns. 
  First of all I've taken use of ` TestMain`

  ```go
  func TestMain(m *testing.M) {
  	// ---------- "global" setup --------------
    	//insert new database credentials to use a test databas e -> confining tests
    code := m.Run()
    	// ---------- "global" teardown -----------
    	// Drop database to clean up after tests've ran
    os.Exit(code)
  } 
  ```

  This is massively helpful and works like a charm when the database variable is global, as I can simply insert _other_ credentials to my database connector, further ensuring that my tests won't interfere with the actual service.
  The next nested testing pattern I've used is called [TDT](https://github.com/golang/go/wiki/TableDrivenTests) (Table Driven Testing) Which allows for grouping similar tests within one test to use the same setup/teardown, but still run them as individual tests. It even lists them in hierarchical order when running go test.
  This allowed me to test large pieces of code where some logic within said code will split in multiple cases. in other words: it allowed me to easily gain coverage by utilizing essentially the same test twice, thrice, etc. (n-ice?) with some minor changes to divert into the different cases.

- Allowing "All" Currencies as Base currency.

  I've taken use of a little hack to get results regardless of what you post as `baseCurrency`, The hack is nothing more advanced than a little piece of math. Though this is not totally representative of the actual conversion an end user would request (because of how money work), it's a proof of concept for now.
  As my system is obtaining currency data from fixer with baseCurrency = EUR, I simply do the calculation
  conversionRate = targetCurrencyRate/baseCurrencyRate

- Using multiple "users" when connecting with my MongoDB hosted on Mongo Atlas

  As I had trouble with understanding how to handle the sensitive data (Posted an issue on the issuetracker on this topic), I ended up essentially posting credentials to connect to the test database (as I wasn't able to load environment variables from the test environment, later I've found gotenv to be useful here), however I realized that _it is a major security flaw_ to allow anyone to connect to the mongo cluster, especially when the credentials would allow them to enter the actual database the service too would use. This is why I created a very basic user that was only allowed to read/write on the "test" database and it's one collection.

- High test coverage. If I've done my calculations right, I'm at 75% test coverage in total, almost the required coverage doubled.
|
- Dockerized project. just run `docker build --tag ass2:latest --file ./cmd/currencytrackr/Dockerfile.
`

## Other Noteworthy Mentions

_aka. Issues Revolving Around my Submission_ 

##### [#39](http://prod3.imt.hig.no/mariusz/imt2681/issues/39)

I, as many others started getting timeouts on my tests when verifying against the submission verification service, as testing some times tool fractionally more than 60 seconds to complete. I've now set it to use localhost as host for the daemon instead, and I'm getting test times closer to 2-8 seconds. 

##### [#51](http://prod3.imt.hig.no/mariusz/imt2681/issues/51)

I ended up in a situation where I was never able to get gofmt up to 100% with the verification service, although running gofmt (even after updating go) on my side never gave me any warnings, trying gofmt -d ./.. returned nothing at all. Even running gofmt -w ./.. did not fix this issue. In other words, this still persists, and I don't understand why.

##### [#48](http://prod3.imt.hig.no/mariusz/imt2681/issues/48) | [#57](http://prod3.imt.hig.no/mariusz/imt2681/issues/57) | [#39](http://prod3.imt.hig.no/mariusz/imt2681/issues/39) 

I touched on this briefly above, but feel like the discussion on sensitive data to connect to a mongoDB, or if we at all should connect to a mongoDB when testing. Is an honorable mention as we gained a lot of insight having it.

---

