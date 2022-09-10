#!/bin/bash
SERVER_KEY='AIzaSyAI-JUxa-iA1yL_zofGqX7paU6mVABJ05s'
DEVICE_REG_TOKEN='eqawwKIgBos8ywEjM8g3ap:APA91bFVkWF1Nz5nlnxp51Q8b0FcSH0JSbObUthfBnfbHgWZaN08h0zepjl2_qsoe4gUllGmMaGZCqAdzXOAGxM6PMbYEdycPHMYlI-0n83ujuW-yv5-6uFfhACkjrcyzPDsHSHKlCgP'

CMD=$(cat <<END
curl -X POST -H "Authorization: key=$SERVER_KEY" -H "Content-Type: application/json"
   -d '{
  "data": {
    "notification": {
        "title": "FCM Message",
        "body": "This is an FCM Message",
        "icon": "/itwonders-web-logo.png",
    }
  },
  "to": "$DEVICE_REG_TOKEN"
}' https://fcm.googleapis.com/fcm/send
END
)

echo $CMD && eval $CMD