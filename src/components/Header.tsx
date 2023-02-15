import {
  MagnifyingGlassIcon,
  HeartIcon,
  UserCircleIcon,
} from "@heroicons/react/24/solid";
import { useState } from "react";

export const Header = () => {
  const [searchInput, setSearchInput] = useState("");

  return (
    // Header
    <div className="flex h-20 justify-around py-4 border-b-2 [&>*]:my-auto">
      {/* Logo */}
      <div>MoviRank</div>
      {/* Search */}
      <div className="flex h-10 w-80 border-2 justify-between px-2 rounded-full [&>*]:my-auto">
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
      {/* Buttons / Icons */}
      <div className="flex gap-x-12">
        {/* Sort / Filter */}
        <div className="flex gap-x-8">
          <button>Sort</button>
          <button>Filter</button>
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
    </div>
  );
};
