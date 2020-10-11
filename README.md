
# financial-manager 

## Components 

- ui 
    - source code for the frontend application  
- server 
    - source code for the web api  
- filewatcher 
    - source code for the application that watches files 
- infra 
    - source code for anything that infrastructure related 


## Run the application 

### PreReqs 
- Go 1.13 or higher 
- latest node and npm 
- docker and docker-compose 


#### To run the application please follow these steps in order 
1. Spin up the local elastic search cluster 
    - ./set_max_map_count.sh 
    - `docker-compose up` 
2. When the elastic search cluster is up and running 
    - `make build-all`
3. in one terminal window run
    - `make run-server`
4. in another terminal window run 
    - `make run-filewatcher` 


### To populate the data into the elastic search cluster 
- ( Assuming all the above components are running ) drop the transaction csv files into the /data directory 
    - this will automatically detect the file was added and populate the data accordingly


### New Issue Template 

```
**[ Current ] -** 

- Details 

**[ Changes Needed ] -** 

- Details 

```


------

### Registering a processor

```
curl --header "Content-Type: application/json" --request POST --data '{"url":"http://localhost:9091","processor_name":"xyz"}' http://localhost:8080/processor

```