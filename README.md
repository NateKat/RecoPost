# Postal system
### About
This program handles delivery of parcel from different post offices (postal system).\
The program holds a number of cities, each with different amount of post offices.\
Each office has it's own limitations on parcel weight.\
Every office keeps the parcel in First In First Out order (FIFO).

## Installation
### Build
From within this git repo, run:
```
$ go build
```

### Instructions how to run
1. Input description of postal system
2. Input of actions to prefrom in the postal system

#### Input of postal system:
The description of the postal system begins with a line containing the number of cities in the country. Next, in sequence, separated by end-of-line, the descriptions of the cities.\
Each description of a city begins with a line containing the name of the city and the next line with the number of post offices in the city.\Next is a sequence, separated by end-of-line, is the descriptions of the post offices in the city.\
Each post office description begins with a row containing three numbers separated by spaces.\
The first number is the number of packages in the branch, the second number is the minimum package weight, the third number is the maximum package weight.\
Packages descriptions is next.\
Each package description consists of a single two-part row separated by a space. The first part contains a string that recognizes a package (name). The second part contains the weight of the package.\

#### Supported Actions/ operations:
Operation number one: Printing the contents of the city offices. There is a single parameter, the name of the city whose ffices you with to print.\
Operation number two: Transfer between offices. Has four parameters that appear in the following order: source city name, source office ID, destination city name, destination office ID.\
Operation number three: Printing the city with the most packages. Has no parameters.

#### Illegal input:
* Duplicate city name.
* Duplicate parcel name.
* Parcel weight must be within the limits of weight determined for the specific office it's created at.
* Minimum weight cannot exceed maximum weight in office creation.
* Parcel weight, number of cities, number of actions to preform must all be non negative?
* Action op number must exist, number of arguments must be in correlation to specific op.
* City name and office number must exist for each op that requires any of the above as input.

### Design and performance notes
#### Computational efficiency:
Different cities can be accessed via a mapping, by city name. (No need to go over all cities to find a specific city).\

Offices in city are accessed by offset, determined according to the office sequence number.\

Parcels (Packages) are stored in a list (FIFO) per office.\
Parcels are also saved in a global map during the initialisation process to allow efficient verification of unique parcel name. The mapping is deleted once the initialisation is done. 

#### Error handling:
Errors are propagated to main and printed to STDIN.\
Input error will cause the program to return with -1 exit code. \
There are no retries for illegal input, the program will exit with a proper message. \
Action/ operation format errors will cause the program to exit before any action is executed. \
On the other hand, action with parameters that doesn’t correlate with the data will cause the program to exit after the execution of previous valid actions. 
