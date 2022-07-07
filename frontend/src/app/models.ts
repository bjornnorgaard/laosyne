export interface SearchFilter {
  take: number;
  skip: number;
  pathContains: string;
  upperRating: number;
  lowerRating: number;
  sortOrder: SortOrder;
}

enum SortOrder {
  RANDOM,
  RATING_DESC,
  RATING_ASC
}

export interface MatchResult {
  winnerId: number;
  loserId: number;
}

export interface Match {
  playerOne: Picture;
  playerTwo: Picture;
}

export interface NewPath {
  path: string;
}

export interface DeletePath {
  pathId: number;
}

export interface Path {
  id: number;
  path: string;
  createdAt: string;
  updatedAt: string;
}

export interface Picture {
  id: number;
  path: string;
  ext: string;
  views: number;
  likes: number;
  rating: number;
  deviation: number;
  wins: number;
  losses: number;
  createdAt: string;
  updatedAt: string;
}
