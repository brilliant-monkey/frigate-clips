# clips

Container Image: [brilliantmonkey.azurecr.io/frigate/clips](brilliantmonkey.azurecr.io/frigate/clips)

Tags: [0.0.1](brilliantmonkey.azurecr.io/frigate/clips:0.0.1)

## Running Frigate Clips

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
