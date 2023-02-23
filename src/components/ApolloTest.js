// File currently only being used as a reference

import { useQuery, gql } from "@apollo/client";

const QUERY = gql`
  query getAllMovies {
    GetAllMovies {
      id
      title
    }
  }
`;

export default function PastLaunches() {
  const { data, loading, error } = useQuery(QUERY);

  if (loading) {
    return <h2>Loading Data...</h2>;
  }

  if (error) {
    console.error(error);
    return <h2>Sorry, there has been an error...</h2>;
  }

  const allMovies = data.GetAllMovies;

  console.log(allMovies);

  return (
    <>
      <h1>test</h1>
    </>
  );
}
