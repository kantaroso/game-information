// Import and configure the Firebase SDK
// These scripts are made available when the app is served or deployed on Firebase Hosting
// If you do not serve/host your project using Firebase Hosting see https://firebase.google.com/docs/web/setup
/**
importScripts('/__/firebase/9.0.1/firebase-app-compat.js');
importScripts('/__/firebase/9.0.1/firebase-messaging-compat.js');
importScripts('/__/firebase/init.js');

const messaging = firebase.messaging();
*/

/**
 * Here is is the code snippet to initialize Firebase Messaging in the Service
 * Worker when your app is not hosted on Firebase Hosting.
 */
 // Give the service worker access to Firebase Messaging.
 // Note that you can only use Firebase Messaging here. Other Firebase libraries
 // are not available in the service worker.
 importScripts('https://www.gstatic.com/firebasejs/9.0.1/firebase-app-compat.js');
 importScripts('https://www.gstatic.com/firebasejs/9.0.1/firebase-messaging-compat.js');
 // Initialize the Firebase app in the service worker by passing in
 // your app's Firebase config object.
 // https://firebase.google.com/docs/web/setup#config-object
 firebase.initializeApp({
   apiKey: 'AIzaSyB1yAm0NyqPUecRbqlEdO08s8pWOk7NiEA',
   authDomain: 'bizyoge-dev.firebaseapp.com',
   projectId: 'bizyoge-dev',
   storageBucket: 'bizyoge-dev.appspot.com',
   messagingSenderId: '783224334963',
   appId: '1:783224334963:web:6c847cbc2b7418293b1806',
   measurementId: 'G-6CGEK8B8HZ',
 });
 // Retrieve an instance of Firebase Messaging so that it can handle background
 // messages.
 const messaging = firebase.messaging();


// If you would like to customize notifications that are received in the
// background (Web app is closed or not in browser focus) then you should
// implement this optional method.
// Keep in mind that FCM will still show notification messages automatically
// and you should use data messages for custom notifications.
// For more info see:
// https://firebase.google.com/docs/cloud-messaging/concept-options
messaging.onBackgroundMessage(function(payload) {
  const notificationTitle = payload.notification.title;
  const notificationOptions = {
    body: payload.notification.body,
    icon: '/favicon.ico',
    tag: payload.notification.title
  };
  // self.registration.showNotification(notificationTitle, notificationOptions);
});
