# LogU - a bullet journal project

## Goals
- Learn a bit about golang, bootstrap and front-end security practices
## Features
- Web UI having a bullet journal per day
    - tasks, notes and events as different sections and different models

- present day is the default view
- browse previous days
## Architecture
### Step 1
- Server side rendered html DONE
    - Calendar browsing DONE
        - go to present day by default DONE
        - ?date=2021-04-05 to browse by date DONE
- Database models DONE
- Basic login/logout (no auth) DONE

### Step 2
- Front api
- Javascript tracks changes to save
- Users api
- Bullet journal API
- User customization (fonts, backgrounds)


### Step 3
- Step 2 +
- User stats
    - tasks done total
    - tasks total
    - tasks(done) by date/month/year
- Tracing
- Metrics
### Step 4
- Step3 +
- Cqrs/kafka