import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
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
  DislikePicture: Scalars['Boolean'];
  LikePicture: Scalars['Boolean'];
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
  CreateMatch: Match;
  GetPaths: Array<Path>;
  GetPicture: Picture;
  GetPictures?: Maybe<Array<Picture>>;
};


export type QueryCreateMatchArgs = {
  input?: InputMaybe<SearchFilter>;
};


export type QueryGetPictureArgs = {
  pictureId: Scalars['Int'];
};


export type QueryGetPicturesArgs = {
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
  Random = 'RANDOM',
  RatingAsc = 'RATING_ASC',
  RatingDesc = 'RATING_DESC'
}

export type GetPictureDetailsQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type GetPictureDetailsQuery = { __typename?: 'Query', GetPicture: { __typename?: 'Picture', id: number, path: string, ext: string, views: number, likes: number, losses: number, wins: number, rating: number, deviation: number, updatedAt: string, createdAt: string } };

export const GetPictureDetailsDocument = gql`
    query GetPictureDetails($id: Int!) {
  GetPicture(pictureId: $id) {
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
  export class GetPictureDetailsGQL extends Apollo.Query<GetPictureDetailsQuery, GetPictureDetailsQueryVariables> {
    document = GetPictureDetailsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }