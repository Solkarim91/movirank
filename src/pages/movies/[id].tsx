import Head from "next/head";
import { useQuery } from "@apollo/client";
import { useRouter } from "next/router";
import { GET_ONE_MOVIE } from "@/graphql/queries/movie";
import { Header } from "@/components/Header";
import Image from "next/image";
import { Footer } from "@/components/Footer";
import { StarIcon, ArrowLeftIcon } from "@heroicons/react/24/solid";
import { ChatBubbleBottomCenterTextIcon } from "@heroicons/react/24/outline";
import { commentsData } from "@/testdata/comments-data";
import { reviewsData } from "@/testdata/reviews-data";
import Link from "next/link";

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
    <div className="bg-stone-200">
      <Head>
        <title>MoviRank</title>
        <meta name="description" content="Rate your favourite movies!" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      {/* Header */}
      <Header />

      <main className="main pt-2">
        <div className="flex flex-col gap-y-4 pt-2">
          <Link
            className="flex items-center gap-x-1 [&>*]:text-blue-400"
            href={"/movies"}
          >
            <ArrowLeftIcon className="h-6 w-4" />
            <h1 className="text-sm">Back</h1>
          </Link>
          <div className="flex p-8 gap-x-8 border-2 relative">
            <div className="relative h-[24vw] w-[16vw] flex-shrink-0">
              <Image
                src={movie.img}
                fill
                sizes="(max-width: 768px) 100vw,
              (max-width: 1200px) 50vw,
              33vw"
                priority
                alt="/"
              />
            </div>
            <div className="flex flex-col gap-y-2 w-[60%]">
              <div className="flex gap-x-20 text-sm text-gray-600 mb-2">
                <div>
                  <h3>Genre</h3>
                  <p>{movie.genre}</p>
                </div>
                <div>
                  <h3>Released</h3>
                  <p>{movie.released}</p>
                </div>
                <div>
                  <h3>Runtime</h3>
                  <p>{movie.runtime}</p>
                </div>
              </div>
              <div className="mb-2">
                <h1 className="font-black text-2xl">{movie.title}</h1>
                <p>{movie.description}</p>
              </div>
              <div className="flex flex-col gap-y-3 text-sm">
                <div className="flex gap-x-1 items-end">
                  <h3 className="font-bold">Director:</h3>
                  <p>{movie.director}</p>
                </div>
                <div className="flex flex-col gap-y-1">
                  <h3 className="font-bold">Cast:</h3>
                  <div className="grid grid-cols-5 gap-x-8 gap-y-2">
                    {movie.cast.map((actor: string) => (
                      <p className="w-60" key={actor}>
                        {actor}
                      </p>
                    ))}
                  </div>
                </div>
              </div>
              <div className="flex gap-x-8 absolute bottom-5 right-6">
                <div>
                  <a
                    href="#reviews"
                    className="flex items-center justify-center gap-x-1 cursor-pointer"
                  >
                    <StarIcon className="h-8 w-6 text-yellow-400" />
                    <h2>Reviews</h2>
                    <p className="text-sm">({reviewsData.length})</p>{" "}
                  </a>
                </div>
                <div>
                  <a
                    href="#comments"
                    className="flex items-center justify-center gap-x-1"
                  >
                    <ChatBubbleBottomCenterTextIcon className="h-8 w-6" />
                    <h2>Comments</h2>
                    <p className="text-sm">({commentsData.length})</p>{" "}
                  </a>
                </div>
              </div>
            </div>
          </div>
          <div className="flex flex-col gap-y-4 mb-16 [&>*]:p-4">
            <div className="border-2 flex flex-col gap-y-2">
              <a id="reviews" className="font-semibold">
                Reviews
              </a>
              <div className="flex flex-col gap-y-4">
                {reviewsData.map((review) => (
                  <div key={review.id} className="flex">
                    <StarIcon className="h-6 w-4 text-yellow-400" />
                    <div className="flex gap-x-2">
                      <p className="font-light">{review.star} </p>
                      <p>-</p>
                      <p>{review.user.username}</p>
                    </div>
                  </div>
                ))}
              </div>
            </div>
            <div className="border-2 flex flex-col gap-y-2">
              <div className="flex gap-x-2 items-center">
                <ChatBubbleBottomCenterTextIcon className="h-6 w-4" />
                <a id="comments" className="font-semibold">
                  Comments
                </a>
              </div>
              <div className="flex flex-col gap-y-4 mt-2">
                {commentsData.map((comment) => (
                  <div key={comment.id}>
                    <h4 className="font-semibold text-sm">
                      {comment.user.username}
                    </h4>
                    <p className="font-light">{comment.text}</p>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <Footer />
    </div>
  );
}
