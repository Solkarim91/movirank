export interface Movie {
  id: string;
  title: string;
  description: string;
  runtime: number;
  released: number;
  rating: number;
  comments: Comment[];
  img: string;
}

export interface MovieCardItem {
  title: string;
  description: string;
  rating: number;
  comments: Comment[];
  img: string;
}

export interface Comment {
  id: string;
  username: string;
  comment: string;
}
