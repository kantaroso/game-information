declare module '*.vue' {
  import Vue from 'vue'
  export default Vue
}

declare module '@/lib/firebase/app' {
  const app: any
  export default app
}

declare module '@/lib/firebase/firestore' {
  const firestore: any
  export default firestore
}