export interface Movie {
  id: string;
  title: string;
  description: string;
  genre: string;
  runtime: number;
  released: number;
  // rating: number;          Need to implement rating/comments on server side first
  // comments: Comment[];
  img: string;
}

export interface MovieCardItem {
  title: string;
  description: string;
  // rating: number;
  // comments: Comment[];
  img: string;
}

export interface Comment {
  id: string;
  username: string;
  text: string;
}

export interface Review {
  id: string;
  user: User;
  star: number;
}

export interface User {
  id: string;
  userType: UserType;
  username: string;
  email: string;
}

enum UserType {
  ADMIN,
  STANDARD,
}
