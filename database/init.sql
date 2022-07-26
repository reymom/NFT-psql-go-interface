CREATE 
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Users
(
    id      int          NOT NULL GENERATED ALWAYS AS IDENTITY,
    user_id uuid         NOT NULL UNIQUE,
    name    varchar(64)  NOT NULL,
    mail    varchar(128) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);
CREATE INDEX idx_users on Users (user_id);

CREATE TABLE Nfts
(
    nft_id            uuid        DEFAULT uuid_generate_v4(),
    owner_id          int                                    NOT NULL,
    nft_type_id       smallint                               NOT NULL,
    created_at        timestamptz DEFAULT current_timestamp  NOT NULL,
    content           json                                   NOT NULL,
    version           smallint                               NOT NULL,
    PRIMARY KEY (nft_id),
    FOREIGN KEY (owner_id) REFERENCES Users (id) ON DELETE RESTRICT,
);
CREATE INDEX idx_nfts on Nfts (owner_id);
