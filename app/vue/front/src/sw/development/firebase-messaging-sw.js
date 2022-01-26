importScripts('https://www.gstatic.com/firebasejs/9.0.1/firebase-app-compat.js');
importScripts('https://www.gstatic.com/firebasejs/9.0.1/firebase-messaging-compat.js');
firebase.initializeApp({
  apiKey: 'process.env.VUE_APP_FIREBASE_API_KEY',
  authDomain: 'process.env.VUE_APP_FIREBASE_AUTH_DOMAIN',
  projectId: 'process.env.VUE_APP_FIREBASE_PROJECT_ID',
  storageBucket: 'process.env.VUE_APP_FIREBASE_STORAGE_BUCKET',
  messagingSenderId: 'process.env.VUE_APP_FIREBASE_MESSAGING_SENDER_ID',
  appId: 'process.env.VUE_APP_FIREBASE_APP_ID',
  measurementId: 'process.env.VUE_APP_FIREBASE_MEASUREMENT_ID',
});
const messaging = firebase.messaging();

messaging.onBackgroundMessage(function(payload) {
  const notificationTitle = payload.notification.title;
  const notificationOptions = {
    body: payload.notification.body,
    icon: '/favicon.ico',
    tag: payload.notification.title
  };
});
