import { AlbumId, SingerId, SingerName } from "./domain";
import { debutSingerWorkflow } from "./workflow";

const CORRECT_SINGER_ID = "singer-123";

const result1 = debutSingerWorkflow({
  kind: "UnValidatedDebutSingerCommand",
  input: {
    singer: {
      id: CORRECT_SINGER_ID as SingerId,
      name: SingerName("John Doe").unwrap(),
      type: "NoDebut"
    },
    debutAlbumId: "album-123" as AlbumId,
  },
});

// Ok
console.log(result1);

const WRONG_SINGER_ID = "singer-1234";
const result2 = debutSingerWorkflow({
  kind: "UnValidatedDebutSingerCommand",
  input: {
    singer: {
      id: WRONG_SINGER_ID as SingerId,
      name: SingerName("John Doe").unwrap(),
      type: "NoDebut"
    },
    debutAlbumId: "album-123" as AlbumId,
  },
});

// Err
console.log(result2);
