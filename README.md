# supreme-spoon

## Aim
 The aim of this project is to provide functionality for movie watching. We want to implement finding and downloading torrents easily.
We plan on adding a Frontend in React to facilitate controlling the downloaded torrents. We also plan to add an API for controlling 
the app that will be running as a local service.
 We want to implement a microservice architecture for the various services we will be running. We also want to create a small Android app.

## Future Plans
 1. VLC support
 2. Web UI
 3. Android App
 4. Note taking during parts of a movie
 5. Letterboxd connection

# Installation
 You can build the package in the standard go way

        go build

# Usage
The CLI offers the following functionality:

## Search
We search for a torrent in the YTS api and return a prompt so the user can select a movie.

        ./supreme-spoon search [Movie Title]