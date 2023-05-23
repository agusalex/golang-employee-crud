# Employee CRUD

### Highlights
* Built a CI pipeline un Gitlab to build the image (also runs tests)
  * Publishes the image to Gitlab artifact repository
* Deployed it to GCP using K8S
  * http://35.222.35.245:8080/api/v1/members
  * Added Ping and Health endpoints
* Added a Search endpoint for easy search by tags and type
  * [Search docs](Documentation.md#search-members)
  * http://35.222.35.245:8080/api/v1/search/members?type=CONTRACTOR&tags=Java
* Added docker-compose for easy local deployment
* DB Auto Migrate using Gorn
* [API Documentation](Documentation.md#member-crud-api)
  * [Swagger](docs/swagger.yaml)
* [Improvements docs](Improvements.md)
* 90% Test Coverage