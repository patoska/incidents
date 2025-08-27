# Incidents Small App

## Quick start
You need docker compose in order to run this app using the following command:

```bash
docker compose up
```

## Tech stack
I used Go as backend with Gin and VueJS for frontend that consumes the Go API

## Tradeoffs
At first the app was kind of broken, it was only adding the title and it was not persisting the data.  
So the tradeoff was reaching a minimal solution that shows the front that communicates with the backend.

## What’s done vs. what you’d add with more time
With more time I would have added a proper database persistence and a bucket for the image  
Basically all bonuses listed in the assignment.

## AI usage
* Simple google searches like "Add CORS to golang gin app", "Use env variable in Vuejs", etc.
* Asked ChatGPT to to generate a simple vuejs app that meets some requirements and based in a package.json that I already had.
* Asked ChatGPT for errors debugging.
* Github Copilot for autocomplete.


## Time spent
14:59 - 15:40 (41 min)  
16:02 - 16:26 (24 min)  

At this time, the app index was only showing the title of the incidents this was due a backend bug. I I spent like ~30 min more to reach the actual state, ie, persisting the data and showing all parameters.

I did not count the time spent writing this README.me

## Screenshots
Create incident view:  
<img width="770" height="795" alt="image" src="https://github.com/user-attachments/assets/cb03cdd2-05d6-4266-aa74-8f93f5f57a76" />

Index view:  
<img width="858" height="692" alt="image" src="https://github.com/user-attachments/assets/0b2afdbe-3196-4e1d-a847-85359c7af3dd" />

