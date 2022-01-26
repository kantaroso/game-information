importScripts('/__/firebase/9.0.1/firebase-app-compat.js');
importScripts('/__/firebase/9.0.1/firebase-messaging-compat.js');
importScripts('/__/firebase/init.js');
const messaging = firebase.messaging();

messaging.onBackgroundMessage(function(payload) {
  const notificationTitle = payload.notification.title;
  const notificationOptions = {
    body: payload.notification.body,
    icon: '/favicon.ico',
    tag: payload.notification.title
  };
});
