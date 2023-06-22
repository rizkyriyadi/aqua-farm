run auto migration first by cd to migration folder

then back to root repo and run main.go

here how to :

1. cd migration
2. go run migrate.go
3. cd ..
4. go run main.go

# Create Farm
localhost:3000/createfarm

{

  "farm_name": "farmName"
  
}



# Create Pond
localhost:3000/createpond
{

    "pond_name ": "pondName",
    
    "farm_id": "refer to farm"
}
