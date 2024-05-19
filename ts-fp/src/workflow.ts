import { Err, Ok, Result } from "ts-results";
import { Album, AlbumId, AlbumTitle, DebutSinger, NoDebutSinger, NormalAlbum } from "./domain";

interface UnValidatedDebutSingerCommand {
  kind: "UnValidatedDebutSingerCommand";
  input: {
    singer: NoDebutSinger;
    debutAlbumId: AlbumId;
  }
}

interface ValidatedDebutSingerCommand {
  kind: "ValidatedDebutSingerCommand";
  input: {
    singer: NoDebutSinger;
    debutAlbumId: AlbumId;
  }
}

interface DebutSingerCreated {
  kind: "DebutSingerCreated";
  input: {
    singer: NoDebutSinger;
    debutAlbumId: AlbumId;
  }
  singer: DebutSinger;
}

export const validateAlbum = (findIdenticalAlbum: (albumId: AlbumId) => Result<Album, Error>) => (command: UnValidatedDebutSingerCommand): Result<ValidatedDebutSingerCommand, Error> => {
  const album = findIdenticalAlbum(command.input.debutAlbumId).unwrap();

  if (album.singerId !== command.input.singer.id) {
    return Err(new Error("Singer is not the same as the album's singer"));
  }

  return Ok({
    kind: "ValidatedDebutSingerCommand",
    input: command.input,
  })
}

export const createDebutSinger = (command: ValidatedDebutSingerCommand): Result<DebutSingerCreated, Error> => {
  return Ok({
    kind: "DebutSingerCreated",
    input: command.input,
    singer: {
      ...command.input.singer,
      debutAlbum: command.input.debutAlbumId,
      type: "Debut", // Add the correct type here
    }
  })
}

const mockFindIdenticalAlbum = (albumId: AlbumId): Result<Album, Error> => {
  const title = AlbumTitle("デュブルアルバム").unwrap();
  return Ok(
    {
      id: albumId,
      type: "Normal",
      title: title,
      releaseDate: new Date(),
      singerId: "singer-123",
      stopped: false,
    } as NormalAlbum
  )
}

type DebutSingerWorkflow = (command: UnValidatedDebutSingerCommand) => Result<DebutSinger, Error>;

export const debutSingerWorkflow: DebutSingerWorkflow = (command: UnValidatedDebutSingerCommand):  Result<DebutSinger, Error> => Ok(command)
  .andThen(validateAlbum(mockFindIdenticalAlbum))
  .andThen(createDebutSinger)
  .map((result) => result.singer)
