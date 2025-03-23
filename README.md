# Let's go wild!

This is an example app showcasing a simple backend API service written in Go.
⚠️ As of 2025, this attempts to follow a very "go-typical" web server structure ⚠️
The idea is to display some of the different design challenges when designing API resources - how do we group data? How do we expose it? What counts as "business logic" vs other things, for example.

For what it's worth, zero LLM/AI was used when implement this application 😘

# The case 🫎🌲⛰

The use case is an API on top of a wildlife database with information about plants and animals.

This is intended to be an API both with "BFF" (back end for front end) capabilities - i.e. serving a website, as well as endpoints which external users can call to learn more about wildlife.

Endpoints exist to list plants, animals and biome specific features.

# Stack
## Explanation
Gowild is written entirely in Go, using primarily standard library features, with the occasional beloved 3rd party library.
Besides go, there following exists:
- Dockerfile for containerizing the app
- Docker compose to spin up a local database on your machine
- Makefile to automate tedious command line actions
- .env + example file to include local environment variables
- Goose migrations: the SQL files in `migrations` are made to work with Goose, a popular database migration tool for go. Goose can be called to migrate both from the command line and from within the code, depending on if you want to migrate from within the application or independent from it.
## Tools you need
- Go (>version 1.24)
- Docker

# Understanding the go folder structure

## Names of subdirectories
- `/cmd` - This is where the main function(s) live. Typically just one, but might be more if you have scripts besides the server.
- `/internal` - This is where the main meat of the application lives. The "classes", or well, structs + interfaces responsible for doing business logic and making calls to the database live here.
- `/pkg` - Often called "lib" in other languages. This is where you put isolated bundles of functions which you treat like your own libraries, which you could in theory swap out. Good example for this is to have **database connections for a specific database type** in pkg, but then have your **database interfacing code** in internal in some type of db/repository package.
- `/migrations` - Here we put SQL files for whenever we want to change the database schema. The nice thing about this is you can step forward and ba

# Understanding the application in action

Here's a basic description of what happens when this application gets spun up:
- Startup:
    - The main function starts and begins "bootstrapping". During this phase, if anything goes wrong we let the whole app crash. Hence we have no "context", and we don't return errors anywhere.
    - Config gets read with info about how to exist in the current environment. Which ports to use, how to connect to resources etc etc.
    - The app connects to its remote resources (i.e. database, in the future maybe a cache)
    - We create our long living structs; services and their repositories. Their methods are registered as handlers for http calls
    - We tell the server to continue listening to traffic. Easy peasy!
- Thereafter:
    - The server listens to incoming calls and automatically handles the threading. From here on we can worry about what happens in our handler functions.

# Design patterns to look into
In no particular order:
 - Service/repository pattern
 - Dependency injection
 - SQL migrations