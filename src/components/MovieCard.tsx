import Image from "next/image";
import { MovieCardItem } from "types/types";

export const MovieCard = ({
  id,
  title,
  description,
  runtime,
  released,
  rating,
  img,
}: MovieCardItem) => {
  return (
    <div className="flex flex-col border-2">
      <div className="relative h-24 w-40 md:h-52 md:w-80 flex-shrink-0">
        <Image
          src={img}
          layout="fill"
          objectFit="cover"
          alt="/"
          className="rounded-2xl"
        />
      </div>
      <div>{title}</div>
      <div>{description}</div>
    </div>
  );
};
