CREATE TABLE IF NOT EXISTS public.players(
    id uuid NOT NULL,
    name VARCHAR(45) NOT NULL,
    email VARCHAR(45) NOT NULL,
    mobile VARCHAR(45),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "player_id_UNIQUE" PRIMARY KEY (id),
    CONSTRAINT "email_UNIQUE" UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.games(
    id uuid NOT NULL,
    name VARCHAR(45) NOT NULL,
    CONSTRAINT "game_id_UNIQUE" PRIMARY KEY (id),
    CONSTRAINT "name_UNIQUE" UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.scores
(
    id uuid NOT NULL,
    game_id uuid NOT NULL,
    player_id uuid NOT NULL,
    score bigint DEFAULT 0,
    region VARCHAR(10) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT game_id_fk FOREIGN KEY (game_id)
    REFERENCES public.games (id) MATCH SIMPLE ON UPDATE CASCADE ON DELETE CASCADE NOT VALID,
    CONSTRAINT player_id_fk FOREIGN KEY (player_id)
    REFERENCES public.players (id) MATCH SIMPLE ON UPDATE CASCADE ON DELETE CASCADE NOT VALID
);
