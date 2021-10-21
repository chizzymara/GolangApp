https://mholt.github.io/json-to-go/

# Golang-Epic-Api

The application retrieves information about the current free games on the website https://www.epicgames.com/store/en-US/free-games and sends a message to slack periodically.

## _Implementation_
1. Creating go structures (struct) to match the structure of the free games promotions payload when a GET request is called.
2. Creating a function "Get free games" that makes a GET api call to the website to  retrieve the response payload, have a loop that goes through the elements and identifying items with the "discountPercentage" of zero, in the "discountSetting".
3. Identified games are added to the "FreeGames" array.
4.  Function  "Main" and "SendSlackNotification" are used to send the contents of the "FreeGames" array to "SLACK_URL"
