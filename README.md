# WebHook with Golang in Docker

A simple WebHook server using golang, dockerized for easy execution in any environment.

## Description

This is a basic webHook server implemented with golang. 
### Changes:
- Now, the server serves an main.go file for easy WebSocket client testing through a browser.

## Technologies Used

- Golang
- WebHook
- Docker

## Prerequisites

To run this project, you need to have Docker installed on your machine. If you don't have it, you can download it from [here](https://www.docker.com/products/docker-desktop).

## Instructions to Run the Project

1. *Clone this repository:*

   If you haven't cloned the repository yet, you can do so with the following command:

   bash
   git clone git clone git clone https://github.com/SantiagoZ98/Go_Webhook.git

2. **Build the Docker image:**

   Before running the container, build the Docker image using the following command:

   bash
   docker build -t santiagozurita26/my-go-webhook .

3. *Push the image to Docker Hub (if needed):*

   If you'd like to upload the image to Docker Hub for others to use, you can do so with:

   bash
   docker push santiagozurita26/my-go-webhook:tagname

4. **Run the Docker container:**

   After building the image, run the container with the following command:

   bash
   docker pull santiagozurita26/my-go-webhook
