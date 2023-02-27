import Head from "next/head";
import { useQuery } from "@apollo/client";
import { useRouter } from "next/router";
import { GET_ONE_MOVIE } from "@/graphql/queries/movie";

export default function MoviePage() {
  const router = useRouter();
  const { id } = router.query;
  const { data, loading, error } = useQuery(GET_ONE_MOVIE, {
    variables: {
      id: id,
    },
  });

  if (loading) {
    return <h2>Loading Data...</h2>;
  }

  if (error) {
    console.error(error);
    return <h2>Sorry, there has been an error...</h2>;
  }

  const movie = data.getOneMovie;

  return (
    <>
      <Head>
        <title>MoviRank</title>
        <meta name="description" content="Rate your favourite movies!" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div>
        <h1>{movie.title}</h1>
        <h1>{movie.description}</h1>
        <h1>{movie.genre}</h1>
        <h1>{movie.released}</h1>
        <h1>{movie.runtime}</h1>
        <h1>{movie.img}</h1>
      </div>
    </>
  );
}
