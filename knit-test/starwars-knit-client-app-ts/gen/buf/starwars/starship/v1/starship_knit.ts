// @generated by protoc-gen-knit-ts v0.0.1 with parameter "target=ts"
// @generated from file buf/starwars/starship/v1/starship.proto (package buf.starwars.starship.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

export type StarshipService = {
  "buf.starwars.starship.v1.StarshipService": {
    fetch: {
    };
    do: {
      getStarships: { $: StarshipsRequest; value: StarshipsResponse; };
    };
    listen: {
    };
  };
};

export interface Starship {
  starshipId: string;
  name: string;
  model: string;
  manufacturer: string;
  costInCredits: bigint;
  length: bigint;
  maxAtmospheringSpeed: bigint;
  crew: bigint;
  passengers: bigint;
  cargoCapacity: bigint;
  consumables: string;
  hyperdriveRating: number;
  mglt: bigint;
  starshipClass: string;
  pilotIds: Array<string>;
  filmIds: Array<string>;
};

export interface StarshipsRequest {
  starshipIds: Array<string>;
  limit: number;
};

export interface StarshipsResponse {
  starships: Array<Starship>;
};

