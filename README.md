# piBle# Goal Server
       The Goal Server is written in arduino/C++ for the esp32. Whenever a goal is scored the esp32 wakes up from deep sleep and tries to send its information for a few seconds to the Pi. When it gets the confirmation that the data has been received or it times out it goes back into deep sleep.

       ## Table of Contents
       This is autogenerated with https://ecotrust-canada.github.io/markdown-toc/ . Just copy the text and click convert button.
         - [Table of Contents](#table-of-contents)
         - [Architecture](#architecture)
             - [Software management](#software-management)
             - [Communication of IoT Devices](#communication-of-iot-devices)
             - [Future](#future)
         - [Polling alive and updates](#polling-alive-and-updates)
         - [Time synchronization](#time-synchronization)
         - [Sleep states](#sleep-states)

       ## Architecture
       Because a pitcure explains more than a thousand words: <br />
       <img src="pictures/architecture.png" width="400" style="margin:50px" /> <br />
       #### Software management
       Management of the software is done in the cloud and then automatically synchronized between the cloud master and the edge controller. Currently, the edge part is running a full fledged K8S solution, but the plan is to switch something a lot smaller, e.g. K3S which has a 40Mb (compressed) binary. Deployments are all done via Containers and we can reap all the benefits of using K8S, i.e. rolling deployments, health checks, networking. A update for the esp32 is build on the client (or server) inside a container. The container is then deployed on the Pi where it sends the data to the esp32. When the esp32 sends the correct/new version in its beacon the container stops. (Maybe: For updates the esp32 enables a full bluetooth connection.)
       <br />

       #### Communication of IoT Devices
       The Pi and esp32 communicate via BLE, where the esp32 is the Server and the Pi is the client. The BLE is in deep sleep for most of the time and only woken up by a new goal (see Future for battery consideration.).

       #### Future
       Make battery tests for laser and sensor. Should they be running all the time? The esp32 could have a game mode which is triggered by the Pi when a game starts.


       ## Polling alive and updates
       In predefined updates (maybe 30 minutes) the esp32 wakes itself up and polls the raspberry for updates, to synch time and maybe more stuff in the future.

       ## Time synchronization
       Time synchronization is done when the esp32 communicates with Pi. But what happens to the time if esp32 goes into deep sleep


       ## Sleep states
       https://lastminuteengineers.com/esp32-sleep-modes-power-consumption/

       ## Some Infos about the Hardware
       https://circuitdigest.com/microcontroller-projects/getting-started-with-esp32-with-arduino-ide
