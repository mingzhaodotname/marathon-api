# convert from raml to swagger 2.0 

swagger 2.0 is also called Open API Spec, i.e. OAS

~/marathon/docs/docs/rest-api/public/api$ oas-raml-converter --from RAML10 --to OAS api.raml > api.swagger.json

# Generate client from swagger spec

java -jar swagger-codegen-cli-2.2.2.jar generate -i api.swagger.json -l go -o go-client

