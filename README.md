# Assignment-kirana Club
The assignment includes the APIs as per the given [requirements](https://kiranaclub.notion.site/Backend-Assignment-4d3c4465c9864422b1791e775c3b5272) and unit test.
### Following is a Flow Chart Diagram for the APIs
![flow-chart](https://github.com/harshitphoenix/Assignment-kiranaclub/assets/56081331/1a251c78-73c7-4b58-a3f7-2ca9fa3daebf)

### Following are the tables used in the assignment
![db-diagram](https://github.com/harshitphoenix/Assignment-kiranaclub/assets/56081331/1c3ba816-4642-4097-a9b9-6942180d14a8)

### API Description 
1. Create Jobs - 
   It takes the payload, which contains store_id, URLs, and visit time, and creates the jobs for each store_id. Respective Job ids are returned to the user as response and
Seperate Goroutines are fired for each job to process them parallely. Subsequently, the status of job is updated in the Job Table and Image info is added to the Meta Data table.

2. Get Job Status - 
   It takes job_id as input and returns the status of job (pending, failed, or completed) accordingly after validation of job id.
3. Get Visit Info - 
   It takes area_code, store_id, start  date and end date as input and returns all the processed images available inside the Meta Data Table.
   
### Scope of Improvements
1. Integration of queue service like Rabbit MQ or Kafka to make it scalable
2. Write more unit tests

