# jobdir

## Live Preview
🔗This REST API is accessible at https://job-directory-n9ak.onrender.com

Try hitting `GET https://job-directory-n9ak.onrender.com/jobs`. Or check the following documentation to see the accessible routes. 

## For Employer
If you have trouble accessing or evaluating this personal project of mine, kindly contact me at
✉️: daffaalexander.work@gmail.com, CC: daffaalex22@gmail.com

## Documentation
📃Access the following URI to view list of possible routes to access:
<br>
https://app.swaggerhub.com/apis-docs/daffaalex22/jobdir/3.0

## Description
This project is created during the [Kemendikbud's Kampus Merdeka Program](https://kampusmerdeka.kemdikbud.go.id/) that the owner do on since August 2021. [Alterra Academy](https://academy.alterra.id/) is where the owner do the Independent Study (SI). Other than following a 3 hours daily live session, 3 hours daily self-learning and a daily tasks, participants are instructed to develop a RESTful API implementing all of the material that have been studied. 

In short, the RESTful API is about a Job Seeker Directory where users can have access to search jobs and directly apply. The project is written in Go Programming Language and implements the Clean Architecture Code Structure.

Some of the technologies and libraries implemented on the projects are as follows:
* [Echo Labstack](echo.labstack.com): A Golang Web Framework
* [GORM](https://gorm.io/): A library for implementing ORM
* [Viper](https://github.com/spf13/viper): Used for environment configurations
* [vektra/mockery](https://github.com/vektra/mockery): Used for creating mocks for unit testing 
* [stretchr/testify](https://github.com/stretchr/testify): Used for unit testing
* [codegangsta/gin](https://github.com/codegangsta/gin): A live reload utility

The database management system used are both SQL and NoSQL:
* MySQL: The main database used are relational database, this is where the API stores the entity
* mongoDB (hasn't deployed): Used for storing log messages from echo logger

The deployment of the project uses:
* Docker; and 
* AWS EC2 
* AWS RDS

Continuous Integration and Continuous Deployment (CI/CD) using GitHub Actions is implemented to automate the deployment process.

## Clean Architecture
As previously mentioned, the project implements Clean Architecture. The four layers on the project are:
  * Domain Layer
  * Repository Layer
  * Usecase Layer
  * Controller Layer
 
### The Diagram
![golang clean architecture](https://github.com/daffaalex22/jobdir/raw/main/CleanArch.png)

From the picture above, the four rounded rectangular corresponds to each Clean Architecture layer. The slightly bolder arrow pointing from certain layer to another, pictures the dependency of the layer. For example, the domain layer (red) is pointing to repository layer (purple). This means that the repository layer imports package from domain and thus the repository layer depends on the domain layer.
