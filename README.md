# leader-board
Basic leaderboard for games. Create new leaderboard for your game on a web UI and use the http REST API, secured via API key authorization header.
# Getting started
## Building the project
Copy *.env-sample* to *.env* and fill in the environment variables needed. 
Same process for *.dbenv-sample* to *.db.env*.
In your terminal, go to the root of this project and tappe the command below.
```bash
docker-compose up --build
```
## Usage
### Web UI
Go to your web browser at [localhost:3000](http://localhost:3000) then login using the credentiale you create in the *.env* file. 
Create a new game and copy-paste the game's **API key** and **uuid**.
You can add or delete game through your web access.
### REST API
Every API request will need a **X-Api-Key** header with the provided API key. 
#### GET /api/v1/games/{game-uuid}/leaderboard/{top}
This request has two params.
* game-uuid: which is link to the API key (for authentication) and to get the correct data.
* top (int): which is optional, and is used if you need to get the top **X** of the game's leaderboard.
The response will look like this :
```json
// GET /api/v1/game/7e8r78e-rew89-bs9ds798/leaderboard/10
[
    { 
        "pseudo": "",
        "score": 0.0
    }
]
```
#### POST /api/v1/games/{game-uuid}/leaderboard
This request is used to post a new score into the game's leaderboard. The request's body needs to be :
```json
// POST /api/v1/game/7e8r78e-rew89-bs9ds798/leaderboard
{ 
    "pseudo": "",
    "score": 0.0
}
```
