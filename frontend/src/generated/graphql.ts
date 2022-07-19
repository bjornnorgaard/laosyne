import * as Apollo from 'apollo-angular';
import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';

export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type DeletePath = {
  pathId: Scalars['Int'];
};

export type Match = {
  __typename?: 'Match';
  playerOne: Picture;
  playerTwo: Picture;
};

export type MatchResult = {
  loserId: Scalars['Int'];
  winnerId: Scalars['Int'];
};

export type Mutation = {
  __typename?: 'Mutation';
  AddPath: Path;
  AddToRating: Picture;
  DeletePath: Scalars['Boolean'];
  DislikePicture: Picture;
  LikePicture: Picture;
  ReportMatchResult: Scalars['Boolean'];
  ScanPaths: Scalars['Boolean'];
};


export type MutationAddPathArgs = {
  input: NewPath;
};


export type MutationAddToRatingArgs = {
  pictureId: Scalars['Int'];
};


export type MutationDeletePathArgs = {
  pathId: Scalars['Int'];
};


export type MutationDislikePictureArgs = {
  pictureId: Scalars['Int'];
};


export type MutationLikePictureArgs = {
  pictureId: Scalars['Int'];
};


export type MutationReportMatchResultArgs = {
  input: MatchResult;
};

export type NewPath = {
  path: Scalars['String'];
};

export type Path = {
  __typename?: 'Path';
  createdAt: Scalars['String'];
  id: Scalars['Int'];
  path: Scalars['String'];
  updatedAt: Scalars['String'];
};

export type Picture = {
  __typename?: 'Picture';
  createdAt: Scalars['String'];
  deviation: Scalars['Float'];
  ext: Scalars['String'];
  id: Scalars['Int'];
  likes: Scalars['Int'];
  losses: Scalars['Int'];
  path: Scalars['String'];
  rating: Scalars['Float'];
  updatedAt: Scalars['String'];
  views: Scalars['Int'];
  wins: Scalars['Int'];
};

export type Query = {
  __typename?: 'Query';
  Match: Match;
  Paths: Array<Path>;
  Picture: Picture;
  Pictures?: Maybe<Array<Picture>>;
};


export type QueryMatchArgs = {
  input?: InputMaybe<SearchFilter>;
};


export type QueryPictureArgs = {
  pictureId: Scalars['Int'];
};


export type QueryPicturesArgs = {
  input?: InputMaybe<SearchFilter>;
};

export type SearchFilter = {
  lowerRating?: InputMaybe<Scalars['Int']>;
  pathContains?: InputMaybe<Scalars['String']>;
  skip?: InputMaybe<Scalars['Int']>;
  sortOrder?: InputMaybe<SortOrder>;
  take?: InputMaybe<Scalars['Int']>;
  upperRating?: InputMaybe<Scalars['Int']>;
};

export enum SortOrder {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  Id = 'ID',
  LikesAsc = 'LIKES_ASC',
  LikesDesc = 'LIKES_DESC',
  Random = 'RANDOM',
  RatingAsc = 'RATING_ASC',
  RatingDesc = 'RATING_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC',
  ViewsAsc = 'VIEWS_ASC',
  ViewsDesc = 'VIEWS_DESC'
}

export type PictureDetailsQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type PictureDetailsQuery = { __typename?: 'Query', Picture: { __typename?: 'Picture', id: number, path: string, ext: string, views: number, likes: number, losses: number, wins: number, rating: number, deviation: number, updatedAt: string, createdAt: string } };

export type LikePictureMutationVariables = Exact<{
  id: Scalars['Int'];
}>;


export type LikePictureMutation = { __typename?: 'Mutation', LikePicture: { __typename?: 'Picture', id: number, path: string, ext: string, views: number, likes: number, losses: number, wins: number, rating: number, deviation: number, updatedAt: string, createdAt: string } };

export type DislikePictureMutationVariables = Exact<{
  id: Scalars['Int'];
}>;


export type DislikePictureMutation = { __typename?: 'Mutation', DislikePicture: { __typename?: 'Picture', id: number, path: string, ext: string, views: number, likes: number, losses: number, wins: number, rating: number, deviation: number, updatedAt: string, createdAt: string } };

export type RatePictureMutationVariables = Exact<{
  id: Scalars['Int'];
}>;


export type RatePictureMutation = { __typename?: 'Mutation', AddToRating: { __typename?: 'Picture', id: number, path: string, ext: string, views: number, likes: number, losses: number, wins: number, rating: number, deviation: number, updatedAt: string, createdAt: string } };

export type InspectorSearchQueryVariables = Exact<{
  input?: InputMaybe<SearchFilter>;
}>;


export type InspectorSearchQuery = { __typename?: 'Query', Pictures?: Array<{ __typename?: 'Picture', id: number }> | null };

export type RescanPathsMutationVariables = Exact<{ [key: string]: never; }>;


export type RescanPathsMutation = { __typename?: 'Mutation', ScanPaths: boolean };

export const PictureDetailsDocument = gql`
  query PictureDetails($id: Int!) {
    Picture(pictureId: $id) {
      id
      path
      ext
      views
      likes
      losses
      wins
      rating
      deviation
      updatedAt
      createdAt
    }
  }
`;

@Injectable({
  providedIn: 'root'
})
export class PictureDetailsGQL extends Apollo.Query<PictureDetailsQuery, PictureDetailsQueryVariables> {
  document = PictureDetailsDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}

export const LikePictureDocument = gql`
  mutation LikePicture($id: Int!) {
    LikePicture(pictureId: $id) {
      id
      path
      ext
      views
      likes
      losses
      wins
      rating
      deviation
      updatedAt
      createdAt
    }
  }
`;

@Injectable({
  providedIn: 'root'
})
export class LikePictureGQL extends Apollo.Mutation<LikePictureMutation, LikePictureMutationVariables> {
  document = LikePictureDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}

export const DislikePictureDocument = gql`
  mutation DislikePicture($id: Int!) {
    DislikePicture(pictureId: $id) {
      id
      path
      ext
      views
      likes
      losses
      wins
      rating
      deviation
      updatedAt
      createdAt
    }
  }
`;

@Injectable({
  providedIn: 'root'
})
export class DislikePictureGQL extends Apollo.Mutation<DislikePictureMutation, DislikePictureMutationVariables> {
  document = DislikePictureDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}

export const RatePictureDocument = gql`
  mutation RatePicture($id: Int!) {
    AddToRating(pictureId: $id) {
      id
      path
      ext
      views
      likes
      losses
      wins
      rating
      deviation
      updatedAt
      createdAt
    }
  }
`;

@Injectable({
  providedIn: 'root'
})
export class RatePictureGQL extends Apollo.Mutation<RatePictureMutation, RatePictureMutationVariables> {
  document = RatePictureDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}

export const InspectorSearchDocument = gql`
  query InspectorSearch($input: SearchFilter) {
    Pictures(input: $input) {
      id
    }
  }
`;

@Injectable({
  providedIn: 'root'
})
export class InspectorSearchGQL extends Apollo.Query<InspectorSearchQuery, InspectorSearchQueryVariables> {
  document = InspectorSearchDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}

export const RescanPathsDocument = gql`
  mutation RescanPaths {
    ScanPaths
  }
`;

@Injectable({
  providedIn: 'root'
})
export class RescanPathsGQL extends Apollo.Mutation<RescanPathsMutation, RescanPathsMutationVariables> {
  document = RescanPathsDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}
