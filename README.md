# CS2 Demo Parse (GO)
Modified version of `print_events.go` from the [demoinfocs-golang library](https://github.com/markus-wa/demoinfocs-golang). Modifications include printing only certain kills, pulling headshot, wallbang, and time information from each kill, as well as handling of the team swap initiated at each CS2 match's halftime. It then outputs all resulting data into two CSV files, `data.csv` and `data2.csv` with all information split between kills taking place before and after halftime. 

## Retrieving CS2 Demo File
To retrieve your demo file from previous Counter-Strike 2 Competitive matches:
1. Visit your Steam profile (in the Steam client, mouse over your name on the large header text and click "Profile" within that menu.
2. On the right side of your profile page, click "Games".
3. Search for Counter-Strike 2 within the resulting list
4. Click "My Game Stats", then within the dropdown click "username's Personal Game Data".
5. From here, click any tabs your game fell under (this parser specifically focuses on Ranked Competitive Matches and Premier Matches only).
6. Scroll through your historical matches and upon one of your choosing, click the "Download Replay" button.
7. Using your favorite archive software, unzip the downloaded archive and store the ".dem" file in a directory of your choosing.

## Installation
In order to run this parser (instructions borrowed from demoinfocs): 
1. Download and install the latest version of Go [from golang.org](https://golang.org/dl/) or via your favourite package manager

2. Create a new Go Modules project

```terminal
mkdir my-project
cd my-project
go mod init my-project
go get -u github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs
```

3. Place `parse_demo.go` within your `my-project` directory.

4. Download your demo file from Valve with the above instructions.

5. Run `go run parse_demo.go -demo [path to your demo file]`

The program should print out all kills occuring within your match, then outputting the two CSV files with results. 
