import Image from "next/image";
import { StarIcon } from "@heroicons/react/24/solid";
import { ChatBubbleBottomCenterTextIcon } from "@heroicons/react/24/outline";
import { MovieCardItem } from "../types/types";
import { commentsData } from "@/testdata/comments-data";
import { reviewsData } from "@/testdata/reviews-data";

export const MovieCard = (movie: MovieCardItem) => {
  return (
    <>
      <div className="flex flex-col">
        <div className="relative h-60 w-fill flex-shrink-0">
          <Image
            src={movie.img}
            fill
            className="object-cover"
            sizes="(max-width: 768px) 100vw,
              (max-width: 1200px) 50vw,
              33vw"
            priority
            alt="/"
          />
        </div>
        <div className="flex flex-col px-3 w-70 gap-y-1">
          <h4 className="text-base font-bold">{movie.title}</h4>
          <p className="text-sm h-10 line-clamp-2">{movie.description}</p>
        </div>
      </div>
      <div className="flex justify-between items-center mx-3 mb-1">
        <div className="flex gap-x-3 [&>*]:gap-x-1 [&>*]:flex [&>*]:items-center">
          <div>
            <StarIcon className="h-8 w-6 text-yellow-400" />
            <p className="text-sm">{reviewsData.length}</p>{" "}
            {/* Review currently hard coded */}
          </div>
          <div>
            <ChatBubbleBottomCenterTextIcon className="h-8 w-6" />
            <p className="text-sm">{commentsData.length}</p>{" "}
            {/* Comments currently hard coded */}
          </div>
        </div>
        <div className="text-sm font-semibold text-gray-500 underline">
          See more
        </div>
      </div>
    </>
  );
};
