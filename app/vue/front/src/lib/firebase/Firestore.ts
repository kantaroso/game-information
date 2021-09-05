import app from '@/lib/firebase/App'
import { getFirestore } from 'firebase/firestore'
const firestore = getFirestore(app.getInstance())
export default firestore
