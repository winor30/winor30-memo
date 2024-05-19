import { Err, Ok, Result } from 'ts-results';
import { v4 as uuidv4 } from 'uuid';

// TODO: what
declare const __newtype: unique symbol;

// newTypeは、基本型(Type)に対して新しい型(Constructor)を作成します。
// これにより、同じ基本型を持つ異なる型を区別できます。
export type newType<Constructor, Type> = Type & { readonly [__newtype]: Constructor };

// SingerIdはstring型を基にした新しい型です。
export type SingerId = newType<Singer, string>;

export const SingerId = (): SingerId => {
  const id = uuidv4();
  return id as SingerId;
}

export type SingerName = newType<Singer, string>;

export const SingerName = (name: string): Result<SingerName, Error> => {
  return name.length > 0 && name.length < 100 ? Ok(name as SingerName) : Err(new Error('Invalid name'));
}

export type SingerType = "Debut" | "NoDebut";

export type Singer = DebutSinger | NoDebutSinger;

interface _Singer {
  id: SingerId;
  name: SingerName;
  type: SingerType;
}

export interface DebutSinger extends _Singer {
  type: "Debut";
  debutAlbum: AlbumId;
}

export interface NoDebutSinger extends _Singer {
  type: "NoDebut";
}

// AlbumIdはstring型を基にした新しい型です。
export type AlbumId = newType<_Album, string>;

export const AlbumId = (): AlbumId => {
  const id = uuidv4();
  return id as AlbumId;
}

export type Album = NormalAlbum | BestAlbum;
type AlbumType = "Best" | "Normal";

export type AlbumTitle = newType<_Album, string>;

export const AlbumTitle = (title: string): Result<AlbumTitle, Error> => {
  return title.length > 0 && title.length < 100 ? Ok(title as AlbumTitle) : Err(new Error('Invalid title'));
}

interface _Album {
  id: AlbumId;
  type: AlbumType;
  title: AlbumTitle;
  releaseDate?: Date;
  singerId: SingerId;
  stopped: boolean;
}

export interface BestAlbum extends _Album {
  type: "Best";
  bestSongs: string[];
}

export interface NormalAlbum extends _Album {
  type: "Normal";
}
