import { gql } from "@apollo/client";

export const GET_ONE_MOVIE = gql`
  query getOneMovie($id: ID!) {
    getOneMovie(id: $id) {
      id
      title
      description
      genre
      released
      runtime
      img
    }
  }
`;
