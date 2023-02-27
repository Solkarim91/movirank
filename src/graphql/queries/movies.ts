import { gql } from "@apollo/client";

export const GET_MOVIES = gql`
  query getMovies {
    getMovies {
      id
      title
      description
      runtime
      released
      img
    }
  }
`;
