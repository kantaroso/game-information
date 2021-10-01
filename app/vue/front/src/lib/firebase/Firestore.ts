import app from '@/lib/firebase/app'
import { getFirestore } from 'firebase/firestore'
const firestore = getFirestore(app.getInstance())
export default firestore
