# ytplaylist-puller
Pulls information from YouTube APIs in the form of playlists

This app is intended to help you explore Google’s API offerings. Do not share your API Credentials with anyone, and do not copy data from your account into apps that you do not trust.

Currently utilizing V3 of Data API.

your client secret json goes in the '.credentials' directory (within your home directory)

# How to run
clone the repo to pull down the code.

### configure .env
`go get github.com/joho/godotenv` so we can use a .env file
then set the vars
`CHANNEL_ID=UC_x5XG1OV2P6uZZ5FSM9Ttw` replace with your channel ID at some point. This is the googleDev channelID which can be used for initial testing.

Currently the program is configured to write the list to a file. It will make the file if it isn't there, it will overwrite if run in the same day, and create a new one each day. This is for comparison over time.

### Scheduling 
TODO manual right now
#### Unix (Mac/Linux)
CRON maybe

#### Windows
The counterpart to cron jobs on Windows is the Scheduled Task, which is managed by the Task Scheduler application. To set up a scheduled task, you can: 
Click Create Task in the right-hand pane 
Enter a name and description for the task in the General tab 
Create a trigger for the task in the Triggers tab 
Create an action for the task in the Actions tab 
Customize the task in the Conditions and Settings tabs 
Click OK to save the task 


## Contributing 
If you're interested in contributing you can reach out with an issue or just make a PR. 


Limitations
apparently the 'watch later' playlist will not be listed in playlists. 
