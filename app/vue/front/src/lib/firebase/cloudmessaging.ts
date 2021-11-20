import app from '@/lib/firebase/app'
import { getMessaging, getToken } from 'firebase/messaging'
const cloudmessaging = getMessaging(app.getInstance())
const token = localStorage.getItem('fcm_token')
if (!token) {
  // この関数で通知許諾の表示、service workerの登録を行う
  // https://firebase.google.com/docs/reference/js/messaging_.md#gettoken
  getToken(cloudmessaging, { vapidKey: process.env.VUE_APP_FIREBASE_CLOUD_MESSAGING_VAPIKEY }).then((currentToken) => {
    if (currentToken) {
      localStorage.setItem('fcm_token', currentToken)
    } else {
      console.log('No registration token available. Request permission to generate one.')
    }
  }).catch((err) => {
    console.log('An error occurred while retrieving token. ', err)
  })
}
export default cloudmessaging
