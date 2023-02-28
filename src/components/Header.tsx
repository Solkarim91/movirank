import {
  MagnifyingGlassIcon,
  HeartIcon,
  UserCircleIcon,
} from "@heroicons/react/24/outline";
import Link from "next/link";
import { useState } from "react";
import Image from "next/image";

export const Header = () => {
  const [searchInput, setSearchInput] = useState("");

  return (
    // Header
    <div className="flex h-20 justify-around py-4 border-b-2 bg-white [&>*]:my-auto">
      {/* Logo */}
      <Link href={"/movies"}>
        <Image
          src={
            "https://res.cloudinary.com/domtywpyc/image/upload/v1677598736/logo-color_ksjklv.png"
          }
          width="100"
          height="100"
          alt={"/"}
        />
      </Link>
      {/* Search */}
      <div className="flex h-10 w-[40%] max-w-[390px] border-2 justify-between px-2 rounded-full [&>*]:my-auto">
        <input
          className="outline-none w-[100%]"
          type="text"
          placeholder="Search movies"
          value={searchInput}
          onChange={(e) => setSearchInput(e.target.value)}
        />
        <button className="w-8 border-l-2 pl-1">
          <MagnifyingGlassIcon className="h-8 w-6" />
        </button>
      </div>
      <div className="flex gap-x-4">
        <button>
          <HeartIcon className="h-8 w-6" />
        </button>
        <button>
          <UserCircleIcon className="h-8 w-6" />
        </button>
      </div>
    </div>
  );
};
