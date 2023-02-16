import Head from "next/head";
import { Header } from "@/components/Header";
import fsPromises from "fs/promises";
import path from "path";
import { MovieCardItem } from "types/types";
import { MovieCard } from "@/components/MovieCard";

export default function Home({ movieData }: any) {
  console.log(movieData);
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
        <div className="container m-x-auto pb-10 grid grid-cols-3 gap-x-8 gap-y-8 pt-8">
          {movieData?.map(
            ({
              id,
              title,
              description,
              runtime,
              released,
              rating,
              comments,
              img,
            }: MovieCardItem) => (
              <MovieCard
                key={id}
                id={id}
                title={title}
                description={description}
                runtime={runtime}
                released={released}
                rating={rating}
                comments={comments}
                img={img}
              />
            )
          )}
        </div>
      </main>
    </div>
  );
}

// Fetching data from the JSON file
export async function getStaticProps() {
  const filePath = path.join(process.cwd(), "mock-data.json");
  const jsonData = await fsPromises.readFile(filePath, "utf-8");
  const movieData = JSON.parse(jsonData);

  return {
    props: { movieData },
  };
}
