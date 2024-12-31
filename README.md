This project consists of basic Golang programs and a step-by-step approach to building a RESTful API in Golang using MongoDB. Find below the steps and solutions to building such a framework in Go.

**Lesson 1 - Socket **

Exercise 
Create an ATM program using a socket-based approach. The code will be in 2 parts 
 
1. Server: This will be the brain of your code that will make decisions. This will take input from the client and based on value it will pass the message to the client.  
 
2. Client: It will take input from the user and pass the user input to the server, and it will display the server response to the user.  E.g It will ask the amount to be withdrawn from the user and pass the message to the server. Based on valid user input and the balance, the server will either give a message that the withdrawal is successful or will give a proper error message.

Link: https://github.com/Poonam-Sama/go-programming/tree/main/01codingquestions/ATMSOCKET

**Lesson 2 - REST  **

Exercise 
Convert Lesson 2 code to REST-based server and client. The client will pass the message as a POST request and the server will send the message with the proper status code. 200 for success and 400 in case of error. The rest of the logic will work the same way as coded in Lesson 1. 

Link: https://github.com/Poonam-Sama/go-programming/tree/main/01codingquestions/websocket/http

**Lesson 3 - Web Server **

Exercise 
 
Convert the program from Lesson 2 to an HTTP-based server. Now your code will have 2 features 
 
1. GET - to know the account balance  
2. POST -  to withdraw from the account
Link: https://github.com/Poonam-Sama/go-programming/blob/main/01codingquestions/webserver/webs.go

**Lesson 4 - Gin **

Exercise 
Convert the program from Lesson 3 to a gin- based code. 
Link: https://github.com/Poonam-Sama/go-programming/tree/main/04Module

**Lesson 5 - MongoDB**

Exercise 
 
Create a table for the user's bank account 

**Lesson 6 - mgo **

Exercise 
Handle the above table with Golang.

Link: https://github.com/Poonam-Sama/go-programming/tree/main/06Module

**Lesson 7 - Final **
About 
We will move the code to gin-based MongoDB solution, where the balance will be handled in the table and all transactions will be recorded in the table.

Exercise 
You have to create 2 tables 
 
1. For user account
2. Transaction 
 
When a user withdraws an amount that record should be recorded in the transaction table and the user balance will be updated in the user table.  The balance will be read from the user table.

Link: https://github.com/Poonam-Sama/go-programming/tree/main/07Module/final
