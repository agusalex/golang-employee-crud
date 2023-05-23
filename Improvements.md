# Improvements
## General
* Tests are definitely not gold standard, some are good, some are boilerplate, more time would have allowed me to craft a more robust testing suite
  * Happy path only this should be seriously improved, to handle corner cases
* Having separate endpoints and separate model and business logic for Contractors and Employees was seriously considered:
  * I don't foresee Employees and Contractors differing that much in the future
    * So Splitting them now would be overdesign
  * If we knew contractors or employees could have different qualities in the future then it would be a **must** starting this way
    * Splitting them later would mean updating the DB down the line
    * Also updating our endpoints and business logic
    * Quite costly, not at all trivial
* Went with MySQL in order to have easy support for Enums (Contractor/Employee), 
  * This locks us up of other solutions. Given the context of this assigment speed of delivery was a priority.
* Using the new version of Gorn would have been better but due to time constraints I went with using the version I already knew how to use
* Adding Auth?
## Model
* In line with separating Employees with Contractors, since Golang does not have inheritance:
  * I would have used a strategy pattern with interfaces for Member
      * I decided against this and went with having the two different fields be empty in the Member struct and a Type
      * ex. a Role interface type inside the Member struct and have an Employee Implementation and a Contractor implementation 
      * This would have presented the challenge of more complicated persistence layer logic
* Add an activated/deactivated field for lay off members, this way deleting is not necessary
# Business
* Also consider excluding deactivated members from search / members endpoints unless there's a param or header
* Swagger UI implementation
* Better logging system / pushing logs Splunk or Sumo
* New Relic reporting
* Having persistence/business layers separated with the current amount of complexity of the project is a bit overkill
  * But avoids refactors in the future and makes things a lot more clear
  * I think the go-way is running a bit leaner on the abstractions
  * That being said having handlers separated from services is 100% worth it even at this stage

## Deployment
* Add CD to our CI by pushing GCP secrets onto gitlab and changing the pipeline to push new image to GCP registry and update cluster
  * I was about to do this, but I don't have maintainer access :(
* Adding replica pods to the service and scaling if necessary to manage more throughput
* Same with DB adding more replicas to DB if needed.
* Having docker image separated in two layers one for tests and building and another for deployment is a better practice


## Side Notes
* I really liked this assigment, it was challenging but easy at the same time, right balance imo.