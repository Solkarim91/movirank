import Head from "next/head";
import { Header } from "@/components/Header";
import fsPromises from "fs/promises";
import path from "path";
import { MovieItem } from "types/types";
import { MovieCard } from "@/components/MovieCard";
import { Footer } from "@/components/Footer";
import Link from "next/link";
import { InferGetStaticPropsType } from "next";

// Fetching data from the JSON file
export async function getStaticProps() {
  const filePath = path.join(process.cwd(), "mock-data.json");
  const jsonData = await fsPromises.readFile(filePath, "utf-8");
  const movieData = JSON.parse(jsonData);

  return {
    props: { movieData },
  };
}

export default function Movies({
  movieData,
}: InferGetStaticPropsType<typeof getStaticProps>) {
  // console.log(movieData);
  return (
    <div className="bg-stone-200">
      <Head>
        <title>MoviRank</title>
        <meta name="description" content="Rate your favourite movies!" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      {/* Header */}
      <Header />

      {/* Main */}
      <main className="max-w-7xl h-[100%] mx-auto px-8 bg-white">
        <div className="text-center pt-8 pb-4">
          <h1 className="font-bold text-2xl">All movies</h1>
        </div>

        {/* View options */}
        <div className="flex justify-around">
          <div className="border-1">
            <p>Results per page:</p>
          </div>

          <div className="flex gap-x-8">
            <button>Sort</button>
            <button>Filter</button>
          </div>
        </div>

        {/* Card grid */}
        <div className="container m-auto grid grid-cols-2 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-4 gap-8 pt-8 pb-10">
          {movieData?.map((movie: MovieItem) => (
            <Link
              key={movie.id}
              className="flex flex-col justify-between border-2 cursor-pointer w-82 h-[48vh] sm:h-[48vh] md:h-[48vh] lg:h-[48vh] hover:shadow-md transition duration-200 ease-out"
              href={`/movies/${movie.id}`}
            >
              <MovieCard
                title={movie.title}
                description={movie.description}
                rating={movie.rating}
                comments={movie.comments}
                img={movie.img}
              />
            </Link>
          ))}
        </div>
      </main>

      {/* Footer */}
      <Footer />
    </div>
  );
}
