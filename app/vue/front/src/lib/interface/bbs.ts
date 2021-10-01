interface BbsThread {
    id: string;
    title: string;
    name: string;
    body: string;
    createdAt: number;
    updatedAt: number;
  }

interface BbsResponse {
  id: string;
  name: string;
  body: string;
  createdAt: number;
  updatedAt: number;
}

export { BbsThread, BbsResponse }
