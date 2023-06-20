run auto migrate first, cd migrate then go run migrate.go
after that cd .. then you can run main file by using go run main.go

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
