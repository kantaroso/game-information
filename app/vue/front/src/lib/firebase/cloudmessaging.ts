import app from '@/lib/firebase/app'
import { Messaging, getMessaging, getToken, onMessage } from 'firebase/messaging'
export default class FCMWrapper {
  private static instance: Messaging

  private static getInstance () {
    if (FCMWrapper.instance) {
      return FCMWrapper.instance
    }
    FCMWrapper.instance = getMessaging(app.getInstance())
    return FCMWrapper.instance
  }

  static getToken () {
    const token = localStorage.getItem('fcm_token')
    if (!token) {
      // この関数で通知許諾の表示、service workerの登録を行う
      // https://firebase.google.com/docs/reference/js/messaging_.md#gettoken
      getToken(FCMWrapper.getInstance(), { vapidKey: process.env.VUE_APP_FIREBASE_CLOUD_MESSAGING_VAPIKEY }).then((currentToken) => {
        if (currentToken) {
          localStorage.setItem('fcm_token', currentToken)
        } else {
          console.log('No registration token available. Request permission to generate one.')
        }
      }).catch((err) => {
        console.log('An error occurred while retrieving token. ', err)
      })
    }
  }

  static watchForegroudMessage () {
    onMessage(FCMWrapper.getInstance(), (payload) => {
      // スコープの確認用処理
      // console.log(navigator.serviceWorker.getRegistrations())
      navigator.serviceWorker.getRegistration('/firebase-cloud-messaging-push-scope').then(function (registration) {
        if (!payload.notification) {
          return
        }
        if (registration) {
          const notificationTitle = payload.notification.title ?? ''
          const notificationOptions = {
            body: payload.notification.body,
            icon: '/favicon.ico'
          }
          registration.showNotification(notificationTitle, notificationOptions)
        }
      })
    })
  }
}
