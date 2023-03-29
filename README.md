
# Clips

![demo](https://user-images.githubusercontent.com/11224731/228419650-565ffcc4-2213-4312-8331-7c3e0cef01b3.gif)

Clips is a thumbnail generation tool. It will receive a Frigate event once it is finished and turn it into a short 5 second video clip. After generation, it publishes an event to `frigate/clips` topic which allows for automation (sending a notification of the recording).

[Container Images](https://github.com/brilliant-monkey/frigate-clips/pkgs/container/frigate-clips)
Tags:
0.0.1

## Running Clips

### Dependencies

- MQTT
- Frigate

### Setup

1. Copy `docker-compose.yml` to the server running Frigate in Docker
1. Using the `config.yml` as a template, create a config file and place in a location accessible to Docker
1. Run `docker compose up -d` to run the project

## Home Assistant

### Recommended Flow

1. Frigate publishes `new` event of a detection to MQTT
1. Home Assistant triggers a new push notification to a device using the `after.id` property as a `tag` key
1. Clips receives the MQTT `new` event message from the `frigate/events` topic and generates a clip if the event is an `end` (detection is complete).
1. Clips publishes a complete event on the `frigate/clips` MQTT topic
1. Home assistant triggers a new push notification to update the existing notification using the `after.id` property as a `tag` key
1. User can see looping clip on mobile device

### Setup

- Create or edit an existing Frigate notification automation
- Add a trigger of `MQTT event` trigging by event on topic `frigate/clips`
- Using the `notify` service, publish a notification using the below block as a template:
  ```
  - service: notify.notify
    data:
      message: A {{ label }} was detected on the {{ cameraName }} camera.
      data:
        tag: "{{ eventId }}"
        group: frigate
        video: "https://<my_clips_endpoint>{{ trigger.payload_json['clip_url'] }}"
        push:
          sound: none
  ```
