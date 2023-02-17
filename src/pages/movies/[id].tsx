import fsPromises from "fs/promises";
import { InferGetStaticPropsType } from "next";
import path from "path";
import { MovieItem } from "types/types";

// export const getStaticPaths = async () => {
//   const filePath = path.join(process.cwd(), "mock-data.json");
//   const jsonData = await fsPromises.readFile(filePath, "utf-8");
//   const movieData = JSON.parse(jsonData);

//   const paths = movieData.map((movie: MovieItem) => {
//     return {
//       params: { id: movie.id.toString() },
//     };
//   });

//   return {
//     paths,
//     fallback: false,
//   };
// };

export const getStaticPaths = async () => {
  const res = await fetch(
    "https://my-json-server.typicode.com/solkarim91/movirank-placeholder/movies/"
  );
  const data = await res.json();

  const paths = data.map((movie: MovieItem) => {
    return {
      params: { id: movie.id.toString() },
    };
  });

  return {
    paths,
    fallback: false,
  };
};

export const getStaticProps = async (context: any) => {
  const id = context.params.id;
  const res = await fetch(
    "https://my-json-server.typicode.com/solkarim91/movirank-placeholder/movies/" +
      id
  );
  const data = await res.json();

  return {
    props: { movie: data },
  };
};

export default function Movie({
  movie,
}: InferGetStaticPropsType<typeof getStaticProps>) {
  console.log(movie);
  return (
    <div>
      <h1>{movie.title}</h1>
    </div>
  );
}
