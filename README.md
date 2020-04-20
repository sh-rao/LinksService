# LinksService

# Prelude
Manages creation and retrieval of links per user via RESTful endoints:
- To create a link for a specified user
- To retrieve links for a user with an option of sorting the links by their created date 

# Design
**_Resources_**:
* User - represents a person/organisation that owns the links - /users/{id} where id is unique per user.
For simplicity purposes of this assignment this resource is not managed by this service. 
* Link - represents links owned by a user and this services supports three types of links:
  - Classic
  - ShowsList
  - MusicPlayer

With high cohesion and low coupling design principle, using DI(Dependency Injection technique)
  - `linkService` is injected to the `handler` so that appropriate `POST` or `GET` handler can create or retrieve links through `linkService`.
  - `handler` has got all the logic to transform the incoming request object to the following domain/business objects based on the **_linktype_**.
     - Classic Link
     - ShowsList
     - MusicPlayer

As there can be hundreds of these link types, the `LinkData` model has been designed in such a way that the actual data of the link is
of generic type so that requests for creating and retrieving links return just one model which is `LinkData` and based on the type of the link
this model can be interpreted (marshalled/unmarshalled) accordingly.

Data Model
----------
To support fast retrievals and scaling of the API service a NoSQL database (DynamoDB) has been chosen to persist the business/domain objects.
The model is a simple model designed with keeping in mind the access patterns mentioned in the requirement (basically YAGNI principle).
The data is retrieved by the API consumers by:
  - Retrieve all links for a given user
  - Retrieve all links for a given user sorted by link created date
With these two access patterns in mind `UserLinks` table has been modelled with the following attributes:
  - **_Primary Key Attributes_**
     - UserId: Primary key for storing a link of a specific type against this UserId
     - LinkDateCreated: Sort Key - so that the links can be sorted on this attribute for a given user
  - **_Other Attributes_**
     - LinkId - Unique id of the link (this can be made as Global Secondary Index if in the future there will be an access pattern to retrieve links by link ids).
     - LinkType - Type of the link  (this can be made as Global Secondary Index if in the future there will be an access pattern to retrieve links by link types).
     - LinkData - The actual content representing the real business data.

# Implementation Notes
- There are certain idiosyncrasies of Go language and hence variable names or code style will look and feel
  different from other languages. A quick read of this - https://golang.org/doc - may assist reading of the code.
- I am not a big believer is adding comments in the source code, the reason being source code itself should be
  self-explanatory. You won't see any comments in the source code unless there is a very good reason. :-)

# Assumptions
* For simplicity purposes user is not managed by this API.
* There is not much information about URLs and what data the links actual contain. All fields in `Music Player` has been made URIs.
  
### What could have been done better
I hate to use time as an excuse but given an extra few hours or so, I would improvise the current design and implementation with the following:

- The schemas that are constants in all the validators can be pulled into their own separate files and manged by a config service
- PACT can be used to do contract testing for requests coming in and responses going out so that the `validate` files can concentrate more on business validation instead of json attribute level validation.
- The handler should be agnostic about the actual link type and SHOULD NOT have to do a switch/case to marshal the link data.
  A factory can be used for this purpose by injected it to the handler.
- Logging and error handling can be improved by passing context and using context logger where ever necessary.
- Local SAM with DyanmoDB hosted locally would have solved the problem of persistence. This way `GetHandler` could have returned
  real data instead of dummy data.

# Prerequisites
- Make sure you have installed the latest version of Golang from https://golang.org/
  This service has been built and tested with go1.13 darwin/amd64 (on MacOS Mojave v10.14.6)
  in GoLand 2020.1 IDE.

# How to visualise the data model
- Download NoSQL Workbench from https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.settingup.html
- Open the app and import the Linktree data model `Linktree Data Model.json` stored under schema folder by following the instructions
  specified here - https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.Modeler.ImportExisting.html
- Data model can be visualised by following these instructions - https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.Visualizer.Facets.html

# How to run the service
- Clone this repository into your local folder using git. e.g.
  ~~~
  mkdir LinksService
  cd LinksService
  git clone https://github.com/sh-rao/LinksService.git
  ~~~
  
- From the project root folder (e.g. LinksService), run this command to download all the dependencies
  ~~~
  go get -u ./...
  ~~~
  
- Run the service from the project root folder (e.g. LinksService) by running the following command
  ~~~
  go run main.go
  ~~~
  
  # Running unit tests
  Units tests can be run by the following command from the project's root folder (e.g. LinksService)
  Due to date format issue `validate_shows_list` tests are failing.
  ~~~
  go test ./... -v
  ~~~
  

