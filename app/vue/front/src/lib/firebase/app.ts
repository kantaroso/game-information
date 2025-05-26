import { FirebaseApp, initializeApp } from 'firebase/app'
export default class FirebaseAppWrapper {
  private static instance: FirebaseApp

  static getInstance () {
    if (FirebaseAppWrapper.instance) {
      return FirebaseAppWrapper.instance
    }
    FirebaseAppWrapper.instance = initializeApp(FirebaseAppWrapper.getConfig())
    return FirebaseAppWrapper.instance
  }

  private static getConfig () {
    const config = {
      apiKey: process.env.VUE_APP_FIREBASE_API_KEY,
      authDomain: process.env.VUE_APP_FIREBASE_AUTH_DOMAIN,
      projectId: process.env.VUE_APP_FIREBASE_PROJECT_ID,
      storageBucket: process.env.VUE_APP_FIREBASE_STORAGE_BUCKET,
      messagingSenderId: process.env.VUE_APP_FIREBASE_MESSAGING_SENDER_ID,
      appId: process.env.VUE_APP_FIREBASE_APP_ID,
      measurementId: process.env.VUE_APP_FIREBASE_MEASUREMENT_ID
    }
    return config
  }
}
