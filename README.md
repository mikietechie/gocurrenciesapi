# SAPP Rates for Gophers

## Introduction

A Fintech API. This is a backend server for fetching, storing, and serving currency and rates data. Examples of companies that provide this service are Microsoft, Google, Polygon and Beacon. My motivation for making this API is to prove how well I understand backend development, how much I know about APIs, Security, Caching, Optimazation and Documentation. I also wish to present it as a portfolio project since I am trying to move from being a fullstack Python JavaScript developer to a "Real" backend developer who is worth paying thousands of dollars.

Gophers is the name given to Golang Developers. These developers generally look happier, healthier and wealtier than other developers. (Why because Go is simple unlike Rust, well compensated unlike me, has typing unlike Python, good errors style unlike JavaScript and compiles fast unlike Java which takes five a tenth of a developers life.)

## Technologies Used

Programming language: Golang
The system is built using the Go programming language. Golang is a modern, super fast, statically typed programming  language developed by Google. It is used in building Cloud Infrastructure, High Performance Web Servers and Developer Tools. Golang powers some of our favourite tools like Docker and Traefik. I chose Golang because I wanted to make the API as fast as possible and also use something that is not Python or JavaScript.

Web Framework: Gin
Golang provides a decent internal HTTP library that can be used for web server development, but I chose to use Gin (Go-Gin). Gin is a blazing fast Golang backend web framework. Me coming from Python and JavaScript it feels like Express JS or Fast API. It has all the features of a modern simplistic, unopinionated web framework. It supports data binding, easy serialization, crash handling, middlewares e.t.c . With Gin, the choices of authentication, database, templating and caching are left to the developer. I chose Gin because it feels like Express JS or Flask, (honestly I should have chosen Fiber)

Database: Postgres
Postgres is an open source Relational Database Management System. Postgres is quite popular in the Python community where I come from. It has advanced features like text search, json support and multi schema support. SQL Relational databases are greate because querying data using SQL is simple and clean as compared to their NoSQL counterparts. I chose Postgres because I am used to working with Postgres and it looks like the industry standard as well. Gone are the days of MySQL and MS Server.

Document Storage: Mongo
MongoDB is an open source NoSQL database. It is actually the most popular. Personally I do not like NoSQL databases, but they are fast and efficient when it comes to storing large blobs of data because of the way they are built, they are also easier to scale should the need arise, plus I needed to prove that yes I do know and understand what a collection and a document is.

Caching: Redis
Redis is an open source cache software. Redis stores data in key value pairs. You can read and write data to Redis way faster than you would to MongoDB or Postgres. Redis will give your Server massive perfomance gains. I chose Redis because it is easy to setup and manage.

API Documentation: Swagger.
Swagger is a technology for documenting APIs. Using swagger I can export client libraries for my API, making it easy for users to consume my API using their favourite languages, Python, TypeScript e.t.c . I chose Swagger over alternatives like Redoc, because Golang has good support for Swagger and I have also used Swagger with Fast API

Data Source: Beacon Currency API
Beacon is a fintech company providing data services through an API. I chose Beacon because it is the first fintech data service I could access from the beautiful city of St Petersburg Russia.

Authentication: JWT and API Key
JWT (JSON Web Token) is an authentication protocol where data is encoded by a key and passed around. API keys are basically cryptic ID's which client apps use, I will explain later why I had to use both JWT's and API Keys.

Testing
TODO

## Architecture

The system has 3 data models User, Client and Rate.

### Workflow

Users signup and have the role of clients. On passing Signup a client is created for the user. The client has an API Key and stores the number of API hits which is decremented for every hit and can be replenished when user buys some hits. Our service exposes different endpoints to client softwares.

The client can:

1. Fetch prevailing exchange rates for all internationally recognized currencies national and crypto.
2. Fetch a list of all currency names
3. Convert from one currency to another a given amount.
4. Query the rate at a given point in time.
5. Query the time series data for selected currencies in a given tim period

From an Admin Dashboard Users with the admin role can add more manage users, clients and client top ups.

### Configuration

Me coming from Python Django, the app has a config file that loads environment variables from the environment and an optional .env file. Here we set postgres, mongo, redis connection parameters, server environment, hashing secret key e.t.c .

### Data Storage and Relationships

A user is used for authentication. A client is used for tracking API usage. A client points to a user. We could have added payments to point to Clients. Users and Clients are stored in Postgres.

Rates are stored in Mongo DB. These are fetched from an external API and stored for future historical queries. They are stored in Mongo because Mongo excels at these type of things storing data blobs for frequent reads.

### Authentication

Users can register and they are auto assigned the role of Client. A client is created for the user. Users login using email and password then they get their user details and a JWT token. This token is used for accessing the user dashboard. The user can access and alter this client from the user dashboard. Here they can also check how many hits  they have left. User client softwares access the API using an API key which is passed as a query parameter.

Admins have the user role Admin. These users manage the system, and they can create more admins.

For more technical details please go through the swagger API.

## Options for future development

1. Better use of Golang Contexts
2. Micro Service Style (We could have created seperate services for managing clients, auth, and the expensies API). Honestly I hate the idea of Micro services I prefare modular monoliths for smaller projects like this one.
3. Use of the builtin Golang HTTP library. The Go community loves using built in libraries only. They hate using 3rd party libraries. Personally I wasn't going to reinvent the wheel.
4. Use of Raw SQL in queries. I love SQL, I enjoy writing SQL but I doubt there were any complex ORM operations that needed me to write Raw SQL. Fun fact I used to maintain an ERP for a multi million dollar company writen using Python, Flask and no ORM.
5. Code clean ups and refactoring. I wish to set Database connections, Redis connections and MongoDb connections to the drivers package. I do not know where I really need that initialize folder.
6. We could have used a faster programming language called Rust or Zig, or even C. Honestly I think languages like Rust, C, C++ and Erlang are too complex for a guy like me. Go is the lowest I will Go. (Before I dropped out of college when I was forced to learn PHP after learning C).
7. Clients and Users pagination.
8. We need more security like rate limiting for DDOS protection.

## Challenges faced

1. Being new to Go I had a hard time choosing which HTTP web frame work to use between Gin and Fiber. I guess I spent two days making that descision. Gin is more popular and mature, whilst Fiber is like a better version of Gin. It's like choosing between Flask and FastAPI.
2. For some weird reason I had a hard time connecting to Postgres. before you shout at me I have been connecting to Postgres using Python and JavaScript for about 5 years. I first wrote the code on my own, it failed, then copied it from the official docs, guess what it failed again. It turns out the problem was with my pre-existing Docker databases.
3. Structuring the code. Unlike Django, Angular and Nest JS. Golang Gin does not give you a predifined way to structure the project. In my journey to learn Go I met way to many structures and I ended up confused. I tried to take the best from all of them.
4. Lack of time, I did not have enough time per day to dedicate to this project. I work 8 hours a day as a Python Backend Developer. (I work on peripheral services only, because I cannot understand properly the language used in my city of residance)
5. *I had no one to give me feedback as I was developing this system. I have 4 years of experience in developing production grade Python JavaScript software, but I have always worked under the watch of an alpha male, a senior software engineer, a 10x developer, a seasoned developer. Truth be told the people who hold these titles have invaluable experience, they know everything, they have seen everything.*
