# NATS in Go

<h3>Pre-requisites</h3>
<li>Go SDK - to compile and run go code</li>
<li>Docker Desktop - to run the NATS Server (without downloading on host machine) </li>
<hr>

<h3>About the app</h3>
<li>In this application, I have tried to demonstrate the Pub-Sub Pattern in Go using NATS</li>
<li>We have an Order Service - which publishes events time to time</li>
<li>Email and SMS Service have subscribed to certain subjects</li>
<li>This app presents how an E-Commerce application must be using Messaging Systems between different microservices</li>
<hr>

<h3>To Run the App</h3>
<li>Go to terminal and start the NATS Server
  
```
  docker run -p 4222:4222 -p 8222:8222 -p 6222:6222 --name nats-server -ti nats:latest
```
-p 4222:4222: Maps the NATS client port (4222) to your host. <br>
-p 8222:8222: Maps the NATS monitoring port (8222) to your host. <br>
-p 6222:6222: Maps the NATS system port (6222) to your host. <br>
--name nats-server: Gives the container a name. <br>
-ti nats:latest: Runs the latest NATS server image interactively. <br>
</li>

<li> Go to SMS Service and Email Service and run the program
  
```
  go run main.go
```
Now subscribers are ready to listen to subscribed subjects
</li>

<li> Go to Order and run the program
  
```
  go run main.go
```
Now publisher will start publishing messages on subjects.
</li>
